#include "iostream"

using namespace std;

int main() {
    int arr[100000] = {0};
    int a;
    cin >> a;
    int s = 0;
    int d = -11111;
    int e = 0;
    int i = 0;
    for (i = 0; i < a; i++) cin >> arr[i];
    for (i = 0; i < a; i++) {
        int j = 0;
        for (j = 0; j < a - i; j++) {
            int k;
            s = 0;
            for (k = 0; k <= i; k++) {
                s += arr[j + k];
                if (s >= d && k == i) {
                    d = s;
                    e = i + 1;
                }
            }
        }
    }

    cout << "MAX Sum= " << d << ", MAX K= " << e;

    return 0;
}