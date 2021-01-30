import json
import urllib.request, urllib.parse, urllib.error

address = input('Enter location: ') #http://py4e-data.dr-chuck.net/comments_862211.json

uh = urllib.request.urlopen(address)
data = uh.read().decode()

print(data)

print('Retrieved', len(data), 'characters')

info = json.loads(data)
print('User count:', len(info))


total = 0

for item in info["comments"]:
    temp = int(item["count"])
    total += temp

print(total)
