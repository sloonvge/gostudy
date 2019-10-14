#!/bin/python
# -*- coding: utf8 -*-
import sys
import os
import re

#请完成下面这个函数，实现题目要求的功能
#当然，你也可以不按照下面这个模板来作答，完全按照自己的想法来 ^-^ 
#******************************开始写代码******************************


def  DynamicPathPlanning(grid, rows, cols):
    steps = []
    for i in range(rows):
        tmp = [0 for j in range(cols)]
        steps.append(tmp)
    for j in range(cols):
        steps[0][j] = 1
        if grid[0][j] == 1:
            steps[0][j] = 0
            break
    for i in range(rows):
        steps[i][0] = 1
        if grid[i][0] == 1:
            steps[i][0] = 0
            break
    for i in range(1, rows):
        for j in range(1, cols):
            if grid[i][j] == 0:
                steps[i][j] = steps[i - 1][j] + steps[i][j - 1]
            else:
                steps[i][j] = 0
    return steps[rows - 1][cols - 1]
#******************************结束写代码******************************


_matrixGrid_rows = 0
_matrixGrid_cols = 0
_matrixGrid_rows = int(input())
_matrixGrid_cols = int(input())

_matrixGrid = []
for _matrixGrid_i in range(_matrixGrid_rows):
    _matrixGrid_temp = map(int,re.split(r'\s+', input().strip()))
    _matrixGrid.append([i for i in _matrixGrid_temp])

  
res = DynamicPathPlanning(_matrixGrid, _matrixGrid_rows, _matrixGrid_cols)

print(str(res) + "\n")