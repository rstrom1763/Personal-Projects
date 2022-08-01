from sense_hat import SenseHat

sense = SenseHat()
sense.clear()

temp_celcius = sense.get_temperature()
temp_farenheit = round((temp_celcius * 9/5) + 32,2)

pressure = round(sense.get_pressure(),2)
humidity = round(sense.get_humidity(),2)

print(temp_farenheit,pressure,humidity)