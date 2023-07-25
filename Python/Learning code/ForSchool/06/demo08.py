# Author - York
# Create Time - 7/26/2023 2:24 AM
# File Name - demo08
# Project Name - Learning code

# built-in function
# sorted
listA = [5, 7, 6, 3, 4, 1, 2]
listB = sorted(listA)
print("origin listA ->", listA)
print("sorted listA ->", listB)

# map
print(list(map(lambda x: x ** 2, [1, 2, 3, 4, 5])))
print(list(map(lambda x, y: x + y, [1, 2, 3, 4, 5], [2, 3, 4, 5, 6])))

# zip
listB = [1, 2, 3, 4, 5]
print("listA ->", listA)
print("listB ->", listB)
print("zip(listA, listB) ->", list(zip(listA, listB)))

# eval / exec
x, y = 3, 7
print(eval('x + 3 * y - 4'))
exec('print("Hello World!")')

# all / any
n = 47
print("all ->", all([1 if n % k != 0 else 0 for k in range(2, n)]))
n = 15
print("all ->", all([1 if n % k != 0 else 0 for k in range(2, n)]))
print("any ->", any([1 if n % k != 0 else 0 for k in range(2, n)]))
