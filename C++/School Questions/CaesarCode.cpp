/**
 * 双向链表的凯撒密码
 */

#include <iostream>

using namespace std;

#define NUM 26
#define LEN sizeof(struct LinkList)

// 构建一个双循环链表结点的结构体
struct LinkList {
    char word;
    struct LinkList *prior;
    struct LinkList *next;
};

// 开始初始化这个结构体，让26个字母插入，并形成双循环链表
struct LinkList *init(void) {
    struct LinkList *L, *temp, *head;
    int i = 0;  //用于计数
    L = (struct LinkList *) malloc(LEN);
    head = temp = L;
    //上面用于处理好头部
    while (i < NUM) {
        L->word = 'A' + i;
        temp->next = L;   // 后继是 L
        L->prior = temp;  // 前驱是 L 的替身 temp
        temp = L;         // 让 temp 再次成为 L 的替身
        L = (struct LinkList *) malloc(LEN); // L 重新分配内存空间
        i++;
    }
    temp->next = head;  //最后一个结点指向头部
    head->prior = temp; //头部的前驱指向尾部，形成闭环
    return head;   //返回头部地址
}

void print_word(struct LinkList *L) {
    struct LinkList *head;
    head = L;
    do {
        cout << L->word << " ";
    } while ((L = L->next) != head);
}

struct LinkList *caesar(struct LinkList *L, int n) {
    if (n > 0) {
        do {
            L = L->next;
        } while (--n);
    }

    if (n < 0) {
        do {
            L = L->prior;
        } while (++n);
    }

    return L;
}


int main() {
    struct LinkList *word_p;
    int n;
    cin >> n;

    word_p = init();

    word_p = caesar(word_p, n);
    print_word(word_p);
}