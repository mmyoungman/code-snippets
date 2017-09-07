import requests
from bs4 import BeautifulSoup

response = requests.get("http://feeds.bbci.co.uk/news/rss.xml?edition=uk")

# pip install lxml to parse xml
soup = BeautifulSoup(response.text, "xml")

rssItems = []

for item in soup.find_all("item"):
	thisDict = {}
	thisDict["title"] = title = item.title.text.strip()
	thisDict["date"] = item.pubDate.text.strip()
	thisDict["link"] = item.link.text.strip()
	rssItems.append( thisDict )

#print rssItems

storyItems = []

for item in rssItems:
	response = requests.get(item["link"])
	soup = BeautifulSoup(response.text, "html.parser")

	if( response.url.startswith("http://www.bbc.co.uk/news/av") or 
			response.url.startswith("http://www.bbc.co.uk/news/blog") or 
			response.url.startswith("http://www.bbc.co.uk/news/in-pictures")	):
		continue

	elif( response.url.startswith("http://www.bbc.co.uk/news/") ):
		thisDict = {}
		thisDict["headline"] = soup.find("div", attrs={"class": "story-body"}).find("h1").text.strip()

		storyBody = soup.find("div", attrs={"property": "articleBody"}).find_all("p")
		for i, p in enumerate(storyBody):
			storyBody[i] = p.text.strip()
		thisDict["storyBody"] = storyBody

		storyItems.append(thisDict)

print storyItems
		