#include <stdio.h>
#include <time.h>

int binsearch_book(int x, int v[], int n)
{
    int lo = 0;
    int hi = n - 1;
    while (lo <= hi)
    {
        int mid = (lo + hi) / 2;
        // printf("%d %d %d\n", lo, mid, hi);
        if (x > v[mid])
            lo = mid + 1;
        else if (x < v[mid])
            hi = mid - 1;
        else
            return mid;
    }
    return -1;
}

int binsearch_mine(int x, int v[], int n)
{
    int lo = 0;
    int hi = n - 1;
    while (lo <= hi)
    {
        int mid = (lo + hi) / 2;
        // printf("%d %d %d\n", lo, mid, hi);
        if (x > v[mid])
            lo = mid + 1;
        else
            hi = mid - 1;
    }
    if (x == v[lo])
        return lo;
    else if (x == v[hi])
        return hi;
    return -1;
}

clock_t time_binsearch(int (*f)(int, int[], int), int target, int arr[], int size)
{
    clock_t start = clock();
    f(target, arr, size);
    clock_t end = clock();
    return end - start;
}

int main()
{
    // The benchmark is a fail, unfortunately.
    const int size = 2000000;
    int bigarr[size];
    for (int i = 0; i < size; i++)
    {
        bigarr[i] = i;
    }
    printf("%d\n", binsearch_book(3, bigarr, size));
    printf("%d\n", binsearch_mine(3, bigarr, size));

    printf("book sort:\n");
    for (int i = 0; i < 5; i++)
    {
        printf("%lu\n", time_binsearch(binsearch_book, i, bigarr, size));
    }

    printf("my sort:\n");
    for (int i = 0; i < 5; i++)
    {
        printf("%lu\n", time_binsearch(binsearch_mine, i, bigarr, size));
    }
}