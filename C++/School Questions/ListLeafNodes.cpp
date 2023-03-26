#include "iostream"
#include "queue"

using namespace std;

int main() {
    int n;
    cin >> n;

    bool flag = true;

    queue<int> queue;

    bool flags[n];
    for (int i = 0; i < n; i++) flags[i] = true;
    pair<int, int> pairs[n];

    for (int i = 0; i < n; i++) {
        char a;
        cin >> a;
        if (a == '-') pairs[i].first = -1;
        else {
            pairs[i].first = a - '0';
            flags[pairs[i].first] = false;
        }
        cin >> a;
        if (a == '-') pairs[i].second = -1;
        else {
            pairs[i].second = a - '0';
            flags[pairs[i].second] = false;
        }
    }

    int root;
    for (int i = 0; i < n; i++) if (flags[i]) root = i;

    queue.push(root);
    while (!queue.empty()) {
        root = queue.front();
        if (pairs[root].first == -1 && pairs[root].second == -1) {
            if (flag) {
                cout << root;
                flag = false;
            } else cout << " " << root;
        } else {
            if (pairs[root].first != -1) queue.push(pairs[root].first);
            if (pairs[root].second != -1) queue.push(pairs[root].second);
        }

        queue.pop();
    }

    return 0;
}