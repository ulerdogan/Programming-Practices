import sys
from PyQt5 import QtWidgets


def Pencere():

    app = QtWidgets.QApplication(sys.argv)

    pencere = QtWidgets.QWidget()
    pencere.setWindowTitle("PyQt5 Ders 3")

    buton = QtWidgets.QPushButton(pencere)  # buton oluşturma
    buton.setText("Bu bir butondur")  # butona yazı ekleme

    etiket = QtWidgets.QLabel(pencere)
    etiket.setText("Merhaba dünya!")

    etiket.move(200, 30)
    buton.move(190, 80)

    pencere.setGeometry(100, 100, 500, 500)

    pencere.show()

    sys.exit(app.exec_())


Pencere()
