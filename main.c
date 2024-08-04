#include <stdio.h>
#include "libgetconv.h"

int main() {
    float rates[3];
    Getconv(rates);
    printf("EUR rate: %f\n", rates[0]);
    printf("JPY rate: %f\n", rates[1]);
    printf("USD rate: %f\n", rates[2]);
    return 0;
}