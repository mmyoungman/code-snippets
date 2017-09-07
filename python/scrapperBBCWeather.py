import urllib2
from bs4 import BeautifulSoup

webpage = urllib2.urlopen("http://www.bbc.co.uk/weather/s17")

soup = BeautifulSoup(webpage, "html.parser")

tempList = []
postcode = "s17"

for x in range(5):
	day = "/weather/"+postcode+"?day=" + str(x)
	thisDay = ( soup.find("ul", attrs={"class": "daily"})
									.find("a", attrs={"href": day})
									.find("span", attrs={"class": "day-name"}) )
	thisTemp = ( soup.find("ul", attrs={"class": "daily"})
									.find("a", attrs={"href": day})
									.find("span", attrs={"class": "max-temp"})
									.find("span", attrs={"class": "temperature-value"}) )
	thisList = [ thisDay.text.strip(), ''.join(i for i in thisTemp.text.strip() if i.isdigit()) ]
	tempList.append(thisList)

print tempList