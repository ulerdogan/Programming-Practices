# 9.4 Write a program to read through the mbox-short.txt and figure out who has sent the greatest number of mail messages. The program looks for 'From ' lines and takes the second word of those lines as the person who sent the mail. The program creates a Python dictionary that maps the sender's mail address to a count of the number of times they appear in the file. After the dictionary is produced, the program reads through the dictionary using a maximum loop to find the most prolific committer.

name = input("Enter file:")
if len(name) < 1:
    name = "mbox-short.txt"
handle = open(name)

senders_list = dict()

for line in handle:
    if line.startswith("From "):
        sender_line = line.split()
        sender = sender_line[1]
        if sender in senders_list:
            senders_list[sender] = senders_list[sender] + 1
        else:
            senders_list[sender] = 1
    else:
        continue

max = 0
max_sender = ""

for sender, count in senders_list.items():
    if count >= max:
        max = count
        max_sender = sender
    else:
        continue

print(max_sender, max)
