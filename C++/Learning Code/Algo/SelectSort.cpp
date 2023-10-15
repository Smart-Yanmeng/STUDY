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

void SelectSort(int arr[], int n) {
    for (int i = 0; i < n - 1; i++) {
        int k = i;
        for (int j = i + 1; j < n; j++) {
            if (arr[j] < arr[k]) k = j;
        }
        if (k != i) {
            warp(arr[i], arr[k]);
            PrintArr(arr, n);
        }
    }
}

int main() {
    int n;
    cin >> n;
    int arr[n];
    for (int i = 0; i < n; i++) cin >> arr[i];

    SelectSort(arr, n);

    return 0;
}