import csv

with open('NAVData20160330.csv', 'r') as f:
  reader = csv.reader(f)
  your_list = list(reader)

print(your_list)
