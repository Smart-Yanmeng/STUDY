# Author - York
# Create Time - 7/23/2023 12:38 AM
# File Name - demo06
# Project Name - Learning code

# format
lower, upper = input().split()
lower, upper = int(lower), int(upper)
for i in range(lower, upper, 2):
    print(i, "{:.1f}".format(5 * (i - 32) / 9))

s = "hello world!"

# default - align left
print("{0}".format(s))

# minimum wide 30 - align right
print("{0:>30}".format(s))

# minimum wide 30 - in the middle
print("{:^30}".format(s))
