from sense_hat import SenseHat

sense = SenseHat()
sense.clear()

# Get the temp
temp_celcius = sense.get_temperature()

# Convert temp to Farenheit
temp_farenheit = round((temp_celcius * 9/5) + 32,2)

# Get air pressure
pressure = round(sense.get_pressure(),2)

# Get humidity
humidity = round(sense.get_humidity(),2)

# Print out all the variables
print(temp_farenheit,pressure,humidity)