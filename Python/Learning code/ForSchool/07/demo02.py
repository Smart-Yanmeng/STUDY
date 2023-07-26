# Author - York
# Create Time - 7/26/2023 4:47 PM
# File Name - demo02
# Project Name - Learning code

# Pandas and Plotly
import pandas as pd
from plotly import figure_factory as FF
import plotly as py

pyplt = py.offline.plot

data1 = pd.DataFrame([["01", "York", "computer", "18", "90"],
                      ["02", "Mike", "computer", "18", "90"],
                      ["03", "Lucy", "computer", "18", "100"],
                      ["04", "Nan", "computer", "20", "80"],
                      ["05", "Lin", "computer", "18", "70"]],
                     columns=["number", "name", "subject", "age", "score"])

table1 = FF.create_table(data1)
pyplt(table1, show_link=False)

data2 = pd.DataFrame({"number": ["01", "02", "03", "04", "05"],
                      "name": ["York", "Mike", "Lucy", "Nan", "Lin"],
                      "subject": ["computer", "computer", "computer", "computer", "computer"],
                      "age": ["18", "18", "18", "20", "18"],
                      "score": ["90", "90", "100", "80", "70"]})

table2 = FF.create_table(data2)
pyplt(table2, show_link=False)

print("student1's number ->", data2.at[0, "number"])
print("student1's detail ->\n", data2.iloc[0], sep='')
print("all the students' number ->\n", data2["number"], sep='')
