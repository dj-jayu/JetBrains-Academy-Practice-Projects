package chucknorris;

import java.util.Scanner;

public class ValidOptionGetter {
    public String getValidOption() {
        Scanner scanner = new Scanner(System.in);
        String input = "";

        while (true) {
            System.out.println("Please input operation (encode/decode/exit):");
            input = scanner.nextLine();
            if (input.equals("decode") || input.equals("encode") || input.equals("exit")) {
                return input;
            }
            System.out.format("There is no '%s' operation\n\n", input);
        }
    }
}
