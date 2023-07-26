# Author - York
# Create Time - 7/27/2023 12:11 AM
# File Name - demo03
# Project Name - temp-plot.html

import plotly as py
import plotly.graph_objs as pygo
import pandas as pd

pyplt = py.offline.plot
data = pd.read_csv("student.csv", encoding="GBK")

xdata = data["name"].tolist()
ydata = data["score"].tolist()

trace0 = pygo.Scatter(x=xdata, y=ydata, name="score")

score = pygo.Data([trace0])
layout = pygo.Layout(
    title="scores",
)

dataDraw = pygo.Figure(data=score, layout=layout)
pyplt(dataDraw)
