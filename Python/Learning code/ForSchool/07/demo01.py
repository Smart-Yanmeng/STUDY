# Author - York
# Create Time - 7/26/2023 4:35 PM
# File Name - demo01
# Project Name - Learning code

# file
textFile = open("07.txt", "rt")
t = textFile.readline()
print("t ->", t)
textFile.close()

binFile = open("07.txt", "rb")
b = binFile.readline()
print("b ->", b)
binFile.close()
