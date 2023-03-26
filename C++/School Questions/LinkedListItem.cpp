#include "iostream"

using namespace std;

typedef struct node {
    int data;
    struct node *next;
} *LinkList;

int main() {
    LinkList l1 = nullptr, l2 = nullptr;

    int k, num = 0;
    cin >> k;

    int member;
    cin >> member;
    while (member >= 0) {
        // 头插法
        l1 = (LinkList) malloc(sizeof(struct node));
        l1->data = member;
        l1->next = l2;

        l2 = l1;
        num++;

        cin >> member;
    }

    if (k > num) {
        cout << "NULL";

        return 0;
    }

    for (int i = 0; i < k - 1; i++) {
        l2 = l2->next;
    }

    cout << l2->data;

}