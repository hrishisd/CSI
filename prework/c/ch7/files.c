#include <stdio.h>
#include <stdlib.h>

int main(void)
{
    FILE *f = fopen("files.c", "r");
    char c;
    while ((c = getc(f)) != EOF)
    {
        putchar(c);
    }
    putchar('\n');
    fclose(f); // called automatically for all open file descriptors when a program terminates normally
    exit(0);
}