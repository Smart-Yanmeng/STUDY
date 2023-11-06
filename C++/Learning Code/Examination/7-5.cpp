#include "iostream"

using namespace std;

int main() {
    int T;
    cin >> T;

    string gan[10] = {"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"};
    string zhi[12] = {"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"};

    for (int i = 0; i < T; i++) {
        int year;
        cin >> year;

        int delta = year - 1924;
        int ganIndex = delta % 10;
        int zhiIndex = delta % 12;

        if (ganIndex < 0) ganIndex += 10;
        if (zhiIndex < 0) zhiIndex += 12;

//        cout << ganIndex << zhiIndex << endl;

        cout << gan[ganIndex] << zhi[zhiIndex] << endl;
    }

    return 0;
}