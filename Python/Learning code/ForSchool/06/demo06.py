# Author - York
# Create Time - 7/26/2023 1:55 AM
# File Name - demo06
# Project Name - Learning code

listA = [2, 3, 5, 8, 3, 6, 8]
seen = set()
# seen.add(*) always return none, not none always be ture
listB = [i for i in listA if i not in seen and not seen.add(i)]
print(listB)
