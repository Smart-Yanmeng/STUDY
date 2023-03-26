//
// Created by yorky on 2023/2/26.
//

/**
 * 十进制转二进制
 */

#include "iostream"

using namespace std;

void decimalToBinary(int num) {
    int binary[32];
    int i = 0;

    while (num > 0) {
        binary[++i] = num % 2;
        num /= 2;
    }

    cout << "所输入数的二进制形式是：";

    for (int j = i; j > 0; j--) {
        cout << binary[j];
    }

    cout << endl;
}

int main() {
    int num = 0;

    while (1) {
        cout << "请输入一个数字：" << endl;
        cin >> num;

        decimalToBinary(num);
    }
}