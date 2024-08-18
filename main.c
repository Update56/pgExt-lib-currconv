#include <stdlib.h>
#include <stdio.h>
#include <dlfcn.h>

int main() {
    void *handle;
    char* currency = "USD";
    handle = dlopen ("libgetconv.so", RTLD_NOW | RTLD_GLOBAL);
    if (!handle) {
        fputs (dlerror(), stderr);
        exit(1);
    }

    float rates[3];
    char*(*getbody)(char*) = dlsym(handle, "GetBody");
    char* body = getbody(currency);

    printf("%s", body);
    return 0;
}