#include "iostream"

using namespace std;

void printBinary(int num) {
    int numBits = sizeof(num) * 8;  // 计算整数的位数
    int bit;

    // 从最高位开始逐位输出
    for (int i = numBits - 1; i >= 0; i--) {
        bit = (num >> i) & 1;  // 取出第i位的值
        printf("%d", bit);
    }

    cout << endl;
}

int bitCount(int x) {

    int count;

    int tmpMask1 = (0x55) | (0x55 << 8);
    int mask1 = (tmpMask1) | (tmpMask1 << 16);

    int tmpMask2 = (0x33) | (0x33 << 8);
    int mask2 = (tmpMask2) | (tmpMask2 << 16);

    int tmpMask3 = (0x0f) | (0x0f << 8);
    int mask3 = (tmpMask3) | (tmpMask3 << 16);

    int mask4 = (0xff) | (0xff << 16);

    int mask5 = (0xff) | (0xff << 8);

    count = (x & mask1) + ((x >> 1) & mask1);
    count = (count & mask2) + ((count >> 2) & mask2);
    count = (count + (count >> 4)) & mask3;
    count = (count + (count >> 8)) & mask4;
    count = (count + (count >> 16)) & mask5;

    return count;
}

int main() {
    cout << bitCount(0x1234);

    return 0;
}