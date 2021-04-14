char_array = "xyz0123456789"

for letter in char_array[3:6]:
    print(letter)
print("--------------------------")

print(ord("a"), ord("b"))
print(chr(97), chr(99))
print("--------------------------")

for i in range(ord("a"), ord("z")):
    print(chr(i))