#include "iostream"

using namespace std;

int main() {
    int n, k;
    int m, num = 0;
    cin >> n >> k;
    num += n;

    while (n >= k) {
        // 换的水
        m = n / k;
        if (m < 1) {
            cout << num;
            return 0;
        }
        num += m;

        // 余下的水
        n = n % k;
        n += m;
    }

    cout << num;

    return 0;
}