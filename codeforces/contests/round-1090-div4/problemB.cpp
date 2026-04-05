#include <bits/stdc++.h>

using namespace std;

int main() {
    int t;
    cin >> t;
    while (t--){
        int mx = -68;
        int sum = 0;
        for (int i=0; i < 7; i++) {
            int tmp; 
            cin >> tmp;
            mx = max(mx, tmp);
            sum -= tmp;
        }
        cout << sum+(2*mx) << "\n";
    }
}
