#include "iostream"

using namespace std;

int main() {
    int x = 100;
    cout << (x + (1 << 31));

    return 0;
}