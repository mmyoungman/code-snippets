stations = { "Holborn" : [1], 
					   "Earl's Court" : [1, 2], 
						 "Wimbledon" : [3], 
						 "Hammersmith" : [2], 
						 "Royal Albert" : [3] }

journeys = [ [1, 1, 2.50],
						 [2, 2, 2.00],
						 [3, 3, 2.00],
						 [1, 2, 3.00],
						 [1, 3, 3.20],
						 [2, 1, 3.00],
						 [2, 3, 2.25],
						 [3, 1, 3.20],
						 [3, 2, 2.25] ]

maxCharge = 3.20
busCharge = 1.80

class oysterCard(object):
	def __init__(self, balance):
		self.balance = balance
		print "New Oyster card with", self.balance, "pounds"

	def travelBus(self):
		self.balance -= busCharge

	def enterTube(self, location):
		self.balance -= maxCharge
		self.location = location

	def leaveTube(self, destination):
		if hasattr(self, "location"):
			possibleCharges = []
			for x in range( len(stations[self.location]) ):
				for y in range ( len(stations[destination]) ):
					for journey in journeys:
						if journey[0:2] == [ stations[self.location][x], stations[destination][y] ]:
							possibleCharges.append( journey[2] )
							break
			self.balance += maxCharge - min(possibleCharges)
			del self.location
		else:
			print "Cannot leave tube since you never entered!"

newCard = oysterCard(30.00)

newCard.enterTube("Holborn")
newCard.leaveTube("Earl's Court")

newCard.travelBus()

newCard.enterTube("Hammersmith")
newCard.leaveTube("Royal Albert")

newCard.enterTube("Royal Albert")
newCard.leaveTube("Holborn")

newCard.enterTube("Holborn")

print "Oyster card final balance:", newCard.balance