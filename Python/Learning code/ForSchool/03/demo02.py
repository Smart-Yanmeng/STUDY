# Author - York
# Create Time - 7/24/2023 2:24 PM
# File Name - demo02
# Project Name - Learning code

# string
longString = """Hello World"""
print("longString ->", longString)

normalString = "Hello\nWorld"
originString = r"Hello\nWorld"
print("normalString ->", normalString)
print("originString ->", originString)

s = "This is a test."
print("the location of 'is' ->", s.find("is"))

name = "john johnson"
print("name ->", name)
print("name.title ->", name.title())
print("name.upper ->", name.upper())
print("name.lower ->", name.lower())

print("replace 'is' with 'ezz' in list s ->", s.replace("is", "ezz"))
