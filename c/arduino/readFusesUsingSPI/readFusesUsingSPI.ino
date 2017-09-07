// Author: Mark Youngman
// Sketch to read lock bits and fuses of another arduino uno using SPI

#include <SPI.h>

// The below wiring is required
// Arduino with this sketch -> Programmed Arduino
// 5v -> 5v
// GND -> GND
// pin10 -> RESET
// pin11 -> pin11
// pin12 -> pin12
// pin13 -> pin13

//pin10 is used as slave select
int ssPin = 10;

void setup() {
  Serial.begin(19200);
  Serial.println("Starting SPI...");
  
  pinMode(ssPin, OUTPUT);
  //digitalWrite(ssPin, HIGH); // Necessary?
  //delay(10);
  
  SPI.begin();

  Serial.println("Waiting before sending Programming Enable...");
  digitalWrite(ssPin, LOW);
  delay(20); // datasheet says 328p needs at least 20ms before sending Programming Enable instruction

  Serial.println("Sending Programming Enable...");
  SPI.transfer(0xAC);
  SPI.transfer(0x53);
  uint8_t echo = SPI.transfer(0x00);
  SPI.transfer(0x00);

  if(echo == 0x53)
    Serial.println("Programming Enable Success!");
  else
    Serial.println("Programming Enable Fail!");

  Serial.println("Requesting Lock bits...");
  SPI.transfer(0x58);
  SPI.transfer(0x00);
  SPI.transfer(0x00);
  uint8_t lock = SPI.transfer(0x00);

  Serial.println("Requesting Fuse bits...");
  SPI.transfer(0x50);
  SPI.transfer(0x00);
  SPI.transfer(0x00);
  uint8_t fuse = SPI.transfer(0x00);

  Serial.println("Requesting Fuse High bits...");
  SPI.transfer(0x58);
  SPI.transfer(0x08);
  SPI.transfer(0x00);
  uint8_t fuseHigh = SPI.transfer(0x00);

  Serial.println("Requesting Extended Fuse bits...");
  SPI.transfer(0x50);
  SPI.transfer(0x08);
  SPI.transfer(0x00);
  uint8_t fuseExt = SPI.transfer(0x00);

  Serial.print("Lock bits: ");
  Serial.println(lock, BIN);

  Serial.print("Fuse bits: ");
  Serial.println(fuse, BIN);

  Serial.print("Fuse High bits: ");
  Serial.println(fuseHigh, BIN);

  Serial.print("Extended Fuse bits: ");
  Serial.println(fuseExt, BIN);

  digitalWrite(ssPin, HIGH);
  SPI.end();

  Serial.println("SPI ended");

}

void loop() {
}
