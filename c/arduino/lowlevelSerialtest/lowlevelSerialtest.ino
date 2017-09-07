void setup() {
  Serial.begin(19200);

  TCCR1A &= 0b11111100;
  TCCR1B &= 0b11100111;

  // 0 prescaler
  //TCCR1B &= 0b11111001;
  //TCCR1B |= 0b00000001;

  //1024 prescaler
  TCCR1B &= 0b11111101;
  TCCR1B |= 0b00000101;
}

void loop() {
  Serial.println(TCNT1);
}
