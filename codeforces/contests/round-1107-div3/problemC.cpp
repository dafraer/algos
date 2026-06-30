#include <bits/stdc++.h>

using namespace std;

int main() {
    int t;
    cin >> t;
    while (t--){
        int n;
        string s;
        cin >> n;
        cin >> s;
        int cnt0 = 0;
        int cnt1 = 0; 
        for (int i = 0; i < n; i++) {
            if (s[0]) {
                cnt1++;
            } else {
                cnt0++;
            }
        }
        printf("%d\n", (abs(cnt0-cnt1)%2)+1);      
    }
}