import json
#from pprint import pprint

with open('batchSiteData.json') as data_file:    
    data = json.load(data_file)

print( json.dumps(data["items"], sort_keys=True, indent=4) )

#for i in range( len(data["items"]) ):
#	print(data["items"][i])

#print(json.dumps(data, sort_keys=True, indent=4))

#new_list = []
#for i in range( len(data["items"]) ):
#    #new_list.append( data["items"][i]["machineNo"] ) #u appears next to each element because unicode
#    print( data["items"][i]["machineNo"] ) # This prints neatly

#for i in range( len(new_list) ):
#	new_list[i].strip('"')

#print(new_list)

#new_list = []
#for i in machineData:
#	new_list.append( i.values() )
#
#print(new_list)

# Grab first 100 entries:
#print("[")
#for i in range(100):
#	print(json.dumps(data[i], sort_keys=True, indent=4))
#	if(i != 99):
#		print(",")
#print("]")


