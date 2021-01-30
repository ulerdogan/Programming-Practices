import sqlite3
import time
import re


class Music():
    def __init__(self, name, singer, album, length):
        time.sleep(1)
        self.name = name
        self.singer = singer
        self.album = album
        self.length = length
        self.like = False

    def __str__(self):
        return "Song Name: {}, Singer: {}, Album: {}, Length: {}".format(self.name, self.singer, self.album, self.length)

    def likeMusic(self):
        self.like = True


class Playlist():
    def __init__(self):
        self.connect()

    def connect(self):
        self.connection = sqlite3.connect("musics.db")
        self.cursor = self.connection.cursor()
        ask = "Create Table If not exists musics (name TEXT,singer TEXT,album TEXT,length TEXT)"
        self.cursor.execute(ask)
        self.connection.commit()

    def disconnect(self):
        self.connection.close()

    def showList(self):
        ask = "Select * From musics"
        self.cursor.execute(ask)
        musics = self.cursor.fetchall()

        if len(musics) > 0:
            for i in musics:
                music = Music(i[0], i[1], i[2], i[3])
                print(music)
        else:
            time.sleep(1)
            print("Musics couldn't find...")

    def findMusic(self, name):
        ask = "Select * From musics"
        self.cursor.execute(ask)
        musics = self.cursor.fetchall()
        check = True
        for i in musics:
            if re.findall(name.lower(), i[0][0].lower()):
                print(i[0])
                check = False
        if check:
            time.sleep(1)
            print("This song couldn't found...")

    def addMusic(self, music):
        try:
            ask = "Insert into musics Values(?,?,?,?)"
            self.cursor.execute(
                ask, (music.name, music.singer, music.album, music.length))
            self.connection.commit()
            time.sleep(1)
        except:
            time.sleep(1)
            print("An error has occured!")

    def delMusic(self, music):
        try:
            ask = "Delete From musics where name = ?"
            self.cursor.execute(ask, (music,))
            self.connection.commit()
            print("The music deleted succesfuly...")
        except:
            time.sleep(1)
            print("An error has occured!")

    def totalLength(self):
        ask = "Select length From musics"
        self.cursor.execute(ask)
        lengths = self.cursor.fetchall()
        minutes = 0
        seconds = 0
        if len(lengths) > 0:
            for i in lengths:
                minutes += int(i[0].split(":")[0])
                seconds += int(i[0].split(":")[1])

                minutes += int((int(seconds)/60))
                seconds = seconds % 60
        print("The total length of playlist is {} minutes and {} seconds".format(
            minutes, seconds))

    def countMusic(self):
        ask = "Select * From musics"
        self.cursor.execute(ask)
        musics = self.cursor.fetchall()
        counter = 0
        for i in musics:
            counter += 1
        print("There are {} music".format(counter))
