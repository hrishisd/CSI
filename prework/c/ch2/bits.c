#include <stdio.h>

void print_bits(unsigned int x)
{
  if (x == 0) 
  {
    putchar('0');
  }
  else 
  {
    print_bits(x >> 1);
    printf("%d", x & 1);
  }
}

// rotate right by n bits, wrapping around
unsigned rotate_right(unsigned x, int n)
{
  unsigned last_n = x & ~(~0 << n); 
  return x >> n | last_n << ((sizeof(unsigned) * 4) - n);
}

// Get n bits from position p to (p-n+1)
// p = 0 means rightmost bit, p = 4 means 4th from left (0-indexed)
unsigned getbits(unsigned x, int p, int n)
{
  x = x >> (p - n + 1); // shift so that the last n bits are the desired bits
  unsigned mask = ~(~0 << n);
  return x & mask;
}

int main()
{
  // print_bits((-1) << 5);
  unsigned x = 12345;
  print_bits(x);
  printf("\n");
  print_bits(rotate_right(x, 2));
  printf("\n");
  print_bits(getbits(x, 6, 5));
  printf("\n");
  //print_bits(4);
  //printf("\n");
}
