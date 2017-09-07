# Let's pretend this is a slow function
import time
def square(num):
  time.sleep(1)
  print("Calculating square at time " + str(time.time()))
  return num*num

def square_numbers(nums):
  #result = []
  #for i in nums:
  #  result.append(square(i))
  #return result
  return [square(i) for i in nums]

print("When you don't use a generator...")

# Creates list of square numbers
sqNums = square_numbers([1,2,3,4,5])

for num in sqNums:
  print(num)
# OR could just print(sqNums), as the work of squaring the numbers is already done

####
####
####
print("all squares are calculated before any of the results are returned.\n\n")

def square_numbers_gen(nums):
  for i in nums:
    yield square(i)

print("When you do use a generator...")

# Creates iterator
sqNumGen = square_numbers_gen([1,2,3,4,5])

# Each square is calculated as we go through the iterator
for num in sqNumGen:
  print(num)

####
####
####
print("each square is calculated when it is needed.\n\n")

# Boiler plate conversion of above

class square_numbers_bp:
  def __init__(self, nums):
    self.nums = nums
    self.count = 0

  def __iter__(self):
    return self

  def __next__(self):
    if self.count < len(self.nums):
      self.count += 1
      return square(self.nums[self.count-1])
    else:
      raise StopIteration()

print("When you do use a generator, boiler plate style...")

sqNumBP = square_numbers_bp([1,2,3,4,5])

for num in sqNumBP:
  print(num)