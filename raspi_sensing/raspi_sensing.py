from sense_hat import SenseHat

sense = SenseHat()
sense.clear()

temp_celcius = sense.get_temperature()
temp_farenheit = round((temp_celcius * 9/5) + 32,2)

pressure = sense.get_pressure()
humidity = sense.get_humidity()

print(temp_farenheit,pressure,humidity)