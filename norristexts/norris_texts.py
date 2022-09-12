import json
import requests

account_sid = ''
auth_token = ''
from_num = ''

joke = json.loads(requests.get(
    'https://api.chucknorris.io/jokes/random').content.decode())


def send_text(body, to_num, from_num, account_sid, auth_token):
    from twilio.rest import Client
    client = Client(account_sid, auth_token)

    message = client.messages \
                    .create(
                        body=body,
                        from_=from_num,
                        to=to_num
                    )
    print(message.sid)


num_dict = {}

for num in num_dict.values():
    try:
        send_text(body=joke['value'], to_num=num, from_num=from_num,
              account_sid=account_sid, auth_token=auth_token)
    except:
        print("There was an error\n")
        continue

log_file = open('./norris_log.txt', 'a')
log_file.write(joke['value'] + '\n')
log_file.close()