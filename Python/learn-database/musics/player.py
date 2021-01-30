from liste import *


def processInputs():
    print('''
    //\\\//\\\//\\\//\\\//\\\//\\\//\\\//\\\//
    Welcome to the MUSIC PLAYER!
    ----------------------------
    Please, commit the process that you want to apply!
    1. Show the musics...
    2. Find a music...
    3. Add a new music...
    4. Delete a music from list...
    5. Get info about the playlist...
    6. Quit from the MUSIC PLAYER!
    \\\//\\\//\\\//\\\//\\\//\\\//\\\//\\\//\\\\
    ''')


processInputs()
player = Playlist()


while True:
    print("to see the process options again, please enter 'R'...")
    process = input("Chosen process: ")

    if process[0] == "1":
        player.showList()


    elif process[0] == "2":
        isim = input("Please, enter the music that you are finding: ")
        player.findMusic(isim)

    elif process[0] == "3":
        name = input("Name: ")
        singer = input("Singer: ")
        album = input("Album: ")
        length = input("Length (m:sc format): ")
        print("Music is adding to list...")
        music = Music(name, singer, album, length)
        player.addMusic(music)
        print("Song added!")

    elif process[0] == "4":
        name = input("Please, enter the music that you want to delete: ")
        player.delMusic(name)

    elif process[0] == "5":
        player.totalLength()
        player.countMusic()

    elif process[0] == "6":
        break

    elif process[0] == "R":
        processInputs()
    
    else:
        print("This is not a setted command!")