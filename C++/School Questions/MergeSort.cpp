#include<iostream>

using namespace std;

int arr[100];
int n;

void Merge(int arr[], int low, int mid, int high) {
    int i = low, j = mid + 1, k = 0;
    int *temp = new(nothrow) int[high - low + 1];

    if (!temp) {
        cout << "ERROR!";

        return;
    }

    while (i <= mid && j <= high) {
        if (arr[i] <= arr[j])
            temp[k++] = arr[i++];
        else
            temp[k++] = arr[j++];
    }
    while (i <= mid)
        temp[k++] = arr[i++];
    while (j <= high)
        temp[k++] = arr[j++];
    for (i = low, k = 0; i <= high; i++, k++)
        arr[i] = temp[k];
    delete[]temp;
}

void MergeSort(int arr[], int low, int high) {
    if (low < high) {
        int mid = (low + high) / 2;
        MergeSort(arr, low, mid);
        MergeSort(arr, mid + 1, high);
        Merge(arr, low, mid, high);

        for (int i = 0; i < n; i++) {
            if (i == n - 1) cout << arr[i] << endl;
            else cout << arr[i] << " ";
        }
    }
}

int main() {
    while (~scanf("%d", &n)) {
        for (int i = 0; i < n; i++) cin >> arr[i];
        MergeSort(arr, 0, n - 1);
    }

    return 0;
}
//
//#include<iostream>
//
//using namespace std;
//
//int a[110];
//int n;
//
//void merge(int a[], int tem[], int p, int q, int r) {
//
//    int i = p;
//    int j = q + 1;
//    int pos = p;
//
//    while (i <= q && j <= r) {
//        if (a[i] < a[j])
//            tem[pos++] = a[i++];
//        else
//            tem[pos++] = a[j++];
//    }
//    while (i <= q)
//        tem[pos++] = a[i++];
//    while (j <= r)
//        tem[pos++] = a[j++];
//    while (p <= r) {
//        a[p] = tem[p];
//        p++;
//    }
//}
//
//void print(int s[], int n) {
//    int t = 0;
//    for (t = 0; t < n; t++) {
//        cout << s[t];
//
//        if (t == n - 1) cout << endl;
//        else cout << " ";
//    }
//}
//
//void msort(int a[], int tem[], int p, int r) {
//    if (p < r) {
//        int q = (p + r) / 2;
//        msort(a, tem, p, q);
//        msort(a, tem, q + 1, r);
//        merge(a, tem, p, q, r);
//
//        print(a, n);
//    }
//}
//
//void merge_sort(int a[], int n) {
//    int tem[n];
//    msort(a, tem, 0, n - 1);
//}
//
//int main() {
//    int i = 0;
//
//    while (~scanf("%d", &n)) {
//        for (i = 0; i < n; i++) cin >> a[i];
//        merge_sort(a, n);
//    }
//
//    return 0;
//}