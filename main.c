#include <stdio.h>
#include <stdlib.h>
#include "libgetconv.h"

int main() {
    char* rates = Getconv();
    printf("Rates:\n%s\n", rates);
    free(rates);
    return 0;
}