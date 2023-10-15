#include <cstring>
#include "iostream"

using namespace std;


int main() {
    int arr[42], tmp = 0, res = 0;
    memset(arr, 0, sizeof arr);
    for (int i = 0; i < 10; i++) {
        cin >> tmp;
        arr[tmp % 42] = 1;
    }

    for (int i: arr) res += i;
    cout << res;

    return 0;
}