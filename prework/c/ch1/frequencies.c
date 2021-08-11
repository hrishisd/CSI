#include <stdio.h>

int main()
{
  int digit_freqs[10] = {0, 0, 0, 0, 0, 0, 0, 0, 0, 0};
  int whitespaces = 0;
  int other = 0;

  char c;
  while ((c = getchar()) != EOF) 
  {
    if (c >= '0' && c <= '9')
    {
      digit_freqs[c - '0']++;
    } 
    else if (c == ' ' || c == '\t' || c == '\n') 
    {
      whitespaces++;
    }
    else 
    {
      other++;
    }
  }

  putchar('\n');

  int max_freq = 0;
  for (int i = 0; i < 10; i++)
  {
    if (max_freq < digit_freqs[i])
    {
      max_freq = digit_freqs[i];
    }
  }
  printf("max: %d\n\n\n\n", max_freq);


  for (int freq = max_freq; freq > 0; freq--) {
    for (int i = 0; i < 10; i++)
    {
      if (digit_freqs[i] >= freq) 
       putchar('*'); 
      else 
        putchar(' ');
    }
    printf("\n");
  }

  for (int i = 0; i < 10; i++)
    printf("%d", i);

  putchar('\n');

  printf("\ndigit frequencies:\n");
  for (int i = 0; i < 10; i++) 
    printf("%d: %d\n", i, digit_freqs[i]);
  printf("whitespaces: %d\nother: %d\n", whitespaces, other);
}
