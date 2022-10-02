import requests
import json
import time
import os

os.system('cls')
endpoint = "http://maul.lan:8081/uploadComment"

file_path = "Y:/reddit/data/RC_2018-10.json"
count = 0
start_at = 0

with open(file_path) as comments:
    while True:
        comment = comments.readline()
        #For testing purposes
        if not comment:
            break
        if count >= start_at:
            comment = json.loads(comment)
            request = requests.post(endpoint, json=comment)
            #time.sleep(.05)
            #print(comment['author'])
        count += 1
        if count % 100 == 0 and count > start_at:
            print(count)
            time.sleep(.1)
        
