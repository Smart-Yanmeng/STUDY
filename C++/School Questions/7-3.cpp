/**
 * 7-3 凯撒密码
 */
#include "iostream"
#include "list"

using namespace std;

// 字符转置
char strExchange(char str, int lab) {
    if (str >= 'a' && str <= 'z') {
        if (str + lab % 26 > 'z') return str + lab % 26 - 26;
        else if (str + lab % 26 < 'a') return str + lab % 26 + 26;
        else return str + lab % 26;
    } else if (str >= 'A' && str <= 'Z') {
        if (str + lab % 26 > 'Z') return str + lab % 26 - 26;
        else if (str + lab % 26 < 'A') return str + lab % 26 + 26;
        else return str + lab % 26;
    } else return str;
}

int main() {
    list<char> list1;

    while (true) {
        char str = getchar();
        if (str == '\n') break;
        else list1.push_back(str);
    }

    int lab = 0;
    cin >> lab;

    for (char &it: list1) {
        cout << strExchange(it, lab);
    }

    return 0;
}