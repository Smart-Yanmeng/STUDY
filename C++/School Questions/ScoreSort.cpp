#include "iostream"
#include "algorithm"

using namespace std;

struct stuSco {
    int stuNum;
    int chiSco;
    int matSco;
};

bool compare(const stuSco &a, const stuSco &b) {
    if (a.chiSco > b.chiSco) return true;
    else if (a.chiSco == b.chiSco && a.matSco > b.matSco) return true;
    else if (a.chiSco == b.chiSco && a.matSco == b.matSco && a.stuNum < b.stuNum) return true;
    else return false;
}

int main() {
    int n = 0;
    scanf("%d", &n);
    stuSco sco[n];

    for (int i = 0; i < n; i++) scanf("%d %d %d", &sco[i].stuNum, &sco[i].chiSco, &sco[i].matSco);

    sort(sco, sco + n, compare);

    for (int i = 0; i < n; i++) printf("%d\n", sco[i].stuNum);

    return 0;
}