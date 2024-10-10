#include <bits/stdc++.h>

using namespace std;

vector<int> merge_sort(vector<int> a) {
    if (a.size() == 1) {
        return a;
    }
    int m = a.size()/2;
    vector<int> b;
    vector<int> c;
    for (int i = 0; i < a.size(); i++) {
        if (i < m) {
            b.push_back(a[i]);
        } else {
            c.push_back(a[i]);
        }
    }
    b = merge_sort(b);
    c = merge_sort(c);
    vector<int> answer;
    int i = 0, j = 0;
    while (i < b.size() && j < c.size()) {
        if (b[i] < c[j]) {
           answer.push_back(b[i]);
           i++;
        } else {
            answer.push_back(c[j]);
            j++;
        }
    }
    if (j == a.size()) {
        for (int k = j; k < c.size(); k++) {
            answer.push_back(c[k]);
        }
        return answer;
    }
    for (int k = i; k < b.size(); k++) {
            answer.push_back(b[k]);
    }
    return answer;
}

int main() {
    int n;
    cin >> n;

    vector<int> a(n);

    for(int i = 0; i < n; i++) {
        cin >> a[i];
    }

    a = merge_sort(a);

    for(int i = 0; i < n; i++) {
        cout << a[i] << " ";
    }
    cout << "\n";
}