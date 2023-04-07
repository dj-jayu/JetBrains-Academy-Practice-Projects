package chucknorris;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        ValidOptionGetter validOptionGetter = new ValidOptionGetter();
        Encoder encoder = new Encoder();
        Decoder decoder = new Decoder();
        while (true) {
            String answer = validOptionGetter.getValidOption();
            if (answer.equals("decode")) {
                decoder.start();
            } else if (answer.equals("encode")) {
                encoder.start();
            } else {
                break;
            }
        }
        System.out.println("Bye!");
    }
}