#include "iostream"
#include "set"

#define MAXPOINT 100

int adjacencyMatrix[MAXPOINT][MAXPOINT] = {0};
int book[MAXPOINT] = {0};

using namespace std;

int pointNum, edgeNum;

void createAdjacencyMatrix() {
    // 录入顶点数和边数
    cin >> pointNum >> edgeNum;

    // 构建邻接矩阵
    for (int i = 0; i < edgeNum; i++) {
        int pointF, pointS;
        cin >> pointF >> pointS;
        adjacencyMatrix[pointF - 1][pointS - 1] = 1;
        adjacencyMatrix[pointS - 1][pointF - 1] = 1;
    }
}

set<int> s;

void dfs(int u) {
    s.insert(u + 1);
    book[u] = 1;

    for (int i = 0; i < pointNum; i++) {
        if (adjacencyMatrix[u][i] != 0 && !book[i]) {
            dfs(i);
        }
    }
}

void dfsTraverse() {
    for (int i = 0; i < pointNum; i++) {
        if (!book[i]) {
            s.clear();
            cout << "(";

            dfs(i);

            bool flag = true;
            for (auto a: s) {
                if (flag) {
                    cout << a;
                    flag = false;
                } else cout << "," << a;
            }

            cout << ")" << endl;
        }
    }
}

int main() {
    createAdjacencyMatrix();
    dfsTraverse();

    return 0;
}