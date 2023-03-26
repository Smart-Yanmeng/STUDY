#include "iostream"
#include "vector"
#include "queue"

using namespace std;

struct node {
    int data;
    struct node *left;
    struct node *right;
};

vector<string> binaryTreePaths(node *root) {
    if (root == nullptr) return {};
    vector<string> res;

    queue<node *> node_que;
    queue<string> path_que;
    node_que.push(root);
    path_que.push(to_string(root->data));

    while (!node_que.empty()) {
        node *node = node_que.front();
        node_que.pop();

        string path = path_que.front();
        path_que.pop();

        if (node->left == nullptr && node->right == nullptr) {
            res.push_back(path);
        } else {
            if (node->left != nullptr) {
                node_que.push(node->left);
                path_que.push(path + " " + to_string(node->left->data));
            }
            if (node->right != nullptr) {
                node_que.push(node->right);
                path_que.push(path + " " + to_string(node->right->data));
            }
        }
    }

    return res;
}


void printPaths(struct node *node) {
    vector<string> res = binaryTreePaths(node);

    for (const auto &a: res) {
        cout << a << endl;
    }
//
//    struct node *root, *q;
//    int i;
//    int top = 0;
//    struct node *arr[1000];
//
//    q = nullptr;
//    root = node;
//
//    while (root != nullptr || top != 0) {
//        if (root != nullptr) {
//            arr[++top] = root;
//            root = root->left;
//        } else if (top > 0) {
//            root = arr[top];
//            if (root->right == nullptr || root->right == q) {
//                if (root->left == nullptr && root->right == nullptr) {
//                    for (i = 1; i <= top; i++) {
//                        if (i != top) cout << arr[i]->data << " ";
//                        else cout << arr[i]->data << endl;
//                    }
//                }
//                q = root;
//                top--;
//
//                root = nullptr;
//            } else root = root->right;
//        }
//    }

}
//请补足这个函数，函数目标主要是递归的输出所有从根结点到叶子结点的路径

struct node *creatTree() {
    struct node *T = NULL;
    int val;
    scanf("%d", &val);
    if (val == -1) {
        T = NULL;//说明该元素无效
        return T;
    } else {
        T = (struct node *) malloc(sizeof(struct node));
        if (T == NULL)//递归结束标识
            return NULL;
        T->data = val;
        T->left = NULL;//先置空
        T->right = NULL;
        T->left = creatTree();
        T->right = creatTree();
        return T;
    }
}

//主函数
int main() {
    struct node *T = creatTree();//用先序遍历创建树
    printPaths(T); //需要补充此函数

    return 0;
}