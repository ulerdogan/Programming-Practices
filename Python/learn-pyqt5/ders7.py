import sys

from PyQt5.QtWidgets import QWidget,QApplication,QCheckBox,QLabel,QPushButton,QVBoxLayout #sadece bunları import ettik


class Pencere(QWidget):
    def __init__(self):

        super().__init__()

        self.init_ui()

    def init_ui(self):

        self.checkbox = QCheckBox("Python'u seviyor musunuz?")
        self.yazi_alani = QLabel("")
        self.button = QPushButton("Bana Tıkla")

        v_box = QVBoxLayout()

        v_box.addWidget(self.checkbox)
        v_box.addWidget(self.yazi_alani)
        v_box.addWidget(self.button)

        self.setLayout(v_box)

        self.setWindowTitle("Check Box")

        self.button.clicked.connect(lambda : self.click(self.checkbox.isChecked(), self.yazi_alani))

        self.show()
    
    def click(self, checkbox, yazi_alani):
        if checkbox:
            yazi_alani.setText("Python'u seviyorsun!")

        else:
            yazi_alani.setText("NEDEN?")


app = QApplication(sys.argv)

pencere = Pencere()

sys.exit(app.exec_())