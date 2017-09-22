int sensePin = 0;
int ledPin = 11;
int ledLevel = 0;

void setup()
{
  pinMode(sensePin, INPUT);
  analogReference(DEFAULT); // this line isn't necessary
  pinMode(ledPin, OUTPUT);
  
  Serial.begin(9600);
}

void loop()
{
  //Serial.println(analogRead(sensePin));
  //ledLevel = 254 - ((analogRead(sensePin) / 1024.0) * 255.0);
  int val = constrain(analogRead(sensePin), 450, 650); // limits val between 450-650
  ledLevel = map(val, 450, 650, 255, 0); // changes values 450-650 to 255-0
  //Serial.println(ledLevel);
  if( ledLevel >= 0 && ledLevel <= 255) // just to make sure
    analogWrite(ledPin, ledLevel);
  delay(100);
}
