# Author - York
# Create Time - 7/26/2023 12:49 AM
# File Name - demo01
# Project Name - Learning code

# function
def fibs(n):
    result = [0, 1]
    for i in range(n - 2):
        result.append(result[-2] + result[-1])
    return result


print(fibs(10))
