#include <bits/stdc++.h>
using namespace std;


int main() {
    int t;
    cin >> t;
    while (t--){
        int n;
        cin >> n;
        unsigned long long k = 1;
        while (n--) {
            printf("%llu ", k * (k+2));
            k+=2;
        }
        printf("\n");
    }
}
