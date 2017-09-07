def add(x,y):
  return x + y

def makeSubFunc():
  def newFunc(x,y):
    return x - y
  return newFunc

add = makeSubFunc()

print(add(2,1))