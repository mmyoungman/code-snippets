class Person(object):
  """ PERSON CLASS DOCUMENTATION """
  def __init__(self, name, age):
    self.name = name
    self.age = age

  def ageIncrement():
    self.age += 1

mark = Person("Mark", 33)
print("Their name is " + mark.name + " and they're " + str(mark.age) + " years old\n")

# dir is a built-in function that lists all of an object's attributes
print("dir(mark):")
print(dir(mark))
print("\n")

# __dict__ contains a list of variables and their values
print("mark.__dict__:")
print(mark.__dict__)
print("\n")



print("mark.__class__.__mro__:")
print(mark.__class__.__mro__)
print("\n")

print("mark.__getattribute__:")
print(mark.__getattribute__)
print("\n")

print("Person.__bases__:")
print(Person.__bases__)
print("\n")

# __class__ returns the class, so mark is of class Person, and Person is of class type
print("mark.__class__:")
print(mark.__class__)
print("\n")

print("Person.__class__:")
print(Person.__class__)
print("\n")

# __doc__ returns the doc string
print("mark.__doc__:")
print(mark.__doc__)
print("\n")

print("Person.__doc__:")
print(Person.__doc__)
print("\n")