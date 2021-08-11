#define MAX_LINE_SIZE 1000

#include <stdio.h>

int get_line(char dest[]);
void print_str(char s[]);
void copy_str(char src[], char dest[]);
int length(char string[]);
void reverse(char s[]);

int main()
{
  char line[MAX_LINE_SIZE];
  char longest[MAX_LINE_SIZE];
  int max_size = 0;
  int size = 0;
  while ((size = get_line(line)) != 0) 
  {
    if (size > max_size)
    {
      max_size = size;
      copy_str(line, longest);
    }
  }
  reverse(longest);
  printf("%s\n", longest);
}

void print_str(char s[])
{
  char c;
  for (int i = 0; i < MAX_LINE_SIZE && (c = s[i]) != '\0'; i++)
  {
    putchar(s[i]);
  }
  putchar('\n');
}

// populates dest array with next line
// returns number of chars in array
int get_line(char dest[])
{
  int i = 0;
  char c;
  for (; i < MAX_LINE_SIZE-1 && (c = getchar()) != EOF && c != '\n'; i++) 
  {
    dest[i] = c;
  }
  dest[i] = '\0';
  return i;
}

void copy_str(char src[], char dest[]) 
{
  for (int i = 0; i < MAX_LINE_SIZE; i++)
  {
    char c = src[i];
    dest[i] = c;
    if (c == '\0')
      return;
  }
}

int length(char string[])
{
  int length = 0;
  for (int i = 0; i < MAX_LINE_SIZE && string[i] != '\0'; i++)
    length++;
  return length;
}

void reverse(char s[]) 
{
  int len = length(s);
  for (int i = 0; i < len/2; i++)
  {
    char temp = s[i];
    s[i] = s[len-i-1];
    s[len-i-1] = temp;
  }
}
