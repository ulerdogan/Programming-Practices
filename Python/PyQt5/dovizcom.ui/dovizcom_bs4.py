import requests
from bs4 import BeautifulSoup

url = "https://www.doviz.com/"

response = requests.get(url)
html_icerigi = response.content
soup = BeautifulSoup(html_icerigi, "html.parser")

def kurBul(c):
    for doviz in soup.find_all("div", {"class": "item"}):
        if doviz.find("span", {"class": "name"}).text == c:
            ck = doviz.find("span", {"class": "value"}).text

    kur = float(ck.split(",")[0] + "." + ck.split(",")[1])
    return kur

print(kurBul("DOLAR"))
