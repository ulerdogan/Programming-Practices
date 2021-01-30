import sys
from PyQt5 import  QtWidgets


class Pencere(QtWidgets.QWidget):
    def __init__(self):

        super().__init__()

        self.init_ui()

    def init_ui(self):

        self.yazı_alanı = QtWidgets.QLineEdit() #input alanı oluşturma
        self.temizle = QtWidgets.QPushButton("Temizle") #temizle ve yazdır butonları
        self.yazdır = QtWidgets.QPushButton("Yazdır")

        v_box = QtWidgets.QVBoxLayout()

        v_box.addWidget(self.yazı_alanı) # oluşturduğumuz şeyleri vboxa ekleme
        v_box.addWidget(self.temizle)
        v_box.addWidget(self.yazdır)
        v_box.addStretch()

        self.setLayout(v_box)


        self.temizle.clicked.connect(self.click) #fonksiyanlar bağlama
        self.yazdır.clicked.connect(self.click)



        self.show()

    def click(self):
        sender = self.sender() #miras aldığım sender fonksiyonu üstündeki yazıyı almayı sağlar, hangi butona basıldığını alırız

        if sender.text() == "Temizle":
            self.yazı_alanı.clear()

        else:
            print(self.yazı_alanı.text()) #yazı alanındaki texti alma




app = QtWidgets.QApplication(sys.argv)

pencere = Pencere()

sys.exit(app.exec_())