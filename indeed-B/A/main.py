# coding: utf-8

if __name__ == "__main__":
    x1, y1 = [int(i) for i in raw_input().split()]
    x2, y2 = [int(i) for i in raw_input().split()]

    count = abs(x1 - x2) + abs(y1 - y2)
    count += 1

    print(count)
