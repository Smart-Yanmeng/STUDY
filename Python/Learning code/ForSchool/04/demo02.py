# Author - York
# Create Time - 7/24/2023 5:10 PM
# File Name - demo02
# Project Name - Learning code

# while
# calculate log(2, x)
x = int(input("please enter a number:\n"))
count = 0
while x > 1:
    x /= 2
    count += 1
print("count ->", count)

# for
listA = [1, 2, 3, 4, 5]
for i in listA:
    print(listA[i - 1])
    i += 1

# judge prime number
num = int(input("please enter a number:\n"))
for i in range(num - 1, 1, -1):
    if num % i == 0:
        print("not a prime number!")
        break
else:
    print("it is a prime number!")
