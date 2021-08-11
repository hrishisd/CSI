#include <stdio.h>

int power(int, int);

int main() 
{
  printf("2^5 = %d", power(2, 5));
}

int power(int base, int exp) 
{
  int res = 1;
  for (int i = 0; i < exp; i++)
  {
    res *= base;
  }
  return res;
}
