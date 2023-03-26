/**
 * 两个有序链表序列的合并
 */

#include "iostream"

using namespace std;

typedef struct node {
    int data;
    struct node *next;
} *LinkList;

LinkList createLinkList() {
    LinkList head, L, pro;
    int member;

    L = (LinkList) malloc(sizeof(struct node));
    head = L;

    cin >> member;
    if (member == -1) {
        L = nullptr;

        return L;
    }

    while (true) {
        if (member == -1) {
            pro->next = nullptr;
            free(L);

            return head;
        }

        L->data = member;
        L->next = (LinkList) malloc(sizeof(struct node));
        pro = L;
        L = L->next;

        cin >> member;
    }
}

LinkList mergeLinkedList(LinkList l1, LinkList l2) {
    LinkList L, head;
    L = (LinkList) malloc(sizeof(LinkList));
    head = L;

    while (l1 && l2) {
        if (l1->data <= l2->data) {
            L->next = l1;
            L = L->next;
            l1 = l1->next;
        } else {
            L->next = l2;
            L = L->next;
            l2 = l2->next;
        }
    }

    if (l1) {
        L->next = l1;
        L = head;
        head = head->next;

        return head;
    } else if (l2) {
        L->next = l2;
        L = head;
        head = head->next;

        return head;
    } else {
        return nullptr;
    }
}

int main() {
    LinkList L, l1, l2;
    l1 = createLinkList();
    l2 = createLinkList();
    L = mergeLinkedList(l1, l2);
    if (!L) cout << "NULL";
    else {
        while (L->next) {
            cout << L->data << " ";
            L = L->next;
        }

        cout << L->data;
    }

    return 0;
}