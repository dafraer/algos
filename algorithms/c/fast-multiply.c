#include <stdio.h>

long int product (int x, int y) {
    long int z = 0;
    while (y) {
        if (y % 2) {
            y--;
            z += x;
        } else {
            y>>=1;
            x <<=1;
        }
    }
    return z;
}

int main() {
    int a, b;
    a = 29;
    b = 38;
    printf("%d", product(a,b));
}
