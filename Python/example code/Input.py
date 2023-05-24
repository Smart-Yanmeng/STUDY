# Author - York
# Create Time - 5/24/2023 6:32 PM
# File Name - Input
# Project Name - Multithreading.py
def isWord(ch):
    return 'a' <= ch <= 'z' or 'A' <= ch <= 'Z'


def judgeWord(ch):
    global flag
    global res

    for index, char in enumerate(res):
        if ch == char or ch.upper() == char or ch.lower() == char:
            break
        elif ch != char and ch.upper() != char and ch.lower() != char and index == len(res) - 1:
            res += ch
            flag += 1


if __name__ == '__main__':
    res = ''
    flag = 0
    a = input()

    if len(a) < 10:
        print("not found")
    else:
        for ch in a:
            if isWord(ch):
                if len(res) == 0:
                    res += ch
                else:
                    judgeWord(ch)
                    if flag == 9:
                        print(res)
                        break
            else:
                continue

        if flag < 9:
            print("not found")
