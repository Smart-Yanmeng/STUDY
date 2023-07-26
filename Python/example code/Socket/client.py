# Author - York
# Create Time - 5/18/2023 9:19 PM
# File Name - client
# Project Name - example code
import socket
from multiprocessing import Queue

myQueue = Queue(1)
myQueue.put_nowait('app')
sock = socket.socket()

sock.connect(('39.105.219.78', 1000))

sock.send(myQueue.get_nowait().encode('utf-8'))

sock.close()
