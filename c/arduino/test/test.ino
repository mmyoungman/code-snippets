  uint8_t ucsr0a;
  uint8_t ucsr0b;
  uint8_t ucsr0c;
  uint8_t ubrr0l;
  uint8_t ubrr0h;

void setup() {
  ucsr0a = UCSR0A;
  ucsr0b = UCSR0B;
  ucsr0c = UCSR0C;
  ubrr0l = UBRR0L;
  ubrr0h = UBRR0H;

  Serial.begin(9600);
}

void loop() {
  Serial.print("UCSR0A: ");
  Serial.println(UCSR0A, BIN);
  Serial.print("UCSR0B: ");
  Serial.println(UCSR0B, BIN);
  Serial.print("UCSR0C: ");
  Serial.println(UCSR0C, BIN);
  Serial.print("UBRR0L: ");
  Serial.println(UBRR0L, BIN);
  Serial.print("UBRR0H: ");
  Serial.println(UBRR0H, BIN);
  Serial.print("UBRR0: ");
  Serial.println(UBRR0, BIN);
  Serial.println();

  Serial.println("Before Serial.begin was called...");
  Serial.print("UCSR0A: ");
  Serial.println(ucsr0a, BIN);
  Serial.print("UCSR0B: ");
  Serial.println(ucsr0b, BIN);
  Serial.print("UCSR0C: ");
  Serial.println(ucsr0c, BIN);
  Serial.print("UBRR0L: ");
  Serial.println(ubrr0l, BIN);
  Serial.print("UBRR0H: ");
  Serial.println(ubrr0h, BIN);
  Serial.println();
}

//ISR(USART_UDRE_vect) {}
//ISR(USART_RX_vect) {}
//ISR(USART_TX_vect) {}

