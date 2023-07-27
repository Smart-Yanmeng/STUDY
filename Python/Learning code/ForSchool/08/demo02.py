# Author - York
# Create Time - 7/27/2023 11:06 PM
# File Name - demo02
# Project Name - temp-plot.html

class Personal:
    def __init__(self, name, age):
        self.__name = name
        self.__age = age

    @property
    def name(self):
        """I'm the 'x' property."""
        return self.__name


p = Personal("York", 18)
print(p.name)
