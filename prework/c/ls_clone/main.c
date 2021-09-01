#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>
#include <dirent.h> // for opendir and dirent definition
#include <sys/stat.h>
#include <unistd.h> // for getopt

static const int MAX_FILES_IN_DIR = 1024;

typedef struct
{
    char *name;
    struct stat *_stat;
} file_metadata;

static file_metadata files[MAX_FILES_IN_DIR];

size_t read_files(DIR *parent_dir);
void print_files(size_t num_files);

int main(int argc, char *argv[])
{
    bool display_hidden = false;
    bool long_form = false;
    int opt;
    enum
    {
        CHARACTER_MODE,
        WORD_MODE,
        LINE_MODE
    } mode = CHARACTER_MODE;

    while ((opt = getopt(argc, argv, "hl")) != -1)
    {
        switch (opt)
        {
        case 'h':
            display_hidden = true;
            break;
        case 'l':
            long_form = true;
            break;
        default:
            fprintf(stderr, "Usage: %s [-hl] [file...]\n", argv[0]);
            exit(EXIT_FAILURE);
        }
    }
    // Now optind (declared extern int by <unistd.h>) is the index of the first non-option argument.
    // If it is >= argc, there were no non-option arguments.
    bool no_arguments = optind >= argc;
    int num_root_files; // number of files to list / directories to list contents of
    char **root_files;  // an array of root file names; the length of the array is num_root_files
    if (no_arguments)
    {
        char *current_dir = ".";
        root_files = &current_dir;
        num_root_files = 1;
    }
    else
    {
        root_files = argv + optind;
        num_root_files = argc - optind;
    }

    // display_files(root_files, num_root_files, display_hidden, long_form);

    char *dir_name = root_files[0];

    // char *dir_name = argc < 2 ? "." : argv[1];
    DIR *dir = opendir(dir_name);
    if (dir == NULL)
    {
        fprintf(stderr, "Unable to read directory: %s\n", dir_name);
        return 1;
    }
    size_t num_files = read_files(dir);
    free(dir);

    print_files(num_files);

    // free allocated stat structs
    for (int i = 0; i < num_files; i++)
    {
        free(files[i]._stat);
    }
}

// Populates `files` array with the file metadata of all files and subdirectories in `parent_dir`
// Returns the number of files and subdirectories inside `parent_dir`.
size_t read_files(DIR *parent_dir)
{
    struct dirent *entry = malloc(sizeof(*entry));
    size_t idx = 0;
    for (; (entry = readdir(parent_dir)) != NULL; idx++)
    {
        struct stat *_stat = malloc(sizeof(*_stat));
        stat(entry->d_name, _stat);
        files[idx].name = entry->d_name;
        files[idx]._stat = _stat;
    }
    free(entry);
    return idx;
}

void print_files(size_t num_files)
{
    for (int i = 0; i < num_files; i++)
    {
        file_metadata file = files[i];
        printf("%s\n", file.name);
    }
}