class Circle(object):
    PI = 3.142
    def __init__(self, radius):
        self.radius = radius

    def area(self):
        return self.PI * self.radius ** 2.0

    def circumference(self):
        return 2 * self.radius * self.PI

    # Attach a function to the class -- no need for self
    @staticmethod
    def angle_to_grade(angle):
      return math.tan(math.radians(angle)) * 100.0

myCir = Circle(5)