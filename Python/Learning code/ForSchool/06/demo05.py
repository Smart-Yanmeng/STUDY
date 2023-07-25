# Author - York
# Create Time - 7/26/2023 1:28 AM
# File Name - demo05
# Project Name - Learning code

# variable length number parameter
def countNum(a, *b):
    print("a ->", a)
    print("b ->", b)
    print("b's length ->", len(b))


countNum(1, 2, 3, 4, 5)

# calculate the number of letters in each word in a sentence
s = "It is great to see you here"
print("the number of letters in each word in s ->", *[len(word) for word in s[0:-1].split()])


def countNum(a, **d):
    print("a ->", a)
    print("d ->", d)
    print(len(d) + 1)


countNum(1, x1=2, x2=3, x3=4, x4=5)
