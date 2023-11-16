import socket


def server():
    host = '127.0.0.1'  # Localhost
    port = 12345  # Arbitrary non-privileged port

    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_socket.bind((host, port))
    server_socket.listen(1)
    print("Server is listening on port", port)

    conn, addr = server_socket.accept()
    print('Connected by', addr)

    try:
        while True:
            data = conn.recv(4096)
            if not data:
                break
            # Process data if needed
    finally:
        conn.close()
        server_socket.close()


if __name__ == "__main__":
    server()
