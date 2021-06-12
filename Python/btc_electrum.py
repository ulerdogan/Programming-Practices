import os
import sys
import subprocess

start = "electrum daemon -d --testnet"
create = "electrum create --testnet"
load = "electrum load_wallet --testnet"
newAddr = "electrum --testnet createnewaddress"
newAddress = ""
pay =  "electrum --testnet payto"
broadcast = "electrum --testnet broadcast -"
balance = "electrum --testnet getbalance"
ext = "electrum stop --testnet"

print("Electrum Wallet!\n")
os.system(start)
if subprocess.check_output(load, shell=True).decode("utf8").strip() != "true":
    os.system(create)

while True:
    print("\n1. Yeni adres oluşturun")
    print("2. Yeni adrese BTC gönderin")
    print("3. Bakiye sorgulayın")
    print("4. Çıkış yapın")
    inp = input("Seçiminiz: ")

    print("\n")
    if inp == "1":
        newAddress = subprocess.check_output(newAddr, shell=True).decode("utf8").strip()
        print("Yeni adresiniz:", newAddress)
    elif inp == "2":
        combine = pay + " " + newAddress + " " + input("Lütfen miktar giriniz: ") + " | " + broadcast
        os.system(combine)
    elif inp == "3":
        newbalance = subprocess.check_output(balance, shell=True).decode("utf8").strip()
        print(newbalance)
    elif inp == "4":
        os.system(ext)
        break