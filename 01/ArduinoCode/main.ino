String command;
 
void setup() {
    Serial.begin(9600); 
}

void loop() {
    if(Serial.available()){
        command = Serial.readStringUntil('\n');
 
        Serial.println(command);
        if (command.equals("led_1_open")){
          analogWrite(5,255);
        }else if(command.equals("led_1_close")){
          analogWrite(5,0);
        }else if (command.equals("led_3_open")){
          analogWrite(7,255);
        }else if(command.equals("led_3_close")){
          analogWrite(7,0);
        }
    }
}