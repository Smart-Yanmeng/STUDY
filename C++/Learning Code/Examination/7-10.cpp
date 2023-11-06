#include "iostream"

using namespace std;

const int maxn = 200;
int n, m, sx, sy, ex, ey, ans = 0x3f3f3f3f;
char map[maxn][maxn];
bool vis[maxn][maxn];

struct Node {
    int x, y, step;
} q[maxn * maxn];

int dir[4][2] = {{0,  1},
                 {0,  -1},
                 {1,  0},
                 {-1, 0}};

void bfs(int x, int y) {
    int head = 0, tail = 1;
    q[head].x = x;
    q[head].y = y;
    q[head].step = 0;
    vis[x][y] = true;

    while (head < tail) {
        Node now = q[head++];
        if (now.x == ex && now.y == ey) {
            ans = min(ans, now.step);
            return;
        }

        for (int i = 0; i < 4; i++) {
            int nx = now.x + dir[i][0];
            int ny = now.y + dir[i][1];

            if (nx >= 0 && nx < n && ny >= 0 && ny < m && !vis[nx][ny] && map[nx][ny] != '#') {
                vis[nx][ny] = true;
                q[tail].x = nx;
                q[tail].y = ny;
                q[tail].step = now.step + 1;
                tail++;
            }
        }
    }
}

int main() {
    cin >> n >> m;

    for (int i = 0; i < n; i++) {
        cin >> map[i];
        for (int j = 0; j < m; j++) {
            if (map[i][j] == 'a') {
                sx = i;
                sy = j;
            } else if (map[i][j] == 'r') {
                ex = i;
                ey = j;
            }
        }
    }

    bfs(sx, sy);

    if (ans == 0x3f3f3f3f) cout << "No way!" << endl;
    else cout << ans + 1 << endl;

    return 0;
}