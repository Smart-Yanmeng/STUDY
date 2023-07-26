# Author - York
# Create Time - 7/26/2023 2:19 PM
# File Name - shuffle_port
# Project Name - Learning code

import random


def _shuffle(ls):
    random.shuffle(ls)
    return ls


def ShufflePort(pre_list):
    return list(zip(*[_shuffle(pre_list.copy()) for _ in range(6)]))


preList = [10000, 10001, 10002, 10003]

resultList = ShufflePort(preList)
print("port list ->", resultList)
