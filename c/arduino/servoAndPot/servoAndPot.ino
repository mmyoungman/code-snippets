#include <Servo.h>

Servo serv;
int servoPin = 9;

int potPin = 1;

void setup() 
{                
   serv.attach(servoPin);
   pinMode(potPin, INPUT);
   
   Serial.begin(9600);
}

void loop() 
{
  //Serial.println(analogRead(potPin));
  int value = map(analogRead(potPin), 0, 1023, 20, 180);
  serv.write(value);
}
