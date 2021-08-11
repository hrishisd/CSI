#include <stdio.h>
#include <string.h>
#include <stdbool.h>

// Deletes all characters from s1 that appear in s2
void squeeze(char s1[], char s2[]);
bool contains(char s[], char c);

int main()
{
  char s[] = "hello world";
  char exclude[] = "abcde";
  squeeze(s, exclude);
  printf("%s", s);
}

void squeeze(char s1[], char s2[])
{
  int len = strlen(s1);
  int insert_idx = 0;
  for (int i = 0; i < len; i++)
  {
    char c = s1[i];
    if (!contains(s2, c))
    {
      s1[insert_idx] = c;
      insert_idx++;
    }
  }
  s1[insert_idx] = '\0';
}

bool contains(char s[], char c)
{
  size_t len = strlen(s);
  for (int i = 0; i < len; i++)
    if (s[i] == c)
      return true;

  return false;
}

