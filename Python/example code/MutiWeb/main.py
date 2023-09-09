from flask import Flask, render_template
import threading
import time

app = Flask(__name__)


# 一个简单的任务，模拟耗时操作
def time_consuming_task():
    time.sleep(5)  # 模拟耗时操作，睡眠5秒
    return "Task completed!"


@app.route('/')
def index():
    return "Hello, Flask!"


@app.route('/run_task')
def run_task():
    # 创建一个线程来执行任务
    task_thread = threading.Thread(target=time_consuming_task)
    task_thread.start()
    print(threading.active_count())
    return "Task started in a separate thread."


if __name__ == '__main__':
    app.run()
