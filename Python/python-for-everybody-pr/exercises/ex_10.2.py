# 10.2 Write a program to read through the mbox-short.txt and figure out the distribution by hour of the day for each of the messages. You can pull the hour out from the 'From ' line by finding the time and then splitting the string a second time using a colon.
# From stephen.marquard@uct.ac.za Sat Jan  5 09:14:16 2008
# Once you have accumulated the counts for each hour, print out the counts, sorted by hour as shown below.


name = input("Enter file:")
if len(name) < 1 : name = "mbox-short.txt"
handle = open(name)

hours = dict()

for line in handle:
    if not line.startswith("From "): continue

    from_line = line.split()
    from_hour = from_line[5].split(":")[0]

    if from_hour in hours:
        hours[from_hour] += 1
    else:
        hours[from_hour] = 1

for x,y in sorted(hours.items()):
    print(x,y)