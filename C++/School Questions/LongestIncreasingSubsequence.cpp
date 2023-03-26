#include "iostream"

using namespace std;

int main() {
    int n, increasingNum = 0, maxIncreasingNum = 0, ss = 0, x = 0, y = 0;
    cin >> n;
    int arr[n];

    for (int i = 0; i < n; i++) cin >> arr[i];

    for (int i = 0; i < n - 1; i++) {
        y = i + 1;
        if (arr[i + 1] > arr[i]) {
            increasingNum++;

            if (increasingNum > maxIncreasingNum) {
                maxIncreasingNum = increasingNum;
                ss = x;
            }
        } else {
            increasingNum = 0;
            x = y;
        }
    }

    for (int i = ss; i <= ss + maxIncreasingNum; i++) {
        cout << arr[i];
        if (i < ss + maxIncreasingNum) cout << " ";
    }

    return 0;
}