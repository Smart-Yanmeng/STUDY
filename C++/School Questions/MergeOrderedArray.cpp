#include "iostream"

using namespace std;

int main() {
    int fArrLong, sArrLong;

    cin >> fArrLong;
    int fArr[fArrLong];
    for (int i = 0; i < fArrLong; i++) cin >> fArr[i];

    cin >> sArrLong;
    int sArr[sArrLong];
    for (int i = 0; i < sArrLong; i++) cin >> sArr[i];

    int rstArr[fArrLong + sArrLong];

    int fPoint = 0, sPoint = 0;
    int fArrLongTemp = fArrLong, sArrLongTemp = sArrLong;

    while (fArrLongTemp && sArrLongTemp) {
        if (fArr[fPoint] <= sArr[sPoint]) {
            rstArr[fPoint + sPoint] = fArr[fPoint];

            fPoint++;
            fArrLongTemp--;
        } else {
            rstArr[fPoint + sPoint] = sArr[sPoint];

            sPoint++;
            sArrLongTemp--;
        }
    }

    if (fArrLongTemp > 0) {
        for (int i = 0; i < fPoint + sPoint; i++) cout << rstArr[i] << " ";
        for (int i = fPoint; i < fArrLong; i++) cout << fArr[i] << " ";
    } else if (sArrLongTemp > 0) {
        for (int i = 0; i < fPoint + sPoint; i++) cout << rstArr[i] << " ";
        for (int i = sPoint; i < sArrLong; i++) cout << sArr[i] << " ";
    } else {
        for (int i = 0; i < fPoint + sPoint; i++) cout << rstArr[i] << " ";
    }

    return 0;
}