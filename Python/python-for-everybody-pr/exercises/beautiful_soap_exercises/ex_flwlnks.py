# To run this, download the BeautifulSoup zip file
# http://www.py4e.com/code3/bs4.zip
# and unzip it in the same directory as this file

import urllib.request, urllib.parse, urllib.error
from bs4 import BeautifulSoup
import ssl

# Ignore SSL certificate errors
ctx = ssl.create_default_context()
ctx.check_hostname = False
ctx.verify_mode = ssl.CERT_NONE


# Retrieve the members
counter = 0
big_counter = 0
catched = ""

url = input('Enter - ')
count = int(input("Enter count: "))
position = int(input("Enter position: "))



while True:
    big_counter += 1
    html = urllib.request.urlopen(url, context=ctx).read()
    soup = BeautifulSoup(html, 'html.parser')
    print("Retrieving:" , url)
    members = soup('a')

    for member in members:
        counter += 1
        if counter == position:
            catched = member.contents
            url = member.get("href", None)
            counter = 0
            break

    if big_counter == count:
        print(catched)
        break
    