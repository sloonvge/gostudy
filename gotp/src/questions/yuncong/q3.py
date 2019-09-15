
find = False

def findPath(grid, rows, columns):
    vis = list()
    for i in range(0, rows):
        t = list()
        for j in range(0, columns):
            t.append(0)
        vis.append(t)
    hasPath(grid, rows - 1, columns - 1, 0, 0, vis)
    if find:
        print(1)
        return 1
    print(0)
    return 0

def hasPath1(grid, rows, cols, row, col, vis):
    if col < 0 or col > cols or row < 0 or row > rows or grid[row][col] == 0:
        return False
    if grid[row][col] == 9:
        return True
    return hasPath1(grid, rows, cols, row, col + 1, 
    vis) or hasPath1(grid, rows, cols, row, col -1, 
    vis) or hasPath1(grid, rows, cols, row + 1, col, 
    vis) or hasPath1(grid, rows, cols, row - 1, col, vis)

def hasPath(grid, rows, cols, row, col, vis):
    global find
    if find:
        return
    if grid[row][col] == 9:
        find = True
        return

    if col < cols and vis[row][col + 1] == 0 and grid[row][col] == 1:
        vis[row][col + 1] = 1
        hasPath(grid, rows, cols, row, col + 1, vis)
        vis[row][col + 1] = 0
    if col > 0 and vis[row][col - 1] == 0 and grid[row][col] == 1:
        vis[row][col - 1] = 1
        hasPath(grid, rows, cols, row, col - 1, vis)
        vis[row][col - 1] = 0
    if row < rows and vis[row + 1][col] == 0 and grid[row][col] == 1:
        vis[row + 1][col] = 1
        hasPath(grid, rows, cols, row + 1, col, vis)
        vis[row + 1][col] = 0
    if row > 0 and vis[row - 1][col] == 0 and grid[row][col] == 1:
        vis[row - 1][col] = 1
        hasPath(grid, rows, cols, row - 1, col,vis)
        vis[row - 1][col] = 0


rows = 3
cols = 4
grid = [
    [1, 1, 1, 1],
    [0, 1, 0, 0],
    [0, 0, 9, 1]
]

findPath(grid, rows, cols)