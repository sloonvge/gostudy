import sys

line = sys.stdin.readline()
numbers = [int(n) for n in line.split(",")]

ans = 0
dict = {}
for n in numbers:
    dict[n] = 1
for n in numbers:
    if n-1 in dict:
        continue
    t = n
    while True:
        if t not in dict:
            break
        del dict[t]
        t += 1
    d = t - n
    if d > ans:
        ans = d

print(ans)

