#include "iostream"

using namespace std;

typedef struct Node {
    int data;

    struct Node *left, *right;
} Node, *Tree;

bool flag = true;

void BuildTree(struct Node *&tree) {
    int data;
    cin >> data;
    if (data != -1) {
        tree = (Tree) malloc(sizeof(Tree));
        tree->data = data;
        BuildTree(tree->left);
        BuildTree(tree->right);
    } else {
        tree = nullptr;
    }
}

void OutputLeft(struct Node *root) {
    if (root == nullptr) return;

    if (root->left != nullptr) {
        cout << root->data << " ";
        OutputLeft(root->left);
    } else if (root->right != nullptr) {
        cout << root->data << " ";
        OutputLeft(root->right);
    }
}

void OutputLeaves(struct Node *root) {
    if (root == nullptr) return;

    OutputLeaves(root->left);

    if (!(root->left) && !(root->right)) {
        cout << root->data << " ";
    }

    OutputLeaves(root->right);
}

void OutputRight(struct Node *root) {
    if (root == nullptr) return;

    if (root->right != nullptr) {
        OutputRight(root->right);
        cout << root->data << " ";
    } else if (root->left) {
        OutputRight(root->left);
        cout << root->data << " ";
    }
}

int main() {
    struct Node *root = nullptr;
    BuildTree(root);

    OutputLeft(root);
    OutputLeaves(root);
    OutputRight(root->right);
}