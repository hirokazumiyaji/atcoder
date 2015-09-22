if __name__ == "__main__":
    string = "ddeeinnow"
    n = int(raw_input())

    for i in xrange(n):
        sn = raw_input()
        if ''.join(sorted([s for s in sn])) == string:
            print("YES")
        else:
            print("NO")
