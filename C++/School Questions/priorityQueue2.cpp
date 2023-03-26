#include "iostream"
#include "queue"

using namespace std;

priority_queue<int, vector<int>, greater<int> > q;

int main() {
    int testNum;
    cin >> testNum;

    for (int i = 0; i < testNum; i++) {
        int tempNum;
        cin >> tempNum;

        q.push(tempNum);
    }

    int ans = 0;
    while (q.size() > 1) {
        int x = q.top();
        q.pop();

        int y = q.top();
        q.pop();

        q.push(x + y);
        ans += x + y;
    }

    cout << ans;

    return 0;
}