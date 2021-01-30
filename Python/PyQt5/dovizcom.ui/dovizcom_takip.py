import requests
from bs4 import BeautifulSoup

url = "https://www.doviz.com/"

response = requests.get(url)
html_icerigi = response.content
soup = BeautifulSoup(html_icerigi, "html.parser")


while True:
    print('''
    *************************
    Kur Seçenekleri:
    (Girdi olarak numaralarını yazabilirsiniz.)
    1. Dolar (USD)
    2. Euro (EUR)
    3. Sterlin (GBP)2
    *************************
    ''')

    cikti = input("Dönüştürmek istediğiniz kuru giriniz: ")

    try:
        miktar = float(input("Dönüştürmek istediğiniz miktarı giriniz: "))
    except:
        print("Bir hata oluştu, lütfen tekrar deneyiniz...")

    c = None
    ck = 1
    
    if cikti == "1" or cikti == "Dolar" or cikti == "USD":
        c = "DOLAR"
    elif cikti == "2" or cikti == "Euro" or cikti == "EUR":
        c = "EURO"
    elif cikti == "3" or cikti == "Sterlin" or cikti == "GBP":
        c = "STERLİN"
    else:
        print("Bir hata oluştu, lütfen tekrar deneyiniz...")
        continue

    for doviz in soup.find_all("div", {"class": "item"}):
        if doviz.find("span", {"class": "name"}).text == c:
            ck = doviz.find("span", {"class": "value"}).text


    kur = float(ck.split(",")[0] + "." + ck.split(",")[1])

    print("Güncel {} kuru: {}".format(c, kur))

    donusmus_miktar = miktar*kur
    print("Çeviri işlemi sonunda elde edeceğiniz miktar: ", donusmus_miktar)
