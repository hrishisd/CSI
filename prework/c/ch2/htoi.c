#include <stdio.h>

int htoi(char[]);
int length(char s[]);
int hex_digit_to_int(char hex);

const int MAX_STRING_LEN = 1024;

int main() 
{
  char hex[] = {'0', 'x', '1', '2', '3', '\0'};
  printf("%s == %d\n", hex, htoi(hex));
}

int length(char s[])
{
  int len = 0;
  for (;len < MAX_STRING_LEN ; len++)
  {
    if (s[len] == '\0')
      return len;
  }
  return MAX_STRING_LEN;
}

int htoi(char hex[]) 
{
  int result = 0;
  int len = length(hex);
  printf("length is %d\n",  len);
  int start_idx = 
    len > 2 && hex[0] == '0' && (hex[1] == 'x' || hex[1] == 'X')
    ? 2
    : 0;

  for (int i = start_idx ; i < len ; i++)
  {
    result = result * 16 + hex_digit_to_int(hex[i]);
  }
  return result;
}

int hex_digit_to_int(char hex)
{
  if ('0' <= hex && hex <= '9')
  {
    return hex - '0';
  }
  else if ('a' <= hex && hex <= 'f')
  {
    return hex - 'a';
  }
  else 
  {
    return hex - 'A';
  }
}
