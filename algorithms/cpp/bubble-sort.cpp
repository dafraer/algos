#include <bits/stdc++.h>

using namespace std;

vector<int> bubble_sort(vector<int> a) {
  for (int i = 0; i < a.size(); i++) {
    for (int j = 0; j < a.size() - 1; j++) {
      if (a[j] > a[j + 1]) {
        swap(a[j], a[j + 1]);
      }
    }
  }
  return a;
}

int main() {
  // input number of elements
  int n;
  cin >> n;
  // declare vector of size n
  vector<int> a(n);
  for (int i = 0; i < n; i++) {
    cin >> a[i];
  }

  // sort using bubble sort
  a = bubble_sort(a);
  for (int i = 0; i < n; i++) {
    cout << a[i] << " ";
  }
  cout << "\n";
}
