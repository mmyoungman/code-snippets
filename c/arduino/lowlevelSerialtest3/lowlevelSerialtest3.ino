void setup() {
  // Baudrate to 9600
  UCSR0A &= 0b11111101; // U2X0 = 0
  UBRR0 = 103;

  //Baudrate to 250k
  //UCSR0A &= 0b11111101; // U2X0 = 0
  //UBRR0 = 3; 

  // 8 data bits
  UCSR0B &= 0b11111011; // UCSZ2 = 0
  UCSR0C |= 0b00000110; // UCSZ1:0 = 0b11

  // Enable transmission
  UCSR0B |= 0b00001000; // TXEN0 = 1
}

uint8_t transmitByte = 0;

void loop() {
  while((UCSR0A & 0b00100000) == 0); // Wait until UDRE0 bit is 1, i.e. USART ready for new data
  UDR0 = transmitByte;
  //while((UCSR0A & 0b01000000) == 0); // Wait until TXC0 bit is 1, i.e. transmission complete

  //if(transmitByte > 127)
  //  transmitByte = 0;
  //else
    transmitByte++;
}
