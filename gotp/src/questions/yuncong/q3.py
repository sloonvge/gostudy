import sys
[m, s, t] = sys.stdin.readline().split(" ")
m = int(m)
s = int(s)
t = int(t)

for tx := 1; tx <= t; tx++:
    for ty := 1; ty <= t; ty++:
        ta := int((4 * tx + m) / 10)
        p := ta * 50 + 13 * ty
        if p > s:
            if tx + ty + ta <= 10 {
                fmt.Println(tx + ty + ta, p)
=