/**
* 7-2 简单计算器
* 本题要求你为初学数据结构的小伙伴设计一款简单的利用堆栈执行的计算器。
* 如上图所示，计算器由两个堆栈组成，一个堆栈 S1 存放数字，另一个堆栈
* S2 存放运算符。计算器的最下方有一个等号键，每次按下这个键，计算器
* 就执行以下操作：
* 1. 从 S1 中弹出两个数字，顺序为 n1 和 n2 ;
* 2. 从 S2 中弹出一个运算符 op;
* 3. 执行计算 n2 op n1;
* 4. 将得到的结果压回 S1.
* 直到两个堆栈都为空时，计算结束，最后的结果将显示在屏幕上。
*/
#include "iostream"
#include "stack"

using namespace std;

int main() {
    stack<int> numStack;
    stack<char> symStack;
    int num;
    cin >> num;
    for (int i = 0; i < num; i++) {
        int n;
        cin >> n;
        numStack.push(n);
    }
    for (int i = 0; i < num - 1; i++) {
        char sym;
        cin >> sym;
        symStack.push(sym);
    }

    while (!symStack.empty()) {
        int num1 = numStack.top();
        numStack.pop();
        int num2 = numStack.top();
        numStack.pop();

        char sym = symStack.top();
        symStack.pop();

        switch (sym) {
            case '+':
                numStack.push(num2 + num1);
                break;
            case '-':
                numStack.push(num2 - num1);
                break;
            case '*':
                numStack.push(num2 * num1);
                break;
            case '/':
                if (num1 == 0) {
                    cout << "ERROR: " << num2 << "/0";
                    return 0;
                } else {
                    numStack.push(num2 / num1);
                    break;
                }
        }
    }
    cout << numStack.top();

    return 0;
}