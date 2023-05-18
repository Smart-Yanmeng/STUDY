# Author - York
# Create Time - 5/18/2023 9:11 AM
# File Name - read
# Project Name - Multithreading.py
import json

with open('read.json', 'r') as read_file:
    read_data = json.load(read_file)
    print(read_data)

with open('write.txt', 'w') as write_file:
    write_file.write('Hello, I\'m python!\n')
    write_file.write('This is what I read:\n')
    write_file.write(json.dumps(read_data))
