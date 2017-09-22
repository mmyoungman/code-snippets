#include <CapacitiveSensor.h>

CapacitiveSensor cs_2_4 = CapacitiveSensor(2,4); // 10M resistor between pins 2 & 4, pin 4 is sensor pin, add a wire and or foil if desired
int ledOn = 0;
int justTurnedOn = 0;

void setup()                    
{
   Serial.begin(9600);
}

void loop()                    
{
    long start = millis();
    long total1 =  cs_2_4.capacitiveSensor(60);


    Serial.print(millis() - start);        // check on performance in milliseconds
    Serial.print("\t");                    // tab character for debug windown spacing

    Serial.print(total1);                  // print sensor output 1
    Serial.println("\t");

    if(total1 > 500 && justTurnedOn == 0)
    {
      ledOn = !ledOn;
      justTurnedOn = 1;
      digitalWrite(13, ledOn);
    }

    if(total1 < 500 && justTurnedOn == 1)
       justTurnedOn = 0;

    delay(100);                             // arbitrary delay to limit data to serial port 
}
