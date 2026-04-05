#include <bits/stdc++.h>

using namespace std;

int main() {
    int t;
    cin >> t;
    while (t--){
        int n;
        cin >> n;
        n*=3;
        vector<int>a(n);
        for (int i=n-1; i >= 2; i-=3) { 
            a[i] = n;
            a[i-1] = n-1;
            n-=2;
        }
        for (int i=0; i < a.size(); i+=3) { 
            a[i] = n;
            n-=1;
        } 
        for (int i = 0; i < a.size(); i++) {
            cout << a[i] << " ";
        }
        cout << "\n";
    }
}
