# Author - York
# Create Time - 7/25/2023 10:18 PM
# File Name - demo01
# Project Name - Learning code

# set
fruit = {'apple', 'orange', 'pear', 'banana'}
emp = set()
# list -> set
primeList = [2, 3, 5, 7, 11]
primeSet = set(primeList)
even = set([x * 2 for x in range(1, 100)])

fruits = {'apple', 'apple', 'apple', 'orange', 'orange', 'orange', 'pear', 'banana'}
print("fruits ->", fruits)

s1 = {2, 3, 5, 7}
s2 = {1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
print("s1 is s2's subset ->", s1.issubset(s2))
print("s2 is s1's superset ->", s1.issuperset(s2))
