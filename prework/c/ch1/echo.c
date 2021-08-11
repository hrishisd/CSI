#include <stdio.h>

int main() 
{
  printf("EOF: '%c'", EOF);

  char c;
  while ((c = getchar()) != EOF) 
  {
      putchar(c);
  }
}
