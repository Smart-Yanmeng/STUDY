#include<iostream>
#include<malloc.h>
#include<stdio.h>
#include<math.h>

void radixSort(int *a, int N, int num) {
    int baseNum = 1;

    for (int i = 0; i < num; i++) {
        int past[10000];
        int number[10] = {0};

        for (int i = 0; i < N; i++) number[((a[i] / baseNum) % 10)]++;

        for (int i = 1; i < N; i++)
            number[i] += number[i - 1];

        for (int i = N - 1; i >= 0; i--) {
            past[number[((a[i] / baseNum) % 10)] - 1] = a[i];
            number[((a[i] / baseNum) % 10)]--;
        }

        for (int i = 0; i < N; i++) printf("%d ", past[i]);

        printf("\n");
        for (int i = 0; i < N; i++) a[i] = past[i];

        baseNum *= 10;
    }
}

int main() {
    int N;
    scanf("%d", &N);
    int *a = new int[N];

    int max = 0;
    for (int i = 0; i < N; i++) {
        scanf("%d", &a[i]);
        if (a[i] > max)
            max = a[i];
    }

    int num = 0;
    while (max != 0) {
        max = max / 10;
        num++;
    }

    radixSort(a, N, num);
}
