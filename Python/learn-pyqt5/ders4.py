import sys
from PyQt5 import QtWidgets


def Pencere():

    app = QtWidgets.QApplication(sys.argv)

    # iki buton oluşturduk ama henüz yerleştirmedik
    okay = QtWidgets.QPushButton("Tamam")
    cancel = QtWidgets.QPushButton("İptal")

    # h_box = QtWidgets.QHBoxLayout() # yatay bir layout oluşturdukv
    # h_box.addWidget(okay) #butonları horizontal layouta yerleştirdik
    # h_box.addWidget(cancel)
    # h_box.addStretch() # boşluk eklemeyi sağlar. butonlardan önce sonra veya arasında tanımlanabilir

    # # v_box = QtWidgets.QVBoxLayout() # dikey bir layout oluşturduk
    # # v_box.addWidget(okay) #butonları verticaş layouta yerleştirdik
    # # v_box.addWidget(cancel)
    # # v_box.addStretch() # boşluk eklemeyi sağlar. butonlardan önce sonra veya arasında tanımlanabilir

    h_box = QtWidgets.QHBoxLayout() # butonların hep sağ altta kalması için önce horizontal layoutumu oluşturdum

    h_box.addStretch() #sonra butonlarımı h. ly. yerleştirdim
    h_box.addWidget(okay)
    h_box.addWidget(cancel)

    v_box = QtWidgets.QVBoxLayout() #sonra ver. layout oluşturdum

    v_box.addStretch() #önce boşluk koyduk aşağı itmesi için
    v_box.addLayout(h_box) #horiz. ly. a v.ly i ekledim

    pencere = QtWidgets.QWidget()
    pencere.setWindowTitle("PyQt5 Ders 4")

    # pencere.setLayout(h_box) # horiz. layoutu yerleştirdik
    # # pencere.setLayout(v_box) # vert. layoutu yerleştirdik

    pencere.setLayout(v_box) #en sonda da vly yerleştirdik

    pencere.setGeometry(100, 100, 500, 500)

    pencere.show()

    sys.exit(app.exec_())


Pencere()

# horizontal and ## vertical box layout
