#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <time.h>
#include <unistd.h>   // POSIX api
#include <sys/stat.h> // stat syscall
#include <sys/types.h>
#include <sys/wait.h> // wait syscall
#include <string.h>
#include <errno.h>

size_t MAX_LINE_CHARS = 200;

bool file_exists(char *filename)
{
    struct stat buffer;
    return (stat(filename, &buffer) == 0);
}

int main()
{
    char *line_buffer = malloc(MAX_LINE_CHARS * sizeof(char));
    if (line_buffer == NULL)
    {
        //perror("Unable to allocate buffer");
        exit(1);
    }
    size_t chars_read;
    while (true)
    {
        printf("λ ");
        /*
        chars_read = getline(&line_buffer, &MAX_LINE_CHARS, stdin);
        if (chars_read == -1) // EOF
            break;
	*/
        pid_t pid = fork();
        if (pid == 0)
        {

            char *command = "whoami";
            //printf("hello from child\n");
            char *env = getenv("PATH");
            //printf("child path env: %s\n", env);
            char *argv[] = {command, NULL};
            int result = execvp(command, argv);
            printf("result from child is %d\n", result);
            printf("errno: %d\n", errno);
            exit(0);
        }
        else
        {
            //printf("hello from parent. Child pid is %d\n", pid);
            int status;
            waitpid(pid, &status, 0);
        }
        break;
    }
}
