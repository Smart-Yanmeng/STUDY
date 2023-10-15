#include "iostream"

using namespace std;

void warp(int &a, int &b) {
    int tmp = a;
    a = b;
    b = tmp;
}

void PrintArr(int arr[], int n) {
    for (int i = 0; i < n - 1; i++) cout << arr[i] << " ";
    cout << arr[n - 1] << endl;
}

void BubbleSort(int arr[], int n) {
    int flag = 0;
    for (int i = 0; i < n - 1; i++) {
        for (int j = 0; j < n - i - 1; j++) {
            if (arr[j] > arr[j + 1]) {
                warp(arr[j], arr[j + 1]);
                flag = 1;
            }
        }
        if (flag) PrintArr(arr, n);
        flag = 0;
    }
}

int main() {
    int n;
    cin >> n;
    int arr[n];
    for (int i = 0; i < n; i++) cin >> arr[i];

    BubbleSort(arr, n);

    return 0;
}