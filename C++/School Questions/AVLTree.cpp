#include "iostream"

using namespace std;

typedef struct node *AVLNode;
struct node {
    AVLNode left, right;
    int data;
};

int GetHeight(AVLNode a) {
    if (!a) return 0;
    else {
        int lHeight = GetHeight(a->left);
        int rHeight = GetHeight(a->right);

        return lHeight > rHeight ? (lHeight + 1) : (rHeight + 1);
    }
}

AVLNode LLRotation(AVLNode a) {
    AVLNode b = a->left;
    a->left = b->right;
    b->right = a;

    return b;
}

AVLNode RRRotation(AVLNode a) {
    AVLNode b = a->right;
    a->right = b->left;
    b->left = a;

    return b;
}

AVLNode LRRotation(AVLNode a) {
    a->left = RRRotation(a->left);
    return LLRotation(a);
}

AVLNode RLRotation(AVLNode a) {
    a->right = LLRotation(a->right);
    return RRRotation(a);
}

AVLNode Insert(AVLNode a, int data) {
    if (!a) {
        a = (AVLNode) malloc(sizeof(struct node));
        a->data = data;
        a->left = nullptr;
        a->right = nullptr;

        return a;
    } else if (data < a->data) {
        a->left = Insert(a->left, data);
        if (GetHeight(a->left) - GetHeight(a->right) == 2) {
            if (data < a->left->data) a = LLRotation(a);
            else a = LRRotation(a);
        }
    } else if (data > a->data) {
        a->right = Insert(a->right, data);
        if (GetHeight(a->left) - GetHeight(a->right) == -2) {
            if (data > a->right->data) a = RRRotation(a);
            else a = RLRotation(a);
        }
    }

    return a;
}

int main() {
    AVLNode a = nullptr;
    int n = 0, temp = 0;
    cin >> n;
    for (int i = 0; i < n; i++) {
        cin >> temp;

        a = Insert(a, temp);
    }
    cout << a->data;

    return 0;
}