#include "iostream"

using namespace std;

#define MAXSIZE 50     // MAXSIZE为最大数据元素数目
typedef int ElemType;
typedef struct {
    ElemType r[MAXSIZE + 1];    // r[0]闲置或另作它用
    int length;
} SqList;

int main() {
    int n;
    cin >> n;

    SqList lis;
    lis.length = 0;
    lis.r[0] = -1;

    for (int i = 0; i < n; i++) {
        cin >> lis.r[i + 1];
        lis.length++;
    }

    for (int i = 2; i <= n; i++) {
        lis.r[0] = lis.r[i];
        int j = i - 1;

        while (j >= 0 && lis.r[0] < lis.r[j]) {
            lis.r[j + 1] = lis.r[j];
            j--;
        }

        lis.r[j + 1] = lis.r[0];
    }

    for (int i = 1; i <= n; i++) {
        if (i == n) {
            cout << lis.r[i];
        } else cout << lis.r[i] << " ";
    }

    return 0;
}