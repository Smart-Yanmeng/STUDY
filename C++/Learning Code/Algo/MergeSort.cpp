#include "iostream"

using namespace std;

void Merge(int arr[], int t[], int low, int mid, int high) {
    int i = low, j = mid + 1, k = low;
    while (i <= mid && j <= high) {//将L.r中记录从小到大存入a
        if (arr[i] < arr[j])
            t[k++] = arr[i++];
        else
            t[k++] = arr[j++];
    }
    while (i <= mid)
        t[k++] = arr[i++];//将剩余的L.r[i..mid]存入a
    while (j <= high)
        t[k++] = arr[j++];//将剩余的L.r[j..high]存入a

    while (low <= high) {
        arr[low] = t[low];
        low++;
    }
}

void PA(int arr[], int n) {
    for (int i = 0; i < n - 1; i++) cout << arr[i] << " ";
    cout << arr[n - 1] << endl;
}

void MS(int arr[], int t[], int low, int high, int n) {
    //L.r[low..high]归并排序后存入L.r[low..high]中
    if (low < high) {
        int mid = (low + high) / 2;//将当前序列一分为二
        MS(arr, t, low, mid, n);//将序列L.r[low..mid]递归归并排序结果放入L.r[low,mid]
        MS(arr, t, mid + 1, high, n);//将L.r[mid+1..high]递归归并排序结果放入L.r[mid+1..high]
        Merge(arr, t, low, mid, high);//将L.r[low..mid]和L.r[mid+1,high]归并到L.r[low..high]
        PA(arr, n);
    }
}

int main() {
    int n;
    while (cin >> n) {
        int arr[n];
        for (int i = 0; i < n; i++) cin >> arr[i];
        int t[n];
        MS(arr, t, 0, n - 1, n);
    }

    return 0;
}