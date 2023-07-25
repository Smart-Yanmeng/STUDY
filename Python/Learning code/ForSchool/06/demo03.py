# Author - York
# Create Time - 7/26/2023 1:00 AM
# File Name - demo03
# Project Name - Learning code

from math import sqrt


def dis(x1, x2, y1, y2):
    return sqrt((x1 - x2)**2 + (y1 - y2)**2)


print("dis(1, 4, 1, 5) ->", dis(1, 4, 1, 5))
print("dis(x1=1, y2=5, y1=1, x2=4) ->", dis(x1=1, y2=5, y1=1, x2=4))
