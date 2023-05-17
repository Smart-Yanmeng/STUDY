# Author - York
# Create Time - 5/14/2023 10:44 PM
# File Name - Range
# Project Name - Multithreading.py
b = 0
a = b
s = b
for i in range(11):
    if i % 2 == 0:
        a = a + 1
    else:
        b = b + 1
    s = s + i
print(s, a, b)
