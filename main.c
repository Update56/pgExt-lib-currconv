#include <stdio.h>
#include <stdlib.h>
#include "libgetconv.h"

int main() {
    char *currency = "USD";

    // Получаем строку от Go
    char* body = GetBody(currency);
    
    // Печатаем строку
    printf("Body: %s\n", body);
    
    // Освобождаем память, выделенную в Go
    free(body);
    
    return 0;
}