#include <stdlib.h>
#include <stdio.h>
#include <dlfcn.h>

int main() {
    void *handle;
    handle = dlopen ("./libgetconv.so", RTLD_NOW | RTLD_GLOBAL);
    if (!handle) {
        fputs (dlerror(), stderr);
        exit(1);
    }

    float rates[3];
    void(*getconv)(float[]) = dlsym(handle, "Getconv");
    getconv(rates);

    printf("EUR rate: %f\n", rates[0]);
    printf("JPY rate: %f\n", rates[1]);
    printf("USD rate: %f\n", rates[2]);
    return 0;
}