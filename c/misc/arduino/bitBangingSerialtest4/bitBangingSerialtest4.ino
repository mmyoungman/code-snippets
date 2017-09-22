void setup() {
  pinMode(1, OUTPUT);
  digitalWrite(1, HIGH);
}

void loop() {
  // 104 microsecond delay for 9600 baud
  delayMicroseconds(104);
  digitalWrite(1, LOW); // start bit
  
  // 'A' == 65 == 0b01000001
  // Least signficant bit first
  delayMicroseconds(104);
  digitalWrite(1, HIGH); // 1
  delayMicroseconds(104);
  digitalWrite(1, LOW);  // 0
  delayMicroseconds(104);
  digitalWrite(1, LOW);  // 0
  delayMicroseconds(104);
  digitalWrite(1, LOW);  // 0
  delayMicroseconds(104);
  digitalWrite(1, LOW);  // 0
  delayMicroseconds(104);
  digitalWrite(1, LOW);  // 0
  delayMicroseconds(104);
  digitalWrite(1, HIGH); // 1
  delayMicroseconds(104);
  digitalWrite(1, LOW);  // 0
  
  delayMicroseconds(104);
  digitalWrite(1, HIGH); // stop bit
}
