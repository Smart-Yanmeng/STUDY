#include "iostream"

using namespace std;

int main() {
    int n;
    cin >> n;

    int arr[n], resArr[n - 1];

    for (int i = 0; i < n; i++) {
        cin >> arr[i];

        if (i > 0) {
            resArr[i - 1] = arr[i] - arr[i - 1];
        }
    }

    for (int i = 0; i < n - 1; i++) {
        if (i % 3 == 0 && i != 0) cout << endl << resArr[i];
        else cout << resArr[i];

        if (i % 3 < 2 && i != n - 2) cout << " ";
    }

    return 0;
}