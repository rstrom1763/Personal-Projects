from sense_hat import SenseHat
import requests
import json
import time

auth_code = ''  # Place holder for when tokens are implemented
server = 'http://maul.lan:8090/posttemp'  # The api url

sense = SenseHat()

while True:

    # Get the temp
    temp_celcius = sense.get_temperature()

    # Convert temp to Farenheit
    temp_farenheit = round((temp_celcius * 9/5) + 32, 2)

    # Get air pressure
    pressure = round(sense.get_pressure(), 2)

    # Get humidity
    humidity = round(sense.get_humidity(), 2)

    headers = {
        'Content-Type': 'application/json'
    }  # Headers for the post request

    data = {
        'auth-code': auth_code,
        'temp': temp_farenheit,
        'humidity': humidity,
        'pressure': pressure
    }

    # Send post request to the server
    requests.post(server, data=json.dumps(data), headers=headers)

    time.sleep(60)  # Wait one minute before logging the temperature again
