# Author - York
# Create Time - 7/27/2023 12:31 AM
# File Name - demo01
# Project Name - temp-plot.html

# class
class Student:
    name = "Default"
    number = "Default"

    def __init__(self, name, number):
        self.name = name
        self.number = number

    @classmethod
    def getInfo(cls):
        print("name ->", cls.name, "\nnumber ->", cls.number)

    @staticmethod
    def printInfo(self):
        print("name ->", self.name, "\nnumber ->", self.number)


class ExcStudent(Student):
    def __init__(self, name, number):
        super().__init__(name, number)
        self.score = 90

    @classmethod
    def getInfo(cls):
        print("exc name ->", cls.name, "\nexc number ->", cls.number)


stu1 = Student("York", "01")
Student.getInfo()
Student.printInfo(stu1)

print("stu1 is a student? ->", isinstance(stu1, Student))
