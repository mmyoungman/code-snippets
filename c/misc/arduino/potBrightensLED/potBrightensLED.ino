int potPin = 0;
int ledPin = 11;
int ledLevel = 0;

void setup()
{
  pinMode(potPin, INPUT);
  pinMode(ledPin, OUTPUT);
  
  Serial.begin(9600);
}

void loop()
{
  //Serial.println(analogRead(potPin));
  ledLevel = (analogRead(potPin) / 1024.0) * 255.0;
  //Serial.println(ledLevel);
  analogWrite(ledPin, ledLevel);
  delay(100);
}
