section .text
global sum_to_n
sum_to_n:
  mov rax, 1
  add rax, rdi
  imul rax, rdi
  sar rax, 1
  ret