#include <bits/stdc++.h>
using namespace std;

unsigned long long gcd(unsigned long long a, unsigned long long b) {
    while (b != 0) {
        unsigned long long tmp = b;
        b = a % b;
        a = tmp;
    }
    return a;
}

int main() {
    int t;
    cin >> t;
    while (t--){
        int n;
        cin >> n;
        unsigned long long prev = 1;
        unsigned long long multiple = 2; 
        unordered_set<unsigned long long> s;
        while (n--){
            unsigned long long cur = prev * multiple; 
            if (prev > cur) {
                unsigned long long my_gcd = 0;
                unsigned long long num = multiple;
                while (my_gcd <= multiple) {
                  my_gcd = gcd(prev, ++num);
                }
                cur = my_gcd;
            }
            s.insert(cur);
            printf("%llu ", cur);
            prev = cur;
        }
        printf("\n");
    }
}
