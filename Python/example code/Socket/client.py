# Author - York
# Create Time - 5/18/2023 9:19 PM
# File Name - client
# Project Name - example code
import socket
from multiprocessing import Queue

myQueue = Queue(1)
myQueue.put_nowait('app')
sock = socket.socket()

sock.connect(('127.0.0.2', 10000))

sock.send(myQueue.get_nowait().encode('utf-8'))

sock.close()
