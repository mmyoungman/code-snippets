"""
class Celsius:
    def __init__(self, temperature = 0):
        self.temperature = temperature

    def to_fahrenheit(self):
        return (self.temperature * 1.8) + 32

man = Celsius()
man.temperature = 37
print(man.temperature)
print(man.to_fahrenheit())
man.temperature = -274
"""

# Have to adapt code above so people can still use temperature while enforcing rule that
# temperature must be >= -273

class Celsius:
    def __init__(self, temperature = 0):
        self.set_temperature(temperature)

    def to_fahrenheit(self):
        return (self.get_temperature() * 1.8) + 32

    # new update
    def get_temperature(self):
        return self._temperature

    def set_temperature(self, value):
        if value < -273:
            raise ValueError("Temperature below -273 is not possible")
        self._temperature = value

    temperature = property(get_temperature,set_temperature)

# Same code as above
man = Celsius()
man.temperature = 37
print(man.temperature)
print(man.to_fahrenheit())
#man.temperature = -274

# property(fget=None, fset=None, fdel=None, doc=None)

"""
class Celsius:
    def __init__(self, temperature = 0):
        self._temperature = temperature

    def to_fahrenheit(self):
        return (self.temperature * 1.8) + 32

    @property
    def temperature(self):
        return self._temperature

    @temperature.setter
    def temperature(self, value):
        if value < -273:
            raise ValueError("Temperature below -273 is not possible")
        self._temperature = value
"""