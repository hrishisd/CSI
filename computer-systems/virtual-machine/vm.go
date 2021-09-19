package vm

import (
	"fmt"
)

/*
  This is a VM for an imaginary simplified computer architecture.

  The architecture has 2 regular registers and one register to hold the program counter.

  Main memory is 256 bytes.
  Data is stored in the first 8 bytes.
  The zero'th byte holds the output of the program.

  00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f ... ff
  __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ ... __
  XX ^==OTHER DATA======^ ^==INSTRUCTIONS==============^

  The VM supports 5 instructions. Here are their op codes:
  load    0x01
  store   0x02
  add     0x03
  sub     0x04
  halt    0xff

  This is the specification for the instructions:
  load    r1  addr    # Load value at given address into given register
  store   r2  addr    # Store the value in register at the given memory address
  add     r1  r2      # Set r1 = r1 + r2
  sub     r1  r2      # Set r1 = r1 - r2
  halt
  Each parameter is a single byte.
*/

// output location
const out byte = 0x00

// registers
type Register = byte

const (
	R1 Register = 0x00
	R2 Register = 0x01
)

// instruction op codes
const (
	load  byte = 0x01
	store byte = 0x02
	add   byte = 0x03
	sub   byte = 0x04
	halt  byte = 0xff
	addi  byte = 0x05
	subi  byte = 0x06
	jump  byte = 0x07
	beqz  byte = 0x08
)

func executeProgram(data [8]byte, program []Instruction) byte {
	// load the program into memory
	var memory = makeProgram(data, program)
	var registerFile = map[Register]byte{}
	// Initialize PC to be address of first instruction
	var pc byte = 0x08

	for inst := memory[pc]; inst != halt; inst = memory[pc] {
		switch inst {
		case load:
			r, addr := memory[pc+1], memory[pc+2]
			if !validDataAddr(addr) {
				panic(fmt.Sprintf("attempted to read from invalid data location: %x", addr))
			}
			registerFile[Register(r)] = memory[addr]
			pc += 3
		case store:
			r, addr := memory[pc+1], memory[pc+2]
			if !validDataAddr(addr) {
				panic(fmt.Sprintf("attempted to write to invalid data location: %x", addr))
			}
			memory[addr] = registerFile[r]
			pc += 3
		case add:
			r1, r2 := memory[pc+1], memory[pc+2]
			registerFile[r1] += registerFile[r2]
			pc += 3
		case sub:
			r1, r2 := memory[pc+1], memory[pc+2]
			registerFile[r1] -= registerFile[r2]
			pc += 3
		case addi:
			r, i := memory[pc+1], memory[pc+2]
			registerFile[r] += i
		case subi:
			r, i := memory[pc+1], memory[pc+2]
			registerFile[r] -= i
		case jump:
			addr := memory[pc+1]
			pc = addr
		case beqz:
			r, offset := memory[pc+1], memory[pc+2]
			pc += 3
			if registerFile[r] == 0 {
				pc += offset
			}
		default:
			panic(fmt.Errorf("illegal op-code: %x", inst))
		}
	}
	// return the value of the output location
	return memory[0]
}

// Utility functions to construct programs
type Instruction []byte

func Load(r Register, addr byte) Instruction {
	return Instruction{load, byte(r), addr}
}

func Store(r Register, addr byte) Instruction {
	return Instruction{store, byte(r), addr}
}

func Add(r1 Register, r2 Register) Instruction {
	return Instruction{add, byte(r1), byte(r2)}
}

func Sub(r1 Register, r2 Register) Instruction {
	return Instruction{sub, byte(r1), byte(r2)}
}

func Halt() Instruction {
	return Instruction{halt}
}

func validDataAddr(addr byte) bool {
	return addr <= 7
}

func makeProgram(data [8]byte, instructions []Instruction) [256]byte {
	// Doesn't check for overflow
	var result [256]byte
	// copy data into result
	copy(result[:8], data[:])
	idx := 8
	for _, inst := range instructions {
		for _, b := range inst {
			result[idx] = b
			idx++
		}
	}
	return result
}
