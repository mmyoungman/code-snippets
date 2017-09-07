class Circle(object):
    PI = 3.142
    def __init__(self, radius):
        self.radius = radius

    def area(self):
        return self.PI * self.radius ** 2.0

    @property
    def circumference(self):
        return 2 * self.radius * self.PI


myCircle = Circle(2)
myCircle.radius = 3
print(myCircle.area())
print(myCircle.circumference)