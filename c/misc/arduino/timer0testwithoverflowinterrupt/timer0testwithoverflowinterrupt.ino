volatile int LEDOn = 0;

void setup() {
  // Use PORTD to set pins 5,6 to be off
  *(uint8_t *)0x2B &= 0b10011111;

  // DDRD set pin 5,6, to be OUTPUT
  *(uint8_t *)0x2A |= 0b01100000;

  // Set TCNT0 = 0
  *(uint8_t *)0x46 = 0;

  // OCR0A = 200
  *(uint8_t *)0x47 = 0;
  // OCR0B = 100
  *(uint8_t *)0x48 = 0;

  // Set TCCR0A COM0A0 1, COM0B0 1, WGM01 1
  *(uint8_t *)0x44 |= 0b01010010;

  *(uint8_t *)0x45 |= 0b00000010;
}

void loop() {
  digitalWrite(13, LEDOn);
}

extern "C" void __vector_16(void) __attribute__ ((signal,used,externally_visable));
void __vector_16(void) {
  LEDOn = !LEDOn;
}

/*
ISR(TIMER0_OVF_vect)
{
  LEDOn = !LEDOn;
}
*/

