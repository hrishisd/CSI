#include <stdio.h>
#define MAX_LINE_LEN 1000
#define MAX_LINES 5000

/* store the next line from stdin in dest
   return the length of the line */
int _getline(char *dest);

int main(void)
{
    char dest[MAX_LINE_LEN];
    _getline(dest);
    printf("%s\n", dest);
}

int _getline(char *dest)
{
    char c;
    int i;
    for (i = 0; i < MAX_LINE_LEN && (c = getchar()) != EOF && c != '\n'; i++)
    {
        *(dest + i) = c;
    }
    *(dest + i) = '\0';
    return i;
}