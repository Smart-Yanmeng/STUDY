# Author - York
# Create Time - 7/27/2023 11:49 PM
# File Name - demo01
# Project Name - temp-plot.html

# web
from flask import Flask

app = Flask(__name__)


@app.route("/")
def hello():
    return "Hello World!\nThis is a web program"


if __name__ == "__main__":
    app.run()
