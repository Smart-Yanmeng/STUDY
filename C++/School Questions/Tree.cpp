#include "iostream"

using namespace std;

typedef struct Node {
    int data;

    struct Node *left, *right;
} Node, *Tree;

int flag = false;

Tree BuildTree(int n, int *intermediateOrder, int *postOrder) {
    if (n == 0) return nullptr;

    // 创建树根
    Tree tree = (Tree) malloc(sizeof(Node));
    tree->data = postOrder[n - 1];
    tree->left = nullptr;
    tree->right = nullptr;

    // 左右子树的结点数
    int leftNum = 0, rightNum = 0;
    for (int i = 0; i < n; i++) {
        if (intermediateOrder[i] == postOrder[n - 1]) {
            leftNum = i;
            rightNum = n - i - 1;

            break;
        }
    }

    // 递归遍历
    tree->left = BuildTree(leftNum, intermediateOrder, postOrder);
    tree->right = BuildTree(rightNum, intermediateOrder + leftNum + 1, postOrder + leftNum);

    return tree;
}

void preOrder(Tree tree) {
    if (!tree) return;

    if (!flag) {
        flag = true;

        cout << tree->data;
        preOrder(tree->left);
        preOrder(tree->right);
    } else {
        cout << " " << tree->data;
        preOrder(tree->left);
        preOrder(tree->right);
    }
}

int main() {
    int n;
    cin >> n;

    int intermediateOrder[n], postOrder[n];
    for (int i = 0; i < n; i++) cin >> postOrder[i];
    for (int i = 0; i < n; i++) cin >> intermediateOrder[i];

    Tree tree = BuildTree(n, intermediateOrder, postOrder);

    cout << "Preorder: ";
    preOrder(tree);

    return 0;
}