#include "iostream"

using namespace std;

int n;
int Cnt = 0;
bool A[10], B[20], C[20];

void searchN(int cur) {
    if (cur == n) {
        Cnt++;
        return;
    }

    for (int i = 0; i < n; i++) {
        if (!A[i] && !B[cur + i] && !C[cur - i + n]) {
            A[i] = B[cur + i] = C[cur - i + n] = true;
            searchN(cur + 1);
            A[i] = B[cur + i] = C[cur - i + n] = false;
        }
    }
}

int main() {
    while (true) {
        cin >> n;
        if (n == 0) break;

        searchN(0);
        cout << Cnt << endl;

        Cnt = 0;
    }

    return 0;
}