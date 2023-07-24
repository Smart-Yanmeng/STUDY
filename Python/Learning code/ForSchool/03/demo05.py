# Author - York
# Create Time - 7/24/2023 4:29 PM
# File Name - demo05
# Project Name - Learning code

import random

a = random.random()
print("random a ->", a)

b = random.uniform(0, 10)
print("random b ->", b)

c = random.randint(10, 30)
print("random c ->", c)

d = random.randrange(10, 30, 2)
print("random d ->", d)

e = random.choice([1, 2, 3, 4, 5])
print("random e ->", e)

shuffleList = [1, 2, 3, 4, 5, 6, 7, 8, 9]
print("origin list ->", shuffleList)
random.shuffle(shuffleList)
print("shuffle list ->", shuffleList)

print("sample list ->", random.sample(shuffleList, 3))
