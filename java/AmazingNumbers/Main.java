package numbers;

// imports
import java.util.Locale;
import java.util.Scanner;

// program starts
public class Main {
    // create scanner
    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) {
//        write your code here
        // starts number checker
        NumberChecker numberChecker = new NumberChecker();

        // welcome and instructions
        System.out.println("Welcome to Amazing Numbers!\n" +
                "\n" +
                "Supported requests:\n" +
                "- enter a natural number to know its properties;\n" +
                "- enter two natural numbers to obtain the properties of the list:\n" +
                "  * the first parameter represents a starting number;\n" +
                "  * the second parameter shows how many consecutive numbers are to be processed;\n" +
                "- two natural numbers and properties to search for;\n" +
                "- a property preceded by minus must not be present in numbers;\n" +
                "- separate the parameters with one space;\n" +
                "- enter 0 to exit.");

        // get number from user
        while (!numberChecker.getExitStatus()) {
            System.out.print("Enter a request: ");
            String input = scanner.nextLine();

            // blank line
            System.out.println();

            // pass number to number checker
            numberChecker.processInput(input.toLowerCase());
        }

    }
}
