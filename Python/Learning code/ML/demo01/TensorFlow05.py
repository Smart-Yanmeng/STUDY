import numpy as np
import tensorflow as tf
import keras
from keras import layers

X = np.array([[200.0, 17.0],
              [120.0, 5.0],
              [425.0, 20.0],
              [212.0, 18.0]])
Y = np.array([1, 0, 0, 1])

model = keras.Sequential([
    layers.Dense(units=25, activation="sigmoid"),
    layers.Dense(units=15, activation="sigmoid"),
    layers.Dense(units=10, activation="linear")
])

# Adam Algorithm
model.compile(optimizer=tf.keras.optimizers.Adam(learning_rate=1e-3),
              loss=tf.keras.losses.SparseCategoricalCrossentropy(from_logits=True))

model.fit(X, Y, epochs=100)
