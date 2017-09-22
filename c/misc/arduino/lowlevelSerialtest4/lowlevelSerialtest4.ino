void setup() {
  pinMode(1, OUTPUT);
  digitalWrite(1, HIGH);
}

void loop() {
  // 104 microsecond delay for 9600 baud
  delayMicroseconds(104);
  digitalWrite(1, LOW); // start bit
  delayMicroseconds(104);

  // 65 = 0b01000001
  digitalWrite(1, LOW);
  delayMicroseconds(104);
  digitalWrite(1, HIGH);
  delayMicroseconds(104);
  digitalWrite(1, LOW);
  delayMicroseconds(104);
  digitalWrite(1, LOW);
  delayMicroseconds(104);
  digitalWrite(1, LOW);
  delayMicroseconds(104);
  digitalWrite(1, LOW);
  delayMicroseconds(104);
  digitalWrite(1, HIGH);
  delayMicroseconds(104);
  digitalWrite(1, LOW);
  delayMicroseconds(104);
  digitalWrite(1, HIGH); // stop bit
}
