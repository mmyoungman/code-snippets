class Person(obj):
  """ PERSON DOCUMENTATION """
  def __init__(self, name, age):
    self.name = name
    self.age = age

  def __get__(self, obj, cls=None):
    return [self.name, self.age]

  def __set__(self, obj, details):
    self.name = details["name"]
    self.age = details["age"]

class People(obj):
  mark = Person("Mark", 34)

people = People()

details = people.mark #runs __get__
print("Details: ", details)

name, age = mark.descriptor # runs __get__
print("Name:", name)
print("Age:", age)

people.mark = { "name": "Timmy", "age": 52 } # runs __set__
print("Name:", people.mark[0]) # runs __get__
print("Age:", people.mark[1]) # runs __get__ again