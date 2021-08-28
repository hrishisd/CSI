#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define SIZE 100

struct Entry
{
    struct Entry *next;
    char *key;
    char *value;
};

static struct Entry *table[SIZE];

unsigned hash(char *s)
{
    unsigned result = 0;
    while (*s != '\0')
    {
        result = *s + 31 * result;
        s++;
    }
    return result % SIZE;
}

struct Entry *lookup(char *key)
{
    for (struct Entry *entry = table[hash(key)]; entry != NULL; entry = entry->next)
        if (strcmp(entry->key, key) == 0)
            return entry;
    return NULL;
}

struct Entry *put(char *key, char *value)
{
    struct Entry *existing_entry = lookup(key);
    if (existing_entry != NULL)
    {
        free(existing_entry->value);
        existing_entry->value = value;
        return existing_entry;
    }
    else
    {
        struct Entry *new_entry = malloc(sizeof(struct Entry));
        if (new_entry != NULL)
        {
            new_entry->key = key;
            new_entry->value = value;
            new_entry->next = table[hash(key)];
            return new_entry;
        }
        else
        {
            fprintf(stderr, "Failed to allocate for new entry.");
            return NULL;
        }
    }
}

int main(void)
{
}