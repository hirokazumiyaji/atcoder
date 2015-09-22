n = int(raw_input())
ps = []
for _ in xrange(n):
    p = int(raw_input())
    if p > 0:
        ps.append(p)

ps.reverse()

q = int(raw_input())
for _ in xrange(q):
    k = int(raw_input())
    if k >= len(ps):
        print(0)
    else:
        print(ps[k] + 1)
