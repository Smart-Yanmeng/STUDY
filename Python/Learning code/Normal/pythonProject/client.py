import socket
import numpy as np
import time


def client():
    timeArr = np.array([])

    host = '127.0.0.1'  # The server's hostname or IP address
    port = 12345  # The port used by the server

    client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client_socket.connect((host, port))

    matrix = np.random.rand(1000000, 1)  # 1000000x1 matrix

    for i in range(100):
        start_time = time.time()
        client_socket.send(matrix.tobytes())
        end_time = time.time()
        ti = end_time - start_time
        timeArr = np.append(timeArr, ti)
        print(f"For {i} -> Time taken to send: {ti} seconds")

    client_socket.close()
    print(f"AVG TIME -> {np.mean(timeArr)}")


if __name__ == "__main__":
    client()
