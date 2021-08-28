#include <stdio.h>
#define BUF_SIZE 10000

static char allocbuf[BUF_SIZE];
static int free_idx = 0;

char *alloc(int n)
{
    if (free_idx + n > BUF_SIZE)
    {
        return NULL;
    }
    char *res = &allocbuf[free_idx];
    free_idx += n;
    return res;
}

char *afree(char *p)
{
    free_idx = p - allocbuf;
}