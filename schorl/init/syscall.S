bits 64
section .text
global init_modules
global finit_modules

init_modules:
    mov rax, 0xaf
    syscall
    ret

finit_modules:
    mov rax, 0x139
    syscall
    ret