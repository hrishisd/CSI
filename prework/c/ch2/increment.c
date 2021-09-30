#include <stdio.h>

#define DOUBLE(X) X * X

int main()
{
  int x = 5;
  printf("%d, %d\n", x++, x++);
  printf("%d\n", ++x);
  printf("%d\n", DOUBLE(x));
}
