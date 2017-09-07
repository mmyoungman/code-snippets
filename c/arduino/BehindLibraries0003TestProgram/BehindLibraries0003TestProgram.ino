void setup() {
  Serial.begin(19200);
  *(uint8_t *)0x80 &= 0b11111100;
  *(uint8_t *)0x81 &= 0b11100111;

  TCCR1B |= 0b00000110;
  TCCR1B &= 0b11111110; // count on pin 5 falling?

  //Serial.print("TCCR1A: ");
  //Serial.println(TCCR1A, BIN);
  //Serial.print("TCCR1B: ");
  //Serial.println(TCCR1B, BIN);
}

void loop() {
  Serial.println(*(uint16_t *)0x84);
}


