import sys

from PyQt5.QtWidgets import QWidget,QApplication,QCheckBox,QLabel,QPushButton,QVBoxLayout,QLineEdit,QHBoxLayout, QVBoxLayout, QRadioButton

import dovizcom_bs4

class Pencere(QWidget):
    def __init__(self):
        super().__init__()

        self.init_ui()

    def init_ui(self):
        
        self.money_in_TRY = QLineEdit()
        self.money_in_TRY.setPlaceholderText("TRY")
        self.TRY = QLabel(" ₺ ")
        self.buton = QPushButton("Cevir")

        self.dolar = QRadioButton("Dolar (USD)")
        self.euro = QRadioButton("Euro (EUR)")
        self.sterlin = QRadioButton("Sterlin (GBP")

        self.yazi_alani = QLabel("")
        self.yazi_alani2 = QLabel("")



        h_box1 = QHBoxLayout()
        h_box2 = QHBoxLayout()

        h_box1.addWidget(self.money_in_TRY)
        h_box1.addWidget(self.TRY)
        h_box1.addStretch()
        h_box1.addWidget(self.buton)

        h_box2.addWidget(self.dolar)
        h_box2.addWidget(self.euro)
        h_box2.addWidget(self.sterlin)

        v_box = QVBoxLayout()
        v_box.addLayout(h_box1)
        v_box.addLayout(h_box2)
        v_box.addWidget(self.yazi_alani)
        v_box.addWidget(self.yazi_alani2)
        


        self.setLayout(v_box)

        self.buton.clicked.connect(lambda : self.click(self.dolar.isChecked(), self.euro.isChecked(), self.sterlin.isChecked(), self.yazi_alani))


        self.setWindowTitle("doviz.com Kur Çevirici")
        self.show()


    def click(self, dolar, euro, sterlin, yazi_alani):
        try:
            anapara = float(self.money_in_TRY.text())
            if dolar:
                guncel_kur = dovizcom_bs4.kurBul("DOLAR")
                self.yazi_alani.setText("İşlem Yapılan Güncel Kur : " + str(guncel_kur))
                self.yazi_alani2.setText("İşlem Sonu Elde Edilen Tutar : " + str(anapara/guncel_kur))

            if euro:
                guncel_kur = dovizcom_bs4.kurBul("EURO")
                self.yazi_alani.setText("İşlem Yapılan Güncel Kur : " + str(guncel_kur))
                self.yazi_alani2.setText("İşlem Sonu Elde Edilen Tutar : " + str(anapara/guncel_kur))

            if sterlin:
                guncel_kur = dovizcom_bs4.kurBul("STERLİN")
                self.yazi_alani.setText("İşlem Yapılan Güncel Kur : " + str(guncel_kur))
                self.yazi_alani2.setText("İşlem Sonu Elde Edilen Tutar : " + str(anapara/guncel_kur))
        except:
            self.yazi_alani.setText("Bir Hata Oluştu!")
            self.yazi_alani2.setText("Lütfen Tekrar Deneyiniz...")




app = QApplication(sys.argv)

pencere = Pencere()

sys.exit(app.exec_())