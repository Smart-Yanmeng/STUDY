#include "iostream"
#include "cstring"
#include "algorithm"

using namespace std;

#define Max 10000001

struct Computer {
    char type[25];
    char name[25];
    int price;
    int qua;
};

Computer com[1005];
char list[1005][25];

int getType(char *type, int cnt) {
    for (int i = 0; i < cnt; i++) {
        if (!strcmp(list[i], type)) return i;
    }
    strcpy(list[cnt], type);

    return cnt + 1;
}

int cmp(Computer a, Computer b) {
    return a.qua > b.qua;
}

int temp[1005];

int check(int n, int s) {
    int mark = 0, sum = 0, t;

    for (int i = 0; i < s; i++) temp[i] = Max;

    for (int i = 0; i < n; i++) {
        t = getType(com[i].type, s);
        if (temp[t] > com[i].price) temp[t] = com[i].price;
    }

    for (int i = 0; i < s; i++) {
        if (temp[i] == Max) return Max;
        else sum += temp[i];
    }

    return sum;
}

int binary(int r, int key) {
    int l = 0, mid;

    while (l < r) {
        mid = (r + l) / 2;

        if (com[mid].qua >= key) l = mid + 1;
        else r = mid;
    }

    return r;
}

int main() {
    int T;
    cin >> T;

    while (T--) {
        int n, myMoney;
        cin >> n >> myMoney;

        // 产品种类
        int listSize = 0;

        for (int i = 0; i < n; i++) {
            cin >> com[i].type >> com[i].name >> com[i].price >> com[i].qua;
            listSize = max(listSize, getType(com[i].type, listSize));
        }

        sort(com, com + n, cmp);

        int mid, l = 0, r = n - 1;

        while (l < r) {
            mid = (l + r) / 2;
            if (check(binary(n - 1, com[mid].qua), listSize) > myMoney) {
                l = mid + 1;
            } else r = mid;
        }

        cout << com[r].qua << endl;
    }

    return 0;
}