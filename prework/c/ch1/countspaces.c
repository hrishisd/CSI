#include <stdio.h>

int main()
{
  char c;
  long whitespace_chars = 0;
  while ((c = getchar()) != EOF) 
  {
    if (c == '\t' || c == '\\')
    {
       printf("\\");
    } else {
      putchar(c);
    }
  }
}
