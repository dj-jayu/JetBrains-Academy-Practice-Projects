package chucknorris;
import jdk.jshell.spi.ExecutionControl;

import java.util.Scanner;

public class Encoder {
    private static String toBin(String input) {
        StringBuilder builder = new StringBuilder();
        String rawBinary = null;
        String formattedBinary = null;
        for (int i = 0; i < input.length(); i++) {
            rawBinary = Integer.toBinaryString(input.charAt(i));
            formattedBinary = String.format("%7s", rawBinary).replace(" ", "0");
            builder.append(formattedBinary);
        }

        return builder.toString();
    }

    private static String toChuck(String binary) {
        StringBuilder builder = new StringBuilder();
        int count = 0;
        for (int i = 0; i < binary.length(); i++) {
            if (i != 0 && binary.charAt(i) != binary.charAt(i-1)) {
                if (binary.charAt(i-1) == '1') {
                    builder.append("0 ");
                } else {
                    builder.append("00 ");
                }
                for (int j = 0; j < count; j++) {
                    builder.append("0");
                }
                builder.append(" ");
                count = 1;
                continue;
            }
            count += 1;
        }
        if (binary.charAt(binary.length() - 1) == '1') {
            builder.append("0 ");
        } else {
            builder.append("00 ");
        }
        for (int j = 0; j < count; j++) {
            builder.append("0");
        }
        return builder.toString();
    }

    private String encode(String input) {
        String encodedToBinary = toBin(input);
        return toChuck(encodedToBinary);
    }

    public void start() {
        Scanner scanner = new Scanner(System.in);
        System.out.println("Input string:");
        String input = scanner.nextLine();
        System.out.println("Encoded string:");
        System.out.println(encode(input));
        System.out.println();
        return;
    }
}
