#include "iostream"

using namespace std;

void PA(int arr[], int n) {
    for (int i = 0; i < n - 1; i++) cout << arr[i] << " ";
    cout << arr[n - 1] << endl;
}

void QS(int arr[], int l, int r, int n) {
    if (l >= r) return;

    int i = l, j = r;
    int pivot = arr[l];
    while (i != j) {
        // 从右往左
        while (arr[j] >= pivot && j > i) j--;
        arr[i] = arr[j];
        arr[j] = pivot;

        // 从左往右
        while (arr[i] <= pivot && j > i) i++;
        arr[j] = arr[i];
        arr[i] = pivot;
    }
    PA(arr, n);

    QS(arr, l, i - 1, n);
    QS(arr, i + 1, r, n);
}

int main() {
    int n;
    while (cin >> n) {
        int arr[n];
        for (int i = 0; i < n; i++) cin >> arr[i];
        QS(arr, 0, n - 1, n);
    }
}