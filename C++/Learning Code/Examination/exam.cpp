#include "iostream"

using namespace std;

int n;
int arr[101], brr[101];

void print() {
    for (int i = 0; i < n; i++) cout << arr[i];
    cout << endl;
}

void dfs(int cur) {
    if (cur == n) {
        print();

        return;
    }

    for (int i = 0; i < n; i++) {
        if (brr[i] == 0) {
            arr[cur] = i + 1;
            brr[i] = 1;
            dfs(cur + 1);
            brr[i] = 0;
        }
    }
}

int main() {
    cin >> n;
    dfs(0);

    return 0;
}