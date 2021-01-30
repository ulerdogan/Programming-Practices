import urllib.request, urllib.parse, urllib.error
import xml.etree.ElementTree as ET
import ssl

# Ignore SSL certificate errors
ctx = ssl.create_default_context()
ctx.check_hostname = False
ctx.verify_mode = ssl.CERT_NONE

address = input('Enter location: ')
if len(address) < 1:
    address = "http://py4e-data.dr-chuck.net/comments_862210.xml"
    
print('Retrieving', address)

uh = urllib.request.urlopen(address)  
data = uh.read() 

print('Retrieved', len(data), 'characters')

tree = ET.fromstring(data)

lst = tree.findall('comments/comment')

print("Count:", len(lst))


total = 0

for item in lst:
    num = int(item.find('count').text)
    total = total + num

print(total)
