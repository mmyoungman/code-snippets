import hashlib

chars_found = 0
password = ""
index = 0
id = "reyedfim"

while chars_found < 8:
   str_to_hash = (id + str(index)).encode('utf-8')
   hash = hashlib.md5(str_to_hash).hexdigest()
   if hash[:5] == "00000":
      chars_found += 1
      password += hash[5]
   index += 1

print(password)