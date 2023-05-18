# Author - York
# Create Time - 5/18/2023 9:18 PM
# File Name - server
# Project Name - example code
import socket

sock = socket.socket()

sock.bind(('127.0.0.2', 10000))

sock.listen(5)
print("listen...")
con, addr = sock.accept()

msg = con.recv(1000)
print(msg)
