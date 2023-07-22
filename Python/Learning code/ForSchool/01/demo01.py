# Author - York
# Create Time - 7/22/2023 1:11 AM
# File Name - demo01
# Project Name - Learning code
import math

# input
a = input("please enter an number: ")

# print
print("a ->", a)

# split
a0, a1 = input("please enter two numbers: ").split()
print("a0 ->", a0)
print("a1 ->", a1)

# print in the same line
print("a0 ->", a0, end=' ')
print("a1 ->", a1)

# calculate the triangle area
b0, b1, b2 = input("please enter three numbers: ").split()
b0 = int(b0)
b1 = int(b1)
b2 = int(b2)
s = (b0 + b1 + b2) / 2

area = math.sqrt(s * (s - b0) * (s - b1) * (s - b2))
print("the area of this triangle is:", area)
