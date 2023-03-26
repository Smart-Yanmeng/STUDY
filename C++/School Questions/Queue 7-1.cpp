#include "iostream"
#include "queue"

using namespace std;

int main() {
    int n, jA = 0, jB = 0;
    cin >> n;

    queue<int> qA, qB;

    for (int i = 0; i < n; i++) {
        int v;
        cin >> v;

        if (v % 2 == 1) qA.push(v);
        else qB.push(v);
    }

    jA = qA.size();
    jB = qB.size();

    while (jA && jB) {
        cout << qA.front() << " ";
        qA.pop();
        jA--;
        if (jA) {
            cout << qA.front() << " ";
            qA.pop();
            jA--;
        }

        if (jA == 0 && jB == 1) {
            cout << qB.front();
            return 0;
        }

        cout << qB.front() << " ";
        qB.pop();
        jB--;
    }

    while (!jA && jB) {
        if (jB == 1) {
            cout << qB.front();
            return 0;
        }

        cout << qB.front() << " ";
        qB.pop();
        jB--;
    }

    while (jA && !jB) {
        if (jA == 1) {
            cout << qA.front();
            return 0;
        }

        cout << qA.front() << " ";
        qA.pop();
        jA--;
    }

    return 0;
}