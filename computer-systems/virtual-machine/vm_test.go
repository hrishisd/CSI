package vm

import "testing"

func TestExecuteProgram(t *testing.T) {
	t.Run("add two bytes of data", func(t *testing.T) {
		data := [8]byte{
			1: 2,
			2: 3,
		}
		instructions := []Instruction{
			// Load byte 1 into register 1
			Load(R1, 1),
			// Load byte 2 into register 1
			Load(R2, 2),
			// Add the values in the two registers
			Add(R1, R2),
			// Store the result in R1 into the output location
			Store(R1, out),
			Halt()}

		result := executeProgram(data, instructions)
		if result != 5 {
			t.Errorf("Expected 5 but got %x", result)
		}
	})

	t.Run("subtract two bytes of data", func(t *testing.T) {
		data := [8]byte{
			1: 3,
			2: 2,
		}
		instructions := []Instruction{
			// Load byte 1 into register 1
			Load(R1, 1),
			// Load byte 2 into register 1
			Load(R2, 2),
			// Add the values in the two registers
			Sub(R1, R2),
			// Store the result in R1 into the output location
			Store(R1, out),
			Halt()}

		result := executeProgram(data, instructions)
		if result != 1 {
			t.Errorf("Expected 1 but got %x", result)
		}
	})
}
