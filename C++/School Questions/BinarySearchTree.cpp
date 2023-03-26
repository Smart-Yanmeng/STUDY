#include "iostream"

using namespace std;

typedef struct Node {
    int data;

    struct Node *left, *right;
} Node, *Tree;

bool flag = true;

Tree BuildTree(int n, int *preorderTraverse) {
    if (n == 0) return nullptr;

    Tree tree = (Tree) malloc(sizeof(Node));
    tree->data = preorderTraverse[0];
    tree->left = nullptr;
    tree->right = nullptr;

    int leftNum = 0, rightNum = 0;

    for (int i = 1; i < n; i++) {
        if (preorderTraverse[i] >= preorderTraverse[0]) {
            leftNum = i - 1;
            rightNum = n - i;

            break;
        }
        if (preorderTraverse[i] < preorderTraverse[0] && i == n - 1) {
            leftNum = i;
            rightNum = 0;
        }
    }

    tree->left = BuildTree(leftNum, preorderTraverse + 1);
    tree->right = BuildTree(rightNum, preorderTraverse + leftNum + 1);

    return tree;
}

Tree BuildTreeAsMirror(int n, int *preorderTraverse) {
    if (n == 0) return nullptr;

    Tree tree = (Tree) malloc(sizeof(Node));
    tree->data = preorderTraverse[0];
    tree->left = nullptr;
    tree->right = nullptr;

    int leftNum = 0, rightNum = 0;

    for (int i = 1; i < n; i++) {
        if (preorderTraverse[i] < preorderTraverse[0]) {
            leftNum = i - 1;
            rightNum = n - i;

            break;
        }
        if (preorderTraverse[i] >= preorderTraverse[0] && i == n - 1) {
            leftNum = i;
            rightNum = 0;
        }
    }

    tree->left = BuildTreeAsMirror(leftNum, preorderTraverse + 1);
    tree->right = BuildTreeAsMirror(rightNum, preorderTraverse + leftNum + 1);

    return tree;
}

void recursionPostorderTraversal(Tree tree) {
    if (tree) {
        recursionPostorderTraversal(tree->left);
        recursionPostorderTraversal(tree->right);
        if (flag) {
            cout << tree->data;
            flag = false;
        } else cout << " " << tree->data;
    }
}

bool isBinarySearchTree(int n, int *preorderTraverse) {
    if (n == 0) return true;

    bool flag = true;
    int left = 0, right = 0;
    for (int i = 1; i < n; i++) {
        if (preorderTraverse[i] >= preorderTraverse[0]) {
            flag = false;
            left = i - 1;
            right = n - i;

            break;
        }
    }
    for (int i = left + 1; i < n; i++) if (!flag && preorderTraverse[i] < preorderTraverse[0]) return false;

    bool leftJudge = isBinarySearchTree(left, preorderTraverse + 1);
    bool rightJudge = isBinarySearchTree(right, preorderTraverse + left + 1);

    return leftJudge && rightJudge;
}

bool isBinarySearchTreeIsMirror(int n, int *preorderTraverse) {
    bool flag2 = true;
    for (int i = 1; i < n; i++) {
        if (preorderTraverse[i] <= preorderTraverse[0] && flag2) flag2 = false;
        else if (preorderTraverse[i] > preorderTraverse[0] && !flag2) return false;
    }

    return true;
}

int main() {
    int n;
    cin >> n;

    int preorderTraverse[n];

    for (int i = 0; i < n; i++) cin >> preorderTraverse[i];

    if (isBinarySearchTree(n, preorderTraverse)) {
        cout << "YES" << endl;
        Tree tree = BuildTree(n, preorderTraverse);

        recursionPostorderTraversal(tree);
    } else if (isBinarySearchTreeIsMirror(n, preorderTraverse)) {
        cout << "YES" << endl;
        Tree tree = BuildTreeAsMirror(n, preorderTraverse);

        recursionPostorderTraversal(tree);
    } else cout << "NO";

    return 0;
}