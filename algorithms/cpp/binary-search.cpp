#include "bits/stdc++.h"

using namespace std;

int main() {
  // Binary Search
  vector<int> a = {1, 3, 5, 8, 12};
  int l = 0, r = a.size() - 1, m = 0, want = 8;
  while (l <= r) {
    m = (l + r) / 2;
    if (a[m] > want) {
      r = m - 1;
    } else if (a[m] < want) {
      l = m + 1;
    } else {
      printf("Index of the wanted number is %d\n", m);
      return 0;
    }
  }
  printf("Element is not in the array\n");
}
