import sys
from PyQt5 import QtWidgets  # pencere gibi özelliklerin olduğu kısmı dahil ettik

def Pencere():

    app = QtWidgets.QApplication(sys.argv) #uygulama oluşturmak istediğimizi belirttik
    
    pencere = QtWidgets.QWidget() # pencere oluşturma isteği, ğencere objemiz oluştu

    pencere.setWindowTitle("PyQt5 Ders 1") # title koyduk

    pencere.show() # pencere göster

    sys.exit(app.exec_()) # uygulamanın sürekli çalışır durumda kalmasını sağlama

Pencere()

#pencere oluşturma