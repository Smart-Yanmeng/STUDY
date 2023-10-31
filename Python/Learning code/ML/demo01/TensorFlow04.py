import tensorflow as tf
import keras
from keras import layers
from tensorflow.python.keras.losses import BinaryCrossentropy

# from keras.src.losses import BinaryCrossentropy

model = keras.Sequential([
    layers.Dense(units=25, activation="sigmoid"),
    layers.Dense(units=15, activation="sigmoid"),
    layers.Dense(units=1, activation="sigmoid")
])

model.compile(loss=BinaryCrossentropy())
