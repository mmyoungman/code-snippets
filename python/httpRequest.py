import requests

url = "http://localhost:5000/api/somethingorother"
headers = {}
data = ""
cookies = {}

r = requests.get(url, headers=headers, data=data, cookies=cookies)

print(r.url)
print(str(r.status_code))
print(str(r.headers))
print(str(r.headers['Content-Type']))
print(str(r.cookies))
print(r.text)  # response body