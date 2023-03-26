#include "iostream"
#include "map"

using namespace std;

int main() {
    int pointNum, edgeNum, pointF, pointS, weight;
    cin >> pointNum >> edgeNum;

    // 创建图
    map<int, map<int, int>> myMap;
    for (int i = 0; i < edgeNum; i++) {
        cin >> pointF >> pointS >> weight;
        myMap[pointF][pointS] = weight;
    }

    for (auto i: myMap) {
        int a = i.first;
        string s = "";
        for (auto j: i.second) {
            s += "(" + to_string(a) + "," + to_string(j.first) + "," + to_string(j.second) + ")";
        }
        cout << a << ":" << s << endl;
    }

    return 0;
}