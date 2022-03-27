package cinema;

import java.util.Arrays;
import java.util.Scanner;

public class Cinema {
    // declare class variables
     static int numberOfPurchasedTickets = 0;
     static int currentIncome = 0;

    // return number of rows from user
    private static int getNumRows(Scanner scanner) {
        System.out.println("Enter the number of rows:");
        return scanner.nextInt();
    }

    // return number of columns from user
    private static int getNumColumns(Scanner scanner) {
        System.out.println("Enter the number of seats in each row:");
        return scanner.nextInt();
    }
    // return 2d array with seats and labels
    private static String[][] generateSeats(int numRowsPrint, int numColumnsPrint) {

        String[][] seats = new String[numRowsPrint][numColumnsPrint];
        for (int i = 0; i < numRowsPrint; i++) {
            for (int k = 0; k < numColumnsPrint; k++) {
                if (i == 0 && k == 0) {
                    seats[i][k] = " ";
                } else if (i == 0) {
                    seats[i][k] = Integer.toString(k);
                } else if (k == 0) {
                    seats[i][k] = Integer.toString(i);
                } else {
                    seats[i][k] = "S";
                }
            }
        }
        return seats;

    }

    // display seats array to screen
    private static void printSeats(String[][] seats) {
        System.out.println("Cinema:");
        for (String[] row : seats) {
            System.out.println(Arrays.toString(row).replaceAll("[\\[\\],]", ""));
        }
        System.out.println();
    }

    // return price for tickets on first and second half seats
    private static int[] calculateTicketPrices(int numRows, int numColumns){
        int[] ticketPrice = new int[2];
        ticketPrice[0] = 10;

        if (numRows * numColumns <= 60) {
            ticketPrice[1] = 10;
        } else {
            ticketPrice[1] = 8;
        }
        return ticketPrice;
    }

    // return total income
    private static int calculateTotalIncome(int numRows, int numColumns, int[] ticketPrices) {
        int numFirstHalfRows = numRows / 2;
        int numSecondHalfRows = numRows - numFirstHalfRows;
        int numFirstHalfSeats = numFirstHalfRows * numColumns;
        int numSecondHalfSeats = numSecondHalfRows * numColumns;
        int priceFirstHalfSeats = ticketPrices[0] * numFirstHalfSeats;
        int priceSecondHalfSeats = ticketPrices[1] * numSecondHalfSeats;

        return  priceFirstHalfSeats + priceSecondHalfSeats;
    }

    /*
     print total income (used on previous level)
     private static void printTotalIncome(int totalIncome) {
         System.out.println("Total income:");
         System.out.println("$" + totalIncome);
     }
    */

    // add "B" to chosen seat, returns modified seats arrays
    private static String[][] reserveSeat(int seatRow, int seatColumn, String[][] seats) {
        seats[seatRow][seatColumn] = "B";
        return seats;
    }

    // print chosen seat price
    private static void printSeatTicketPrice(int seatPrice) {
        System.out.println(String.format("Ticket price: $%d", seatPrice));
        System.out.println();
    }

    // return price for chosen seat
    private static int calculateSeatTicketPrice(int seatRow, int seatColumn, int numRows, int[] ticketPrices) {
        if (seatRow <= numRows / 2) {
            return ticketPrices[0];
        } else {
            return ticketPrices[1];
        }
    }

    // return column of users chosen seat
    private static int getSeatColumn(Scanner scanner) {
        System.out.println("Enter a seat number in that row:");
        return scanner.nextInt();
    }

    // return row of users chosen seat
    private static int getSeatRow(Scanner scanner) {
        System.out.println("Enter a row number:");
        return scanner.nextInt();
    }

    // buyTickets and returns seats array
    private static String[][] buyTickets(Scanner scanner, int numRows, int numColumns, int[] ticketPrices, String[][] seats) {
        // Try to get free and valid seat position
        int seatRow;
        int seatColumn;

        while (true) {
            // get row of seat to calculate price
            seatRow = getSeatRow(scanner);

            // get column of seat to calculate price
            seatColumn = getSeatColumn(scanner);

            // blank line
            System.out.println();

            // if position out of bound
            if (seatRow > numRows || seatColumn > numColumns) {
                System.out.println("Wrong input!");
                System.out.println();
            } else if ("B".equals(seats[seatRow][seatColumn])) {
                System.out.println("That ticket has already been purchased!");
                System.out.println();
            } else {
                break;
            }
        }

        // calculate ticket price for seat
        int seatPrice = calculateSeatTicketPrice(seatRow, seatColumn, numRows, ticketPrices);

        // add sell price to currentIncome
        Cinema.currentIncome += seatPrice;

        // add 1 to number of purchased tickets
        Cinema.numberOfPurchasedTickets++;

        // print ticket price for seat
        printSeatTicketPrice(seatPrice);

        // mark seat as reserved
        seats = reserveSeat(seatRow, seatColumn, seats);
        return seats;
    }

    // print statistics, return seats array
    private static void printStatistics(int numRows, int numColumns, int[] ticketPrices) {
        int totalNumberOfSeats = numRows * numColumns;
        int numberOfPurchasedTickets = Cinema.numberOfPurchasedTickets;
        double percentageOfSeatsPurchased = 100 * ((float)numberOfPurchasedTickets / totalNumberOfSeats);
        int currentIncome = Cinema.currentIncome;
        int totalIncome = calculateTotalIncome(numRows, numColumns, ticketPrices);


        // print information
        System.out.println(String.format("Number of purchased tickets: %d", numberOfPurchasedTickets));
        System.out.println(String.format("Percentage: %.2f%%", percentageOfSeatsPurchased));
        System.out.println(String.format("Current income: $%d", currentIncome));
        System.out.println(String.format("Total income: $%d", totalIncome));
        System.out.println();
    }

    // receive userOption and call functions to do the chosen option, returns seats array
    private static String[][] processUserOption(int userOption, String[][] seats, Scanner scanner, int numRows, int numColumns, int[] ticketPrices) {
        switch (userOption) {
            case 1:
                System.out.println();
                printSeats(seats);
                return seats;
            case 2:
                System.out.println();
                seats = buyTickets(scanner, numRows, numColumns, ticketPrices, seats);
                return seats;
            case 3:
                System.out.println();
                printStatistics(numRows, numColumns, ticketPrices);
                return seats;
            case 0:
                //System.exit(1); (error on tests)
                // returning null
                return null;
        }
        return null;
    }

    // returns userOption
    private static int getUserOption(Scanner scanner) {
        System.out.println("1. Show the seats\n" +
                "2. Buy a ticket\n" +
                "3. Statistics\n" +
                "0. Exit");
        return scanner.nextInt();
    }


    // start program
    public static void main(String[] args) {

        // create scanner to receive input
        Scanner scanner = new Scanner(System.in);

        // get number of rows from user
        int numRows = getNumRows(scanner);

        // num of rows to display
        int numRowsPrint = numRows + 1;

        // get number of seats from user
        int numColumns = getNumColumns(scanner);

        // number of columns to display
        int numColumnsPrint = numColumns + 1;

        // calculate ticket prices for room
        int [] ticketPrices = calculateTicketPrices(numRows, numColumns);

        // blank line
        System.out.println();

        // declare array to store seats
        String[][] seats = generateSeats(numRowsPrint, numColumnsPrint);

        // enter main loop (getUserOption -> processUserOption)
        while (seats != null) {

            // get user choice (show seats, buy ticket, exit)
            int userOption = getUserOption(scanner);

            // process userOption
            seats = processUserOption(userOption, seats, scanner, numRows, numColumns, ticketPrices);
        }








        //----------previous level-------------
        // print seats array
        //printSeats(seats);


        //seats = buyTickets(scanner, numRows, ticketPrices, seats);

        // print seats array
        //printSeats(seats);

        // calculate total income
        // another level: int totalIncome = calculateTotalIncome(numRows, numColumns, ticketPrice);

        // print total income
        // another level: printTotalIncome(totalIncome);
    }


}

