# To run this, download the BeautifulSoup zip file
# http://www.py4e.com/code3/bs4.zip
# and unzip it in the same directory as this file

from urllib.request import urlopen
from bs4 import BeautifulSoup
import ssl

# Ignore SSL certificate errors
ctx = ssl.create_default_context()
ctx.check_hostname = False
ctx.verify_mode = ssl.CERT_NONE

url = input('Enter - ')  # http://py4e-data.dr-chuck.net/comments_862208.html
html = urlopen(url, context=ctx).read()
soup = BeautifulSoup(html, "html.parser")

# Retrieve all of the spans
spans = soup('span')
total = 0
count = 0

for span in spans:
    # Addition at the values of a span
    total = total + int(span.contents[0])
    count += 1

print("Sum", total)
print("Count", count)
