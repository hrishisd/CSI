#include <stdio.h>

// returns the ith month's name as a string
char *month_name(int i)
{
    static char *months[] = {"jan", "feb", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"};
    return *(months + i);
}

int main(void)
{
    printf("%s\n", month_name(0));
}