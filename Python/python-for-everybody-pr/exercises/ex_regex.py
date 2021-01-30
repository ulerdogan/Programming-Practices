import re

no = input("Please enter the text number that you want to open: regex_sum_???.txt --> fill the ??? : ")

# file option 1: regex_sum_42.txt
# file option 1: regex_sum_862206.txt

try:
    file = open("regex_sum_" + no + ".txt")
    print("regex_sum_" + no + ".txt has opened.")
except:
    print("File couldn't find!")

numbers = list()
print(re.findall("[0-9]+", file.read()))
for line in file:
    numbers += re.findall("[0-9]+", line)

sum = 0

for number in numbers:
    sum += int(number)

print(sum)
