#include "iostream"

using namespace std;

typedef struct Node {
    int data;

    struct Node *left, *right;
} Node, *Tree;

// 创建新节点
Tree NewNode(int nodeData) {
    Tree tree = (Tree) malloc(sizeof(Node));
    tree->data = nodeData;
    tree->left = nullptr;
    tree->right = nullptr;

    return tree;
}

// 插入新节点
Tree InsertNode(Tree tree, int nodeData) {
    if (!tree) tree = NewNode(nodeData);
    else {
        if (nodeData < tree->data) tree->left = InsertNode(tree->left, nodeData);
        else tree->right = InsertNode(tree->right, nodeData);
    }

    return tree;
}

// 建树
Tree BuildTree(int n, int *sequence) {
    Tree tree;

    int root = sequence[0];
    tree = NewNode(root);
    for (int i = 1; i < n; i++) InsertNode(tree, sequence[i]);

    return tree;
}

// 判断两棵树是否相同
bool IsSameTree(Tree tree1, Tree tree2) {
    if (tree1 == nullptr && tree2 == nullptr) return true;
    else if (tree1 == nullptr && tree2 != nullptr || tree1 != nullptr && tree2 == nullptr) return false;
    else if (tree1->data != tree2->data) return false;
    else return IsSameTree(tree1->left, tree2->left) && IsSameTree(tree1->right, tree2->right);
}

int main() {
    while (true) {
        int n, l;

        cin >> n;
        if (n == 0) return 0;

        cin >> l;

        int initialSequence[n];
        for (int i = 0; i < n; i++) cin >> initialSequence[i];
        Tree firstTree = BuildTree(n, initialSequence);

        int sequence[n];
        for (int i = 0; i < l; i++) {
            for (int j = 0; j < n; j++) cin >> sequence[j];
            Tree secondTree = BuildTree(n, sequence);

            if (IsSameTree(firstTree, secondTree)) cout << "Yes" << endl;
            else cout << "No" << endl;
        }
    }

}