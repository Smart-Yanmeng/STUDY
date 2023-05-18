# Author - York
# Create Time - 5/18/2023 9:05 PM
# File Name - CreateQueue
# Project Name - example code
import queue


def create_queue():
    q = queue.Queue()
    q.put("I'm a queue")
    return q
