import sys
from PyQt5 import QtWidgets,QtGui #fotoğrafı eklemek için ekledik

def Pencere():

    app = QtWidgets.QApplication(sys.argv)
    
    pencere = QtWidgets.QWidget()
    pencere.setWindowTitle("PyQt5 Ders 2")

    etiket1 = QtWidgets.QLabel(pencere) # etiket tanımlama ve pencere üzerine yapıştırma işlemi
    etiket1.setText("Burası bir yazıdır.") # etiketi yazma
    etiket1.move(150,30) # etiketi taşıma

    etiket2 = QtWidgets.QLabel(pencere)
    etiket2.setPixmap(QtGui.QPixmap("python.png")) # fotoğrafı pencereye ekle
    etiket2.move(70,100) #fotoğrafı hareket ettik

    pencere.setGeometry(100,100,500,500)  #pencere büyüklüğünü ve açılacağı konumu belirleme
    
    pencere.show()

    sys.exit(app.exec_())

Pencere()


#pencereye yazı resim ekleme