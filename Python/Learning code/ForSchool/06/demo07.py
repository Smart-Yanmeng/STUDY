# Author - York
# Create Time - 7/26/2023 2:14 AM
# File Name - demo07
# Project Name - Learning code

# recursion function
def fib(n):
    if n == 0 or n == 1:
        return 1
    else:
        return fib(n - 1) + fib(n - 2)


print("fib(10) ->", fib(10))

# dictionary - reduce redundancy
pre = {0: 1, 1: 1}


def fibDic(n):
    if n in pre:
        return pre[n]
    else:
        new_value = fibDic(n - 1) + fibDic(n - 2)
        pre[n] = new_value
        return new_value


print(fibDic(100))
