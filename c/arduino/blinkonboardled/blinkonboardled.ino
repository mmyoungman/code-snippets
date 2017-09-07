void setup() 
{
  Serial.begin(19200);
                  
  *(uint8_t *)0x24 &= 0b11101111; // set pin 12 to INPUT
  *(uint8_t *)0x24 |= 0b00100000; // set pin 13 to OUTPUT
 
}

void loop() 
{
  //digitalWrite(13, HIGH);   // turn the LED on (HIGH is the voltage level)
  //*(uint8_t *)0x25 |= 0b00100000; // 0x25 is the address of PORTB register, set bit on for pin 13
  //PORTB |= (1 << PORTB5);

  //delay(1000);               // wait for a 100ms
  
  //digitalWrite(13, LOW);    // turn the LED off by making the voltage LOW
  //*(uint8_t *)0x25 &= 0b11011111; // set bit off for pin 13

  //delay(1000);

  Serial.println((*(uint8_t *)0x23) & 0b00010000);

  if( (*(uint8_t *)0x23) & 0b00010000) // if PINB 5th bit isn't 0
    digitalWrite(13, HIGH);
  else
    digitalWrite(13, LOW);
    
  delay(10);
  
}



