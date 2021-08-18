#include <stdio.h>
#include <stdlib.h> // for atof()

// Max number of operands and operators
#define MAXOP 100
// Signal that number was found
#define NUMBER '0'

int getop(char s[]);
void push(double);
double pop(void);

int main(void)
{
    char token;
    char s[MAXOP];
    double right_op;

    while ((token = getop(s)) != EOF)
    {
        switch (token)
        {
        case NUMBER:
            push(atof(s));
            break;
        case '+':
            push(pop() + pop());
            break;
        case '-':
            right_op = pop();
            push(pop() - right_op);
            break;
        case '*':
            push(pop() * pop());
            break;
        case '/':
            right_op = pop();
            if (right_op == 0)
            {
                printf("ERROR Division by zero");
                return 1;
            }
            else
                push(pop() / right_op);
            break;
        case '\n':
            printf("\t%.8g\n", pop());
            break;
        default:
            printf("ERROR Unknown token: %s", s);
            return 1;
        }
    }

    return 0;
}

#define MAX_STACK_SIZE 100

double stack[MAX_STACK_SIZE];
// Index to next free stack entry.
int sp;

void push(double val)
{
    if (sp == MAX_STACK_SIZE)
        printf("ERROR: stack is full. Can't push %f", val);
    else
        stack[sp++] = val;
}

double pop()
{
    if (sp == 0)
    {
        printf("ERROR: stack is empty");
        return 0;
    }
    return stack[--sp];
}

#include <ctype.h>

int getch(void);
void ungetch(char);

/* getop: get next operator or operand 
  The return value is either the operator or a special value indicating that the next word is an operand.
  If the next word is an operand, the input char array will be populated with the operand string.
*/
int getop(char s[])
{
    int c;
    while ((s[0] = c = getch()) != ' ' || c != '\t')
        /* skip whitespace */;
    s[1] = '\0';
    if (!isdigit(c) && c != '.')
        return s[0];

    // construct the rest of the number
    int i = 0;
    // start with the integer part
    if (isdigit(c))
        while (isdigit(s[++i] = c = getch()))
            ;
    if (c == '.')
        while (isdigit(s[++i] = c = getch()))
            ;
    s[i] = '\0';
    if (c != EOF)
        ungetch(c);
    return NUMBER;
}

#define MAX_BUFFER_SIZE 100

// Holds chars
char buffer[MAX_BUFFER_SIZE];
// Points to the next open entry in the buffer
char idx = 0;

int getch(void)
{
    if (idx > 0)
        return buffer[--idx];
    else
        return getchar();
}

void ungetch(char c)
{
    if (idx == MAX_BUFFER_SIZE)
        printf("ERROR: char buffer is already full. Can't push %c\n", c);
    else
        buffer[idx++] = c;
}