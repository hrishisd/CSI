#include <stdio.h>

enum boolean {NO='N', YES='y'};

int main()
{
  const enum boolean x = NO;
  printf("%c", x);
  x = YES;
  printf("%c", x);
}
