void setup() 
{
  Serial.begin(19200);
                  
  *(uint8_t *)0x24 &= 0b11101111; // set pin 12 to INPUT
  *(uint8_t *)0x24 |= 0b00100000; // set pin 13 to OUTPUT
}

void loop() 
{
  Serial.println((*(uint8_t *)0x23) & 0b00010000);

  if( (*(uint8_t *)0x23) & 0b00010000) // if PINB 5th bit isn't 0
    digitalWrite(13, HIGH);
  else
    digitalWrite(13, LOW);
    
  delay(10);
  
}



