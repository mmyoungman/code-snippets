void setup() {
  Serial.begin(19200);

  pinMode(13, OUTPUT);

  //TCCR1B |= 0b00000101;
  //TCCR1B &= 0b11111101; // clk-io/1024

  //cli();
  
  TCCR1B &= 0b11111001;
  TCCR1B |= 0b00000001; // clk-io/1

  //Serial.print("TCCR1A/B: ");
  //Serial.print(TCCR1A, BIN);
  //Serial.print(" ");
  //Serial.println(TCCR1B, BIN);

  //TCCR1A &= 0b11111100; // waveform bits to mode 0
  //TCCR1B &= 0b11100111; // the other waveform bits for mode 0
  //*(uint8_t *)0x80 &= 0b11111100;
  //*(uint8_t *)0x81 &= 0b11100111;

  TCCR1A &= 0b11111100; // mode 4 CTC (Clear Timer on Compare) TOP OCR1A
  TCCR1B &= 0b11100111;
  TCCR1B |= 0b00001000;

  TIMSK1 |= 0b00000010; // Set COMPA interrupt on
  
  OCR1A = 16000; // COMPA interrupt called every 1ms
  
  //sei();
  
  //cli(); // disable interrupts
}

void loop() {

  //Serial.println(TCNT1);
  //Serial.println(count);
  //Serial.println();

  //if(TCNT1 > 16000) {
  //  count++;
  //  TCNT1 = 0;
  //}

  //noInterrupts();
  //if(count >= 1000) {
    //Serial.println(count);
  //  digitalWrite(13, !digitalRead(13));
  //  count = 0;
  //}
  //interrupts();

  digitalWrite(13, HIGH);
  delay2(1000);
  digitalWrite(13, LOW);
  delay2(1000);
  
}

volatile uint16_t count;
//uint16_t count;

ISR(TIMER1_COMPA_vect) {
  //Serial.println(TCNT1); is always zero
  count++;
}

void delay2(uint16_t milliseconds) {
  count = 0;
  TCNT1 = 0;
  while(count < milliseconds);
}

