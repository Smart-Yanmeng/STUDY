#include "iostream"

using namespace std;

struct student {
    int stuNum;
    double score;
};

int main() {
    int n, number;

    cin >> n >> number;
    student stuList[n];

    for (int i = 0; i < n; i++) cin >> stuList[i].stuNum >> stuList[i].score;

    // bubbleSort
    for (int i = n - 1; i >= 0; i--) {
        bool flag = true;

        for (int j = 0; j < i; j++) {
            if (stuList[j].score < stuList[j + 1].score) {
                student temp = stuList[j];
                stuList[j] = stuList[j + 1];
                stuList[j + 1] = temp;

                flag = false;
            }
        }

        if (flag) break;
    }

    printf("%d %g", stuList[number - 1].stuNum, stuList[number - 1].score);

    return 0;
}