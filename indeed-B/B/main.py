# coding: utf-8

if __name__ == "__main__":
    s = raw_input()
    t = raw_input()

    if len(s) != len(t):
        print(-1)
    elif s == t:
        print(0)
    else:
        for i in xrange(1, len(s) + 1):
            x = s[-1 * i:] + s[:-1 * i]
            if x == t:
                print(i)
                break
        else:
            print(-1)
