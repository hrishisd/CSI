section .text
global pangram
pangram:
    ; %rcx character and shift count (%cl)
    mov rcx, 0
	; %rdi argument (address of string), incremented throughout
	; %r8 bitmap of letters
	mov r8, 0
.loop:
	; store the character in %cl (8 bits of %rcx)
	movzx rcx, byte [rdi]
	; if its the null terminator, jump to the return
	cmp rcx, 0 
	je .return
	; increment rdi to point to next character
	inc rdi
	; %cl contains an ascii letter if mask with 0x40 is non-zero 
	; (not strictly true since 0x40 isn't a letter, but doesn't matter)
	; store the byte in rsi and mask rsi
	mov rsi, rcx
	and rsi, 040h
	je .loop
	; preserve the last 5 bits: %rcx &= 0x1f 
	and rcx, 01Fh
	; put 1 in %rdx, shift left by the value in %cl
	mov rdx, 1
	sal rdx, cl
	; %r8 |= %rdx
	or r8, rdx
	; go to loop beginning
	jmp .loop
.return:
    ; first store 0 in rax
    mov rax, 0 
    ; mask rax to isolate 26 bits
	; we want these bits: 0...01...10
	and r8, 07FFFFFEh
	cmp r8, 07FFFFFEh
	; if all bits were hi, set last 8 bits of rax to 1
	sete al
	; xor with those 26 bits set high, all others low
	; will be 0 if all were set, non-zero if any were unset
	ret
