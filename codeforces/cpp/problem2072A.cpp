#include <bits/stdc++.h>

using namespace std;

int main() {
    int t;
    cin >> t;
    while (t--) {
        int n, k, p;
        cin >> n >> k >> p;
        if (abs(k)>abs(n*p)) {
            printf("-1\n");
            continue; 
        } else {
            printf("%d\n", (abs(k)+abs(p)-1)/abs(p));
        }

    }
}
