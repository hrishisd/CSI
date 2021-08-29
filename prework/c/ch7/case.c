#include <stdio.h>

static const char UPPER = 'u';
static const char LOWER = 'l';

char lower(char c)
{
    if ('A' >= c && c <= 'Z')
    {
        return 'a' + (c - 'A');
    }
    return c;
}

char upper(char c)
{
    if ('a' >= c && c <= 'z')
    {
        return 'A' + (c - 'a');
    }
    return c;
}

int main(int argc, char *argv[])
{
    if (argc < 2 || (argv[1][0] != UPPER && argv[1][0] != LOWER))
    {
        fprintf(stderr, "Must provide an argument (u or l)\n");
    }

    char c;
    while ((c = getchar()) != EOF)
    {
        putchar(
            argv[1][0] == UPPER ? upper(c) : lower(c));
    }
}
