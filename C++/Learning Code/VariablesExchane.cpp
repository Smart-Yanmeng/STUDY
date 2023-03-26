//
// Created by yorky on 2023/2/25.
//
#include "iostream"

using namespace std;

int main() {

    int a;
    int b;

    // Exchange method one
    int tmp;
    cin >> a >> b;
    tmp = a;
    a = b;
    b = tmp;
    cout << a << " " << b;

    // Exchange method two
    cin >> a >> b;
    a = a ^ b;
    cout << a << endl;
    b = a ^ b;
    cout << b << endl;
    a = a ^ b;
    cout << a << " " << b;

    return 0;
}