void setup() {
  Serial.begin(19200);

  pinMode(13, OUTPUT);

  TCCR1B |= 0b00000101;
  TCCR1B &= 0b11111101; // clk-io/1024

  //TCCR1B &= 0b11111001;
  //TCCR1B |= 0b00000001; // clk-io/1

  //Serial.print("TCCR1A/B: ");
  //Serial.print(TCCR1A, BIN);
  //Serial.print(" ");
  //Serial.println(TCCR1B, BIN);

  // Waveform generation mode 0
  *(uint8_t *)0x80 &= 0b11111100;
  *(uint8_t *)0x81 &= 0b11100111;

  //TCCR1A &= 0b11111100; // waveform bits to mode 0
  //TCCR1B &= 0b11100111; // the other waveform bits for mode 0
  
}

int count;

void loop() {
  //Serial.println(TCNT1);
  //Serial.println(*(volatile uint16_t *)0x84);

  if(TCNT1 > 16000) {
    //digitalWrite(13, !digitalRead(13));
    count++;
    TCNT1 = 0;
    
    if(count > 10) {
      digitalWrite(13, !digitalRead(13));
      count = 0;
    }
  }
    //Serial.println(TCNT1);


}
