/**
* 7-4 两个有序链表序列的交集
* 已知两个非降序链表序列S1与S2，设计函数构造出S1与S2的交集新链表S3
*/
#include "iostream"
#include "list"

using namespace std;

int main() {
    list<int> list1;
    list<int> list2;
    list<int> list3;

    int i;
    while (cin >> i) {
        if (i == -1) break;
        else list1.push_back(i);
    }
    while (cin >> i) {
        if (i == -1) break;
        else list2.push_back(i);
    }

    auto it1 = list1.begin();
    auto it2 = list2.begin();

    int j1 = 0, j2 = 0;
    while (j1 < list1.size() && j2 < list2.size()) {
        if (*it1 > *it2) {
            it2++;
            j2++;
        } else if (*it1 < *it2) {
            it1++;
            j1++;
        } else {
            list3.push_back(*it1);
            it1++;
            it2++;
            j1++;
            j2++;
        }
    }

    if (list3.empty()) cout << "NULL";
    else {
        auto it3 = list3.begin();
        for (int i = 0; i < list3.size() - 1; i++) {
            cout << *it3 << " ";
            it3++;
        }
        cout << *it3;
    }

    return 0;
}