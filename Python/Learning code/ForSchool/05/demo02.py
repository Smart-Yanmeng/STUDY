# Author - York
# Create Time - 7/25/2023 11:53 PM
# File Name - demo02.py
# Project Name - Learning code

# dictionary
age = {'York': 18, 'Nan': 20}
print("Nan's age is ->", age['Nan'])

print("the num of the age ->", len(age))

for name in age:
    print(name + ": " + str(age[name]))

print("before del ->", age)
del age['York']
print("after del ->", age)
