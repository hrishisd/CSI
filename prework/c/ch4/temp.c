#include <stdio.h>

#define double(X) X * X

int main(void)
{
  printf("%d", double(1+2));
}
