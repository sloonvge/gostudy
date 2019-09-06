import sys

lines = sys.stdin.readlines()
s1 = lines[0]
s2 = lines[1]

def isSubSequence(s1, s2):
    i = 0
    j = 0
    while i < len(s1) and j < len(s2):
        if s1[i] == s2[j]:
            j += 1
        i += 1
    return j == len(s2)

ans = isSubSequence(s1, s2)
if ans:
    print(1)
else:
    print(0)

