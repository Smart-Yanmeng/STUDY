# Author - York
# Create Time - 7/26/2023 1:13 AM
# File Name - demo04
# Project Name - Learning code

# default value parameter
# default argument value is mutable
def init(arg, result=[]):
    result.append(arg)
    return result


print(init(1))
print(init(2))
print(init(3))
print(init(234, [2, 3, 4]))
print(init(345))
