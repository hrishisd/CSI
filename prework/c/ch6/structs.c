#include <stdio.h>

struct point
{
    int x;
    int y;
};

int magnitude(struct point p)
{
    return p.x * p.x + p.y * p.y;
}

int main(void)
{
    struct point p = {1, 2};
    printf("(%d, %d)\n", p.x, p.y);
    printf("%d\n", magnitude(p));
}