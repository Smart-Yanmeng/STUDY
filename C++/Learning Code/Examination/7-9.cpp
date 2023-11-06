#include "iostream"

using namespace std;

const int maxn = 10;
int n, Cnt = 0;
bool A[maxn], B[maxn], C[maxn];

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
        if (n == 0) return 0;

        searchN(0);
        cout << Cnt << endl;
        Cnt = 0;
    }
}