#include "iostream"
#include "cmath"

using namespace std;

int main() {
    int n, i;
    cin >> n;

    for (i = 2; i <= sqrt(n); i++) {
        while (n % i == 0) {
            printf("%d", i);
            n = n / i;
            if (n != 1) cout << "*";
        }
    }

    if (n != 1) cout << n;

    return 0;
}