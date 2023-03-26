#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <windows.h>
#include <time.h>

#define PI 3.14159265359

char love[20][73] = {
        "111111111111111111111111111111111111111111111111111111111111111111111111",
        "111111111111111111111111111111111100001111111111111111100001000111111111",
        "111111111111111111111100000000000000000111111111111111100011000111111111",
        "111111111111111111111100001110011100011111111111111111000110001111111111",
        "111111111111111111111111000100001100001111111111111111000110000000000011",
        "111111111111111111111111000010001000111111111111111110000100011111100011",
        "111111111111111111111000000000000000000011111111111110000100011111000111",
        "111111111111111111111000000000000000000111111111111100000000110001000111",
        "111111111111111111111000111100011111000111111111111000000000110001111111",
        "111111111111111111111110111000111111001111111111111111000111110001111111",
        "111111111111111111111110000000000000001111111111111111000110000000001111",
        "111111111111111111111111110001111111111111111111111111000110000000000111",
        "111111111111111111111111100011111111111111111111111111000100010001000111",
        "111111111111111111111111000000000000011111111111111111000000010001100011",
        "111111111111111111111110000000110000111111111111111111000000110001100001",
        "111111111111111111111100011100000001111111111111111111000001110001110111",
        "111111111111111111111000111110000111111111111111111111000111110001111111",
        "111111111111111111111111100000000000000111111111111111000111000001111111",
        "111111111111111111111110000001111100000111111111111111000111000011111111",
        "111111111111111111111111111111111111111111111111111111111111111111111111"
};

void gotoxy(int x, int y) {
    HANDLE hout = GetStdHandle(STD_OUTPUT_HANDLE);
    COORD pos = {x, y};
    SetConsoleCursorPosition(hout, pos);
}

void SetColor(int color) {
    if (color == 0) color = 0x04;
    HANDLE hCon = GetStdHandle(STD_OUTPUT_HANDLE);
    SetConsoleTextAttribute(hCon, color);
}

void showSnow() {
    SetColor(rand() % 16);
    int x = rand() % 223;
    int y = rand() % 60;
    gotoxy(x, y);
    putchar('*');
}


float sx, sy;

float sdCircle(float px, float py, float r) {
    float dx = px - sx, dy = py - sy;
    return sqrtf(dx * dx + dy * dy) - r;
}

float opUnion(float d1, float d2) {
    return d1 < d2 ? d1 : d2;
}

#define T px + scale * r* cosf(theta), py + scale * r * sin(theta)

int ribbon() {
    float x = (fmodf(sy, 0.1f) / 0.1f - 0.5f) * 0.5f;
    return sx >= x - 0.05f && sx <= x + 0.05f;
}

float f(float px, float py, float theta, float scale, int n) {
    float d = 0.0f;
    for (float r = 0.0f; r < 0.8f; r += 0.02f)
        d = opUnion(d, sdCircle(T, 0.05f * scale * (0.95f - r)));
    if (n > 0)
        for (int t = -1; t <= 1; t += 2) {
            float tt = theta + t * 1.8f;
            float ss = scale * 0.9f;
            for (float r = 0.2f; r < 0.8f; r += 0.1f) {
                d = opUnion(d, f(T, tt, ss * 0.5f, n - 1));
                ss *= 0.8f;
            }
        }
    return d;
}

int main(int argc, char *argv[]) {

    srand((int) time(nullptr));
    SetConsoleCP(437);
    SetConsoleOutputCP(437);

    int n = argc > 1 ? atoi(argv[1]) : 3;
    float zoom = argc > 2 ? atof(argv[2]) : 1.0f;
    for (sy = 0.8f; sy > 0.0f; sy -= 0.02f / zoom, putchar('\n')) {
        for (sx = -0.35f; sx < 0.35f; sx += 0.01f / zoom) {
            if (f(0, 0, PI * 0.5f, 1.0f, n) < 0.0f) {
                if (sy < 0.1f) {
                    SetColor(rand() % 16);
                    putchar('\3');
                } else {
                    if (ribbon()) {
                        putchar('=');
                    } else {
                        SetColor(rand() % 16);
                        putchar('\3');
                    }
                }
            } else {
                putchar(' ');
            }
        }
    }


    int i, j;
    int x, y = 3;
    for (i = 0; i < 19; i++, y++) {
        x = 60;
        for (j = 0; j < 72; j++, x++) {
            gotoxy(x, y);
            if (love[i][j] == '1') {
                putchar(' ');
            } else {
                SetColor(0x06);
                putchar('*');
            }
        }
    }

    for (n = 0; n++ < 365;) {
        showSnow();
    }
//    gotoxy(0, 100);
    getchar();

    return 0;
}