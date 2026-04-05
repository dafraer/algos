import sys

t = int(input())

for t1 in range(t):
    n = int(input())
    num = 1
    for j in range(n):
        sys.stdout.write(f"{num} ")
        num *= 2
    sys.stdout.write("\n")
