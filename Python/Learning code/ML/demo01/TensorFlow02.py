import numpy as np
import tensorflow as tf
import keras
from keras import layers
from keras.src.losses import BinaryCrossentropy

x = np.array([[200.0, 17.0],
              [120.0, 5.0],
              [425.0, 20.0],
              [212.0, 18.0]])
y = np.array([1, 0, 0, 1])

# layer_1 = Dense(units=3, activation="sigmoid")
# layer_2 = Dense(units=1, activation="sigmoid")

model = keras.Sequential([
    layers.Dense(units=25, activation="ReLU"),
    layers.Dense(units=15, activation="ReLU"),
    layers.Dense(units=1, activation="sigmoid")
])

model.compile(loss=BinaryCrossentropy())
model.fit(x, y, epochs=1000)
