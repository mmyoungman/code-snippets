import hashlib

chars_found = 0
index = 0
id = "reyedfim"
password = {}

while chars_found < 8:
   str_to_hash = (id + str(index)).encode('utf-8')
   hash = hashlib.md5(str_to_hash).hexdigest()
   if hash[:5] == "00000":
      if hash[5].isdigit() and int(hash[5]) < 8 and hash[5] not in password:
         password[hash[5]] = hash[6]
         chars_found += 1
   index += 1

answer = ""
for key in sorted(password):
   answer += password[key]
print(answer)