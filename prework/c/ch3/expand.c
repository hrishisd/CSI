#include <stdio.h>
#include <string.h>

void expand(char shorthand[], char result[])
{
    size_t size = strlen(shorthand);
    int insert_idx = 0;
    // Look two characters back to see if we need to expand.
    int i = 2;
    for (; i < size; i++)
    {
        if (shorthand[i - 1] == '-')
        {
            // Insert all expanded chars into result.
            for (char c = shorthand[i - 2]; c < shorthand[i]; c++)
            {
                result[insert_idx++] = c;
            }
            i += 1;
        }
        else
        {
            result[insert_idx++] = shorthand[i - 2];
        }
    }

    // Insert last two characters if we didn't end on an expansion.
    if (i - 2 < size)
    {
        result[insert_idx++] = shorthand[i - 2];
    }
    if (i - 1 < size)
    {
        result[insert_idx++] = shorthand[i - 1];
    }
    result[insert_idx] = '\0';
}

int main()
{
    char result[1000];
    char shorthand[] = "abzs";
    expand(shorthand, result);
    printf("%s -> %s\n", shorthand, result);

    char shorthand2[] = "a-z";
    expand(shorthand2, result);
    printf("%s -> %s\n", shorthand2, result);

    char shorthand3[] = "a-b-z";
    expand(shorthand3, result);
    printf("%s -> %s\n", shorthand3, result);

    char shorthand4[] = "-a-z-";
    expand(shorthand4, result);
    printf("%s -> %s\n", shorthand4, result);
}