#include <stdio.h>

int _strlen(char *s)
{
    int len = 0;
    while (*s != '\0')
    {
        len++;
        s++;
    }
    return len;
}

void _strcpy(char *src, char *dest)
{
    while ((*src++ = *dest++))
        ;
}

// Copy t to the end of s
void _strcat(char *s, char *t)
{
    while (*s != '\0')
        s++;

    // now s points to the null terminator
    _strcpy(s, t);
}

// Returns <0 if s < t, 0 if s == t, >0 if s > t
int _strcmp(char *s, char *t)
{
    while (*s == *t)
    {
        if (*s == '\0')
            return 0;
        s++;
        t++;
    }
    return *s - *t;
}

int main(void)
{
    char *s = "hello";
    char *t = " world";
    printf("%d\n", _strlen("hello"));
    _strcat(s, t);
    printf("%s\n", s);
}
