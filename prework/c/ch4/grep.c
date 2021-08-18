#include <stdio.h>
#include <string.h>
#include <stdlib.h>

const int MAX_LINE_LEN = 1000;

int readline(char dest[])
{
    char c;
    int i = 0;
    while ((c = getchar()) != EOF && c != '\n')
    {
        dest[i++] = c;
    }
    dest[i] = '\0';
    return i;
}

int stridx(char s[], char target[])
{
    size_t size = strlen(s);
    for (int i = 0; i < size; i++)
    {
        int j;
        int k;
        for (j = 0, k = i; target[j] != '\0' && s[k] != '\0'; j++, k++)
        {
            if (target[j] != s[k])
                break;
        }
        if (target[j] == '\0')
        {
            // we made it to the end of the target
            return i;
        }
    }
    return -1;
}

int main()
{
    char pattern[] = "hello";

    char line[MAX_LINE_LEN];
    while (readline(line) != 0)
    {
        if (stridx(line, pattern) != -1)
        {
            printf("%s\n", line);
        }
    }
}