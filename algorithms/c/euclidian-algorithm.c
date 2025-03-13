#include <stdio.h>

int main() {
    int a, b;
    scanf("%d %d", &a, &b);
    //Make sure a is greater or equal to b
    if (b > a) {
        int tmp = a;
        a = b;
        b = tmp;
    }
    //Run euclidian algorithm
    int mod;
    while (a != 0 && b != 0) {
        mod = a%b;
        a = b;
        b = mod;
    }
    printf("\nGreatest common divider is %d\n", a+b);
}
