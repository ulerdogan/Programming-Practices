largest = None
smallest = None
firstCheck = False

while True:
    try:
        num = input("Enter a number: ")
        if num == "done":
            break

        if firstCheck == False:
            largest = num
            smallest = num
            firstCheck = True

        if float(num) >= float(largest):
            largest = num
        elif float(num) <= float(smallest):
            smallest = num

    except:
        print("Invalid input")

print("Maximum is", largest)
print("Minimum is", smallest)
