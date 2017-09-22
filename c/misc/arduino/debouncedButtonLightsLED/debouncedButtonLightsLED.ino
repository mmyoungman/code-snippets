int switchPin = 8;
int ledPin = 13;
boolean ledOn = false;
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
    ledOn = !ledOn;
  }
  
  lastButton = currentButton;
  digitalWrite(ledPin, ledOn);
}
