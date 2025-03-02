#include <bits/stdc++.h>

using namespace std;

int main() {
    int t;
    cin >> t;
    while (t--) {
        long long n;
        string s;
        cin >> n >> s;
        long long underscore = 0, dash = 0;
        for (int i = 0; i < n; i++) {
            if (s[i] == '_') {
                underscore++;
            } else {
                dash++;
            }
        }
        if (n <= 2 || dash < 2 || underscore < 1) {
            printf("0\n");
        } else {
            cout << (dash/2)*((dash+1)/2)*underscore << "\n";
        }
    }
}
