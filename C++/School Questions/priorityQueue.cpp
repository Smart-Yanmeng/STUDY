#include "iostream"
#include "vector"
#include "queue"
#include "string"

using namespace std;

// 小顶堆
priority_queue<int, vector<int>, greater<int> > q;

int main() {
    int n = 0;
    cin >> n;
    int frequency[n];
    for (int i = 0; i < n; i++) {
        getchar();
        getchar();
        cin >> frequency[i];
        q.push(frequency[i]);
    }

    // 求出霍夫曼编码长度
    int ans = 0;
    while (q.size() > 1) {
        int x = q.top();
        q.pop();
        int y = q.top();
        q.pop();
        q.push(x + y);

        ans += x + y;
    }

    int numTest;
    cin >> numTest;
    while (numTest--) {
        string code[n];
        for (int i = 0; i < n; i++) {
            getchar();
            getchar();
            cin >> code[i];
        }

        int tempAns = 0;
        for (int i = 0; i < n; i++) {
            tempAns += frequency[i] * code[i].length();
        }

        // 判断编码长度是否相同
        if (ans != tempAns) {
            cout << "No" << endl;
        } else {
            // 某个编码不得为其他编码的前缀中出现过
            bool flag = false;
            for (int i = 0; i < n; i++) {
                for (int j = 0; j < n; j++) {
                    if (i != j) {
                        // 保证前缀
                        if (code[i].length() >= code[j].length()) {
                            // 如果相同，则编码无效
                            if (code[j] == code[i].substr(0, code[j].length())) {
                                flag = true;
                            }
                        }
                    }

                    if (flag) break;
                }

                if (flag) break;
            }

            if (flag) cout << "No" << endl;
            else cout << "Yes" << endl;
        }
    }

    return 0;
}

