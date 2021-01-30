import smtplib

from email.mime.multipart import MIMEMultipart

from email.mime.text import MIMEText

import sys


def send_mail(un, pw, fr, to, sj, tx):

    try:
        mails = MIMEMultipart()

        mails["From"] = fr
        mails["To"] = to
        mails["Subject"] = sj

        message = MIMEText(tx, "plain")
        mails.attach(message)

        mail = smtplib.SMTP("smtp.gmail.com", 587) 
        mail.ehlo() 
        mail.starttls() 
        mail.login(un, pw) 
        mail.sendmail(mails["From"], mails["To"], mails.as_string())
        mail.close()
    
    except:
        sys.stderr.write("Mail göndermesi başarısız oldu... \n")
        sys.stderr.flush()
        sys.exit()