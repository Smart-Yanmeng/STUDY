#include "iostream"

using namespace std;

#define MAXSIZE 50     // MAXSIZE为最大数据元素数目
typedef int ElemType;
typedef struct {
    ElemType r[MAXSIZE + 1];    // r[0]闲置或另作它用
    int length;
} SqList;

void exchange(int *p, int *q) {
    int temp = *p;
    *p = *q;
    *q = temp;
}

int quickSort(int lis[], int left, int right) {
    if (left >= right) return 0;

    int datum = lis[left];
    int m = left;
    int n = right;

    while (m != n) {
        while (m < n && lis[n] >= datum) n--;
        exchange(&lis[m], &lis[n]);

        while (m < n && lis[m] <= datum) m++;
        exchange(&lis[m], &lis[n]);
    }

    quickSort(lis, m + 1, right);
    quickSort(lis, left, m - 1);
}

int main() {
    int n;
    cin >> n;

    SqList lis;
    lis.length = 0;
    lis.r[0] = 0;

    for (int i = 0; i < n; i++) {
        cin >> lis.r[i + 1];
        lis.length++;
    }

    quickSort(lis.r, 1, n);

    for (int i = 1; i <= n; i++) {
        if (i == n) cout << lis.r[i];
        else cout << lis.r[i] << " ";
    }

    return 0;
}