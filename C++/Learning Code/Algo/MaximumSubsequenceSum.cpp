#include "iostream"

using namespace std;

int MaximumSubsequenceSum(int arr[], int n) {
    int maxSum = 0, tmpSum = 0;
    for (int i = 0; i < n; i++) {
        tmpSum += arr[i];
        if (tmpSum < 0) tmpSum = 0;
        if (tmpSum > maxSum) maxSum = tmpSum;
    }
    return maxSum;
}

int main() {
    int n;
    cin >> n;
    int arr[n];
    for (int i = 0; i < n; i++) cin >> arr[i];

    cout << MaximumSubsequenceSum(arr, n);

    return 0;
}