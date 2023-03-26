#include <bits/stdc++.h>

using namespace std;

#define MAX_LENGTH 100

vector<int> G[MAX_LENGTH];
int n, m, inDegree[MAX_LENGTH];
stack<int> s;

bool topSort() {
    int num = 0;
    for (int i = 1; i <= n; i++) {
        if (inDegree[i] == 0) {
            s.push(i);
        }
    }

    while (!s.empty()) {
        int u = s.top();
        cout << u << " ";
        s.pop();

        for (int i = 0; i < G[u].size(); i++) {
            int v = G[u][i];
            inDegree[v]--;

            if (inDegree[v] == 0) {
                s.push(v);
            }
        }

        num++;
    }

    // 判断是否有环
    if (num == n) return true;
    else return false;
}

int main() {
    cin >> n >> m;

    // 初始化
    memset(inDegree, 0, sizeof(inDegree));

    for (int i = 1; i <= m; i++) {
        int x, y;
        cin >> x >> y;
        inDegree[y]++;

        // 头插法
        if (G[x].empty()) G[x].push_back(y);
        else G[x].insert(G[x].begin(), y);
    }

    if (!topSort()) cout << endl << 0;

    return 0;
}