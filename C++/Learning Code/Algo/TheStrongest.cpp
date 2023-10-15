#include "iostream"

using namespace std;

int main() {
    int arr[10], res = 0;
    for (int &i: arr) cin >> i;
    for (int i = 0; i < 10; i++) {
        if (i == 0) {
            if (arr[i] > arr[i + 1]) res += arr[i];
        } else if (i == 9) {
            if (arr[i] > arr[i - 1]) res += arr[i];
        } else {
            if (arr[i] > arr[i - 1] && arr[i] > arr[i + 1]) res += arr[i];
        }
    }

    cout << res * res;

    return 0;
}