package machine;

import java.util.Scanner;

public class CoffeeMachine {


    // main program
    public static void main(String[] args) {
        Machine machine = new Machine();

        while (!machine.exit) {
            // get user action
            System.out.println("Write action (buy, fill, take, remaining, exit):");
            String userAction = machine.scanner.nextLine();

            // blank line
            System.out.println();

            // process on user action
            machine.processUserAction(userAction);
        }
    }
}
