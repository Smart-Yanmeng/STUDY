#include "iostream"

using namespace std;

int ans[10], ret[10], n;

void dfs(int flag) {
    if (flag == n + 1) {
        for (int i = 1; i <= n; i++) cin >> ans[i];
        printf("\n");
    } else {
        for (int i = 1; i <= n; i++) {
            if (!ret[i]) {
                ans[flag] = i;
                ret[i] = 1;
                dfs(flag + 1);
                ret[i] = 0;
            }
        }
    }
}

int main() {
    cin >> n;
    dfs(1);

    return 0;
}