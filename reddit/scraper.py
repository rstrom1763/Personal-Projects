import requests
import json
import time

test_users = ['Ryan1763']#,'MADTATER01','wackypaw123']

for user in test_users:
    file = open((user + '.json'),'w')
    data = requests.get('https://reddit.com/u/' + user + '.json',headers = {'User-agent': 'Ryan1763'}).content.decode()
    file.write(data)

    data = json.loads(data)
    for comment in data['data']['children']:
        if comment['kind'] == 't1':
            print(comment['data']['body'] + '\n')

    time.sleep(.1)