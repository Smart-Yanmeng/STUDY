import numpy as np

X = np.array([[200, 17]])
W = np.array([[1, -3, 5],
              [-2, 4, -6]])
B = np.array([[-1, 1, 2]])


def sig(x):
    return 1 / (1 + np.exp(-x))


def dense(a_in, w, b):
    Z = np.matmul(a_in, w) + b
    A_out = sig(Z)
    return A_out


print(dense(X, W, B))
