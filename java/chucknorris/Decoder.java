package chucknorris;
import java.util.Scanner;

public class Decoder {

    private String getInput() {
        Scanner scanner = new Scanner(System.in);
        System.out.println("Input encoded string:");
        String input = scanner.nextLine();
        return input;
    }

    private boolean isValid(String input) {

        // The encoded message includes characters other than 0 or spaces;
        for (int i = 0; i < input.length(); i++) {
            if (input.charAt(i) != '0' && input.charAt(i) != ' ') {
                return false;
            }
        }

        String[] brokenInput = input.split(" ");
        // The number of blocks is odd;
        if (brokenInput.length % 2 == 1) {
            return false;
        }
        // The first block of each sequence is not 0 or 00;
        for (int i = 0; i < brokenInput.length - 1; i += 2) {
            String toBeValidated = brokenInput[i];
            if (!toBeValidated.equals("0") && !toBeValidated.equals("00")) {
                return false;
            }
        }
        // The length of the decoded binary string is not a multiple of 7.
        if (toBinaryString(input).length() % 7 != 0) {
            return false;
        }
        return true;
    }

    private static String toBinaryString(String validInput) {
        String[] brokenInput = validInput.split(" ");
        StringBuilder builder = new StringBuilder();
        for (int i = 0; i < brokenInput.length; i += 2) {
            char c = brokenInput[i].equals("0")? '1': '0';
            for (int j = 0; j < brokenInput[i+1].length(); j++) {
                builder.append(c);
            }
        }
        return builder.toString();
    }

    private static String toEnglish(String binaryString) {
        StringBuilder builder = new StringBuilder();
        char c;
        for (int i = 0; i < binaryString.length() - 6; i += 7) {
            String sevenDigits = binaryString.substring(i, i+7);
            c = (char) Integer.parseInt(sevenDigits, 2);
            builder.append(c);
        }
        return builder.toString();
    }

    private String translateInput(String validInput) {
        String binaryString = toBinaryString(validInput);
        return toEnglish(binaryString);
    }

    public void start() {
        String input = getInput();
        if (!isValid(input)) {
            System.out.println("Encoded string is not valid.");
            System.out.println();
            return;
        }
        System.out.println("Decoded string:");
        System.out.println(translateInput(input));
        System.out.println();
        return;
    }
}
