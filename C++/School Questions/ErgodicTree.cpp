#include "iostream"
#include "queue"

using namespace std;

typedef struct Node {
    int data;

    struct Node *left, *right;
} Node, *Tree;

bool flag = true;

Tree buildTree(int n, int *postOrder, int *intermediateOrder) {
    if (n == 0) return nullptr;

    Tree tree = (Tree) malloc(sizeof(Node));
    tree->data = postOrder[n - 1];
    tree->left = nullptr;
    tree->right = nullptr;

    int leftPoint = 0, rightPoint = 0;
    for (int i = 0; i < n; i++) {
        if (intermediateOrder[i] == postOrder[n - 1]) {
            leftPoint = i;
            rightPoint = n - i - 1;

            break;
        }
    }

    tree->left = buildTree(leftPoint, postOrder, intermediateOrder);
    tree->right = buildTree(rightPoint, postOrder + leftPoint, intermediateOrder + leftPoint + 1);

    return tree;
}

void sequenceOrder(Tree tree) {
    if (!tree) return;

    queue<Tree> queue;

    queue.push(tree);
    while (!queue.empty()) {
        Node node = *queue.front();

        if (flag) {
            cout << node.data;
            flag = false;
        } else cout << " " << node.data;

        if (node.left) queue.push(node.left);
        if (node.right) queue.push(node.right);

        queue.pop();
    }
}

int main() {
    int n;
    cin >> n;

    int postOrder[n], intermediateOrder[n];

    for (int i = 0; i < n; i++) cin >> postOrder[i];
    for (int i = 0; i < n; i++) cin >> intermediateOrder[i];

    Tree tree = buildTree(n, postOrder, intermediateOrder);
    sequenceOrder(tree);

    return 0;
}