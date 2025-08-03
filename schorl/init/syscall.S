bits 64
section .text
global mount
global check
global open
global write

check:
    jmp $

mount:
    mov rax, 165
    syscall
    ret

open:
    mov rax, 2
    syscall
    ret

write:
    mov rax, 1
    syscall
    ret