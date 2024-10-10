#include <bits/stdc++.h>

using namespace std;

vector<int> quick_sort(vector<int> a) {
    int pivot = a.size() / 2;
    vector<int> b, c;
    for (int i = 0; i < a.size(); i++) {
        if (i < pivot) {
            b.push_back(a[i]);
        } else {
            c.push_back(a[i]);
        }
    }
    b = quick_sort(b);
    c = quick_sort(c);
    
}

int main() {
    int n;
    cin >> n;

    vector<int> a(n);

    for(int i = 0; i < n; i++) {
        cin >> a[i];
    }

    a = quick_sort(a);

    for(int i = 0; i < n; i++) {
        cout << a[i] << " ";
    }
    cout << "\n";
}   