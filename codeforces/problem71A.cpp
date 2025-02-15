#include <bits/stdc++.h>

using namespace std;

int main() {
    int n;
    cin >> n;
    string s;
    for (int i = 0; i < n; i++) {
        cin.getline(&s);
        if (s.size() <= 10) {
            printf("%s\n", s);
        }
        s = s[0] + to_string(s.size()) + s[s.size()];
        cout << s << "\n";
    }
}