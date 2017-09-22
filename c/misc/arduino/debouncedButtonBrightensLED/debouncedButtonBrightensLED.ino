int switchPin = 8;
int ledPin = 11;
int ledLevel = 0;
boolean lastButton = LOW;
boolean currentButton = LOW;

void setup()
{
  pinMode(switchPin, INPUT);
  pinMode(ledPin, OUTPUT);
}

void loop()
{
  currentButton = digitalRead(switchPin);
  if(currentButton != lastButton)
  {
    delay(5);
    currentButton = digitalRead(switchPin);
  }
  
  if(lastButton == LOW && currentButton == HIGH)
  {
    ledLevel += 51;
  }
  
  if(ledLevel > 255)
    ledLevel = 0;
  
  lastButton = currentButton;
  analogWrite(ledPin, ledLevel);
}
