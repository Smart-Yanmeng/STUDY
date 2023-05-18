# Author - York
# Create Time - 5/18/2023 9:08 PM
# File Name - get_queue
# Project Name - example code
from create_queue import create_queue


def get_queue(myQueue):
    print(myQueue.get().encode('utf-8'))


myQueue = create_queue()
get_queue(myQueue)
