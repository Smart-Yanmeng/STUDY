# Author - York
# Create Time - 7/24/2023 4:45 PM
# File Name - demo06
# Project Name - Learning code

# if
PASS = 60

score = int(input("please enter your score: "))

if score < PASS:
    print("Your grades are not up to standard!")
else:
    print("Your grades are qualified!")

print("Your grades are not up to standard!" if score > PASS else "Your grades are qualified!")
