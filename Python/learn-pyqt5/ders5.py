import sys

from PyQt5 import QtWidgets

class Pencere(QtWidgets.QWidget):  #qwidgettan miras aldık

    def __init__(self):

        super().__init__() #qwidgetın init fonksiyonunu tanıttık

        self.init_ui()
    def  init_ui(self):

        self.yazı_alanı = QtWidgets.QLabel("Bana henüz tıklanmadı..")
        self.buton = QtWidgets.QPushButton("Bana Tıkla")
        self.say = 0


        v_box = QtWidgets.QVBoxLayout() #vert.box tanımladık

        v_box.addWidget(self.buton)
        v_box.addWidget(self.yazı_alanı)
        v_box.addStretch()


        h_box = QtWidgets.QHBoxLayout() #hbox tanımladık

        h_box.addStretch()
        h_box.addLayout(v_box)  #vboxu araya koyduk
        h_box.addStretch()


        self.setLayout(h_box) #vbox hboxun içinde olduğu için

        self.buton.clicked.connect(self.click)  # butonu fonksiyona bağladık

        self.show()

    def click(self):
        self.say += 1
        self.yazı_alanı.setText("Bana " + str(self.say) + " defa tıklandı.") # butona tıklandıkça yazının değişmesi için





app = QtWidgets.QApplication(sys.argv)

pencere = Pencere()

sys.exit(app.exec_())
