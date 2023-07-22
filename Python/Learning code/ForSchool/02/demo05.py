# Author - York
# Create Time - 7/22/2023 11:48 PM
# File Name - demo05
# Project Name - Learning code

# traverse and output the elements in the array
arr = [1, 2, 3, 4]
for i in arr:
    print(i)

# flashback output
for i in arr[::-1]:
    print(i)

# range
print("from 0 to 10 step 1 ->", list(range(0, 11, 1)))
print("from 10 to 0 step -2 ->", list(range(10, -1, -2)))

# sum
print("the sum of 0 to 9 ->", sum(list(range(0, 10))))

# list derivation
print([numbers for numbers in list(range(0, 5))])
print([numbers * 2 for numbers in list(range(0, 5))])
