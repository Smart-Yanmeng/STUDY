# Author - York
# Create Time - 7/27/2023 12:31 AM
# File Name - demo01
# Project Name - temp-plot.html

# class
class Student:
    def __init__(self, name, number):
        self.name = name
        self.number = number

    @classmethod
    def getInfo(cls):
        print("name ->", cls.name, "\nnumber ->", cls.number)


stu1 = Student("York", "01")
stu1.getInfo()

print("stu1 is a student? ->", isinstance(stu1, Student))
