# 6.5 Write code using find() and string slicing (see section 6.10) to extract the number at the end of the line below. Convert the extracted value to a floating point number and print it out.

text = "X-DSPAM-Confidence:    0.8475"
ws = 0
output = ""

while True:
    ws = text.find(" ", ws+1)

    if ws == -1:
        break
    else:
        output = text[ws:]
print(float(output))
