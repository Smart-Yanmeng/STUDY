#include "iostream"

using namespace std;

void swap(int *x, int *y) {
    int temp;
    temp = *x;
    *x = *y;
    *y = temp;
}

int main() {
    int N;
    cin >> N;
    while (N--) {
        int L, num[51], i, j, count = 0;
        cin >> L;
        for (i = 0; i < L; i++) cin >> num[i];
        for (i = 1; i < L; i++) {
            for (j = 0; j < L - i; j++) {
                if (num[j] > num[j + 1]) {
                    swap(num + j, num + j + 1);
                    count++;
                }
            }
        }

        cout << "Optimal train swapping takes " << count << " swaps.";
    }

    return 0;
}

