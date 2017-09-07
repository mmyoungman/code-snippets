def makeSubFunc(func):
  # we could use the add function here, as it is passed as func
  def newFunc(x,y):
    return x - y
  return newFunc

def add(x,y):
  return x + y

add = makeSubFunc(add)

print(add(2,1))