#include "iostream"
#include "list"

using namespace std;

int main() {
    list<int> list;
    int i = 0;
    while (cin >> i) {
        if (i == -1) break;
        list.push_back(i);
    }

    cout << *list.begin() << " " << *list.end() << " " << list.size();

    return 0;
}