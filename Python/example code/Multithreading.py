# coding=utf-8
# Author - York
# Create Time - 5/13/2023 2:12 PM
# File Name - Multithreading
# Project Name - Python
import _thread
import time


# 为线程定义一个函数
# def print_time(threadName, delay):
#     count = 0
#     while count < 5:
#         time.sleep(delay)
#         count += 1
#         print ("%s: %s" % (threadName, time.ctime(time.time())))
#
#
# # 创建两个线程
# try:
#     _thread.start_new_thread(print_time, ("Thread-1", 2,))
#     _thread.start_new_thread(print_time, ("Thread-2", 4,))
# except:
#     print ("Error: 无法启动线程")
#
# while 1:
#     pass


import gramine

def print_message(message):
    print(message)

# 创建线程对象并设置目标函数和参数
thread1 = gramine.Thread(target=print_message, args=("Hello from Thread 1",))
thread2 = gramine.Thread(target=print_message, args=("Hello from Thread 2",))

# 启动线程
thread1.start()
thread2.start()

# 等待线程完成
thread1.join()
thread2.join()