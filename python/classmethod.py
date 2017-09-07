class Circle(object):
    PI = 3.142
    def __init__(self, radius):
        self.radius = radius

    def area(self):
        return self.PI * self.radius ** 2.0

    def circumference(self):
        return 2 * self.radius * self.PI

    # Alternate constructor/init
    @classmethod
    def from_bbd(cls, bbd):
      'Construct circle from a bounding box diagonal'
      radius = bbd / 2.0 / math.sqrt(2.0)
      #return Circle(radius)
      return cls(radius)

myCir = Circle(5)
myCir2 = Circle.from_bbd(20)