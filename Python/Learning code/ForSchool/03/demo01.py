# Author - York
# Create Time - 7/24/2023 2:06 PM
# File Name - demo01
# Project Name - Learning code

# list
prompt = 'hello'
print(prompt[0])
print(prompt[-1])

listA = [1, 2, 3, 4, 5]
print("list_a[2 to 3] ->", listA[1:3])
print("list_a[2 to 3] ->", listA[1:-2])
print("list_a[2 to last] ->", listA[1:])
print("list_a[last n] ->", listA[-2:])

listB = listA[:]
print("list_b ->", listB)

print("list_a + list_b ->", listA + listB)
print("list_a * 2 ->", listA * 2)

print("judge whether 1 is in the list_a ->", 1 in listA)

print("the length of list_a ->", len(listA))
print("the minimum element of list_a ->", min(listA))
print("the maximum element of list_a ->", max(listA))

nameList = ['Alice', 'Kim', 'Karl', 'John']
del nameList[2]
print("after delete ->", nameList)

listA.append(6)
print("after append ->", listA)
listA.extend([7, 8, 9])
print("after extend ->", listA)
listA.insert(2, 2)
print("after insert ->", listA)
listA.remove(2)
print("after remove ->", listA)
# no parameter means the last element in the list
print("listA.pop ->", listA.pop())
listA.reverse()
print("after reverse ->", listA)
print("the element '3' exits in the list of ->", listA.index(3))
