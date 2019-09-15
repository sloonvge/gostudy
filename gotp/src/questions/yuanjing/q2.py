import sys

m, n = [int(i) for i in sys.stdin.readline().split(",")]
lines = sys.stdin.readlines()
grid = []
for line in lines:
    rows = []
    for x in line.rstrip().split(","):
        rows.append(int(x))
    grid.append(rows)

def maxBoard(grid, m, n):
    answer = 0
    vis = []
    for i in range(m):
        t = []
        for j in range(n):
            t.append(0)
        vis.append(t)
    for i in range(m):
        for j in range(n):
            if vis[i][j] == 0 and grid[i][j] == 1:
                n = board(grid, m - 1, n - 1, i, j, vis)
                if n > answer:
                    answer = n
    
    return answer
def board(grid, m, n, i, j, vis):
    if i < 0 or i > m or j < 0 or j > n or vis[i][j] != 0 or grid[i][j] != 1:
        return 0
    vis[i][j] = True

    return 1 + board(grid, m, n, i - 1, j, vis) + \
    board(grid, m, n, i + 1, j, vis) + \
    board(grid, m, n, i, j - 1, vis) + \
    board(grid, m, n, i, j + 1, vis)

answer = maxBoard(grid, m, n)
print(answer)
