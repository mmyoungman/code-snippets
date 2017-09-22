volatile uint32_t count = 0;

void setup() {
  Serial.begin(19200);
  pinMode(13, OUTPUT);

  // Trigger interrupt on timer2 overflow
  TIMSK2 |= 0b00000001;
  // no interrupt on compare match
  //TIMSK2 &= 0b11111001;

  // Set clock select bits to clk-io/1024 (slow)
  TCCR2B |= 0b00000111;
  //Set clock to stop?
  //TCCR2B &= 0b11111000;

  //Set clock on but fast
  //TCCR2B |= 0b00000001;
  //TCCR2B &= 0b11111001;

  //Set waveform generation mode to 0
  TCCR2A &= 0b11111100;
  TCCR2B &= 0b11110111;

  //Disable compare match
  TCCR2A |= 0b00001111;
}

ISR(TIMER2_OVF_vect)
{
  count++;
}

void loop() 
{
  digitalWrite(13, HIGH);
  delay2(50);
  digitalWrite(13, LOW);
  delay2(50);
  
  //Serial.print("TCNT2: ");
  //Serial.println(TCNT2);
  //Serial.print("Count: ");
  //Serial.println(count);
  //Serial.println();
}

void delay2(uint32_t numOverflows)
{
  count = 0;
  TCNT2 = 0;
  while(count < numOverflows);
}


