/**
* 7-1 堆栈操作合法性
*/
#include "iostream"

using namespace std;

bool judge(const string &str, const int &sqr) {
    int flag = 0;

    for (int i = 0; i < str.length(); i++) {
        if (str[i] == 'S') {
            flag++;
            if (flag > sqr) return false;
        } else {
            flag--;
            if (flag < 0) return false;
        }
    }
    if (flag != 0) return false;
    else return true;
}

int main() {
    int num, sqr;
    cin >> num >> sqr;

    for (int i = 0; i < num; i++) {
        string str;
        cin >> str;

        judge(str, sqr) ? cout << "YES" << endl : cout << "NO" << endl;
    }

    return 0;
}