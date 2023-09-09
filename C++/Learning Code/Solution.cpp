#include "iostream"
#include "cstring"

using namespace std;

int main() {
    int n;
    cin >> n;
    int a[n][n];
    memset(a, 0, sizeof(a));
    int x = 0, y = 0;

    int move1[4] = {0, 1, 0, -1};
    int move2[4] = {1, 0, -1, 0};

    int func = 0;

    for (int i = 0; i < n * n; i++) {
        a[x][y] = i + 1;
        if (x + move1[func] == n || y + move2[func] == n) func = (func + 1) % 4;
        if (a[x + move1[func]][y + move2[func]] != 0) func = (func + 1) % 4;

        x += move1[func];
        y += move2[func];
    }

    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            cout << a[i][j] << " ";
        }
        cout << endl;
    }
}