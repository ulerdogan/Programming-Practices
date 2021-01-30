import sys

from PyQt5.QtWidgets import QWidget,QApplication,QCheckBox,QLabel,QPushButton,QVBoxLayout,QLineEdit,QHBoxLayout, QVBoxLayout, QTextEdit

import gmail



class Pencere(QWidget):
    def __init__(self):
        super().__init__()

        self.init_ui()


    def init_ui(self):

        self.user = "GMAIL USERNAME"
        self.password = "GMAIL PASSWORD"

        self.sender = QLabel("Mail gönderen hesap: " + self.user +"gmail.com")

        self.to = QLineEdit()
        self.to.setPlaceholderText("Mail kime gönderilsin?")

        self.subject = QLineEdit()
        self.subject.setPlaceholderText("Mail konusu:")

        self.text = QTextEdit()
        self.text.setPlaceholderText("Mail metni")

        self.send = QPushButton("Gönder")
        self.clean = QPushButton("Temizle")

        self.final = QLabel("")


        h_box = QHBoxLayout()
        h_box.addWidget(self.send)
        h_box.addWidget(self.clean)

        v_box = QVBoxLayout()
        v_box.addWidget(self.sender)
        v_box.addStretch()
        v_box.addWidget(self.to)
        v_box.addWidget(self.subject)
        v_box.addWidget(self.text)
        v_box.addStretch()
        v_box.addWidget(self.final)
        v_box.addLayout(h_box)

        self.setLayout(v_box)

        self.send.clicked.connect(self.sending)
        self.clean.clicked.connect(self.click_clean)


        self.setWindowTitle("GMail - Mail Gönderici")
        self.show()

    def click_clean(self):
        self.to.clear()
        self.subject.clear()
        self.text.clear()
        self.final.clear()

    def sending(self):        
        try:
            gmail.send_mail(self.user, self.password, self.user+"@gmail.com", self.to.text(), self.subject.text(), self.text.toPlainText())
            self.final.setText("Mail başarıyla gönderildi.")
        except:
            self.final.setText("Bir hata oluştu. Mailiniz gönderilemedi...")


app = QApplication(sys.argv)

pencere = Pencere()

sys.exit(app.exec_())