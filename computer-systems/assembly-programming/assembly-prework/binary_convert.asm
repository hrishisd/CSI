section .text
global binary_convert
binary_convert:
    ; first, zero rax
	mov rax, 0
.loop:
    ; will need to access char array at %rdi
	; store next character from index at %rdi into rsi (sil)
	movzx esi, byte [rdi]
	; if character in rsi is null terminator (0), go to done
	cmp esi, 0
	je .done
	; otherwise, figure out whether it is a '1' or '0'
	; '0' is 011 0000
	; '1' is 011 0001
	; preserve last 4 bits 
	mov edx, 0000Fh ; store a mask in edx
	and esi, edx
	; left shift rax
	sal rax, 1
	; set the last bit of rax to the 0 or 1
	or rax, rsi
	inc rdi
	jmp .loop
.done:
	ret
