#include "iostream"

using namespace std;

int main() {
//    int n;
//    cin >> n;
//
//    int arr[n];
//
//    for (int i = 0; i < n; i++) {
//        cin >> arr[i];
//    }
//
//    for (int i = n - 1; i >= 0; i--) {
//        if (i == n - 1) cout << arr[i];
//        else cout << " " << arr[i];
//    }

    int a;
    int i = 0, arr[1000];

    while (~scanf("%d", &a)) {
        if (a == -1) break;
        else arr[i++] = a;
    }

    for (int j = i - 1; j >= 0; j--) {
        if (j == 0) cout << arr[j];
        else cout << arr[j] << " ";
    }

    return 0;
}