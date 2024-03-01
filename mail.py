#!/usr/bin/python
# -*- coding: UTF-8 -*-
#
import smtplib, email
from email import encoders
from os.path import basename
from email.mime.base import MIMEBase
from email.mime.application import MIMEApplication
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
from email.header import Header

port = 587
smtp_server = "smtp.gmail.com"
password = "tiyzqqlmeejmprck"

sender_email = "composition0936@gmail.com"
#receiver_email = "yatzliu@igs.com.tw,nukejhuang@igs.com.tw,jayanchen@igs.com.tw,chengkangshih@igs.com.tw"
#receiver_emails = ["yatzliu@igs.com.tw","nukejhuang@igs.com.tw","jayanchen@igs.com.tw","chengkangshih@igs.com.tw"]
#receiver_email = "nukejhuang@igs.com.tw,rainliao@igs.com.tw"
#receiver_emails = ["nukejhuang@igs.com.tw","rainliao@igs.com.tw"]
receiver_email = "nukejhuang@igs.com.tw"
receiver_emails = ["nukejhuang@igs.com.tw"]
#receiver_email = "luhongyu@igs.com.tw,chuntingchou@igs.com.tw,peterweng@igs.com.tw,nukejhuang@igs.com.tw,kailihuang@igs.com.tw,muchenghuang@igs.com.tw,rainliao@igs.com.tw,shanesu@igs.com.tw"
#receiver_emails = ["luhongyu@igs.com.tw","chuntingchou@igs.com.tw","peterweng@igs.com.tw","nukejhuang@igs.com.tw","kailihuang@igs.com.tw","muchenghuang@igs.com.tw","rainliao@igs.com.tw","shanesu@igs.com.tw"]

def sendmail(subject,text,files=None):
    message = MIMEMultipart('alternative')
    message.set_charset('utf8')
    message["From"] = sender_email
    message["To"] = receiver_email
    message["Subject"] = Header(subject,'utf-8')
    message["Bcc"] = receiver_email

    part1 = MIMEText(text, "html", "UTF-8")
    message.attach(part1)

    #for f in files or []:
    #    with open(f,"rb") as fil:
    #        part = MIMEApplication(fil.read(),Name=basename(f))
    #        part['Content-Disposition']='attachment; filename="%s"' % basename(f)
    #        message.attach(part)

    for path in files or []:
        part = MIMEBase('application', "octet-stream")
        with open(path, 'rb') as file:
            part.set_payload(file.read())
        encoders.encode_base64(part)
        part.add_header('Content-Disposition','attachment; filename="%s"' % basename(path))
        message.attach(part)

    smtp = smtplib.SMTP(smtp_server, port)
    smtp.ehlo()
    smtp.starttls()
    smtp.login(sender_email,password)
    status=smtp.sendmail(sender_email,receiver_emails,message.as_string())
    smtp.quit()
    if status == {}:
        return 1
    return 0
