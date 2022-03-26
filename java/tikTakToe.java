package tictactoe;
import java.util.Scanner;

public class Main {

    // apply move to the board and returns the updated board
    private static char[][] applyChoice(int[] choice, char[][] positions, char turn) {
        int row = choice[0] - 1;
        int column = choice[1] - 1;
        positions[row][column] = turn;
        return positions;
    }

    // get coordinates for the next move
    private static int[] getChoice(char[][] positions) {
        String row;
        String column;
        while (true) {
            System.out.println("Enter the coordinates: ");
            Scanner scanner = new Scanner(System.in);
            row = scanner.next();
            if (!isNumerical(row)) {
                System.out.println("You should enter numbers!");
                continue;
            } else if (Integer.parseInt(row) > 3 || Integer.parseInt(row) < 1) {
                System.out.println("Coordinates should be from 1 to 3!");
                continue;
            }
            column = scanner.next();
            if (!isNumerical(column)) {
                System.out.println("You should enter numbers!");
            } else if (Integer.parseInt(column) < 1 || Integer.parseInt(column) > 3) {
                System.out.println("Coordinates should be from 1 to 3!");
            } else if (positions[Integer.parseInt(row) - 1][Integer.parseInt(column) - 1] == 'O' || positions[Integer.parseInt(row) - 1][Integer.parseInt(column) - 1] == 'X') {
                System.out.println("This cell is occupied! Choose another one!");
            } else {
                break;
            }
        }
        return new int[]{Integer.parseInt(row), Integer.parseInt(column)};


    }

    // print current board
    private static void printBoard(char[][] positions) {

        // make positions into a string
        StringBuilder s = new StringBuilder();

        for (char[] row : positions) {
            for (char letter : row) {
                s.append(letter);
            }
        }
        System.out.println("---------");

        for(int i = 0; i < 3; i++){
            System.out.print("| ");
            for(int k = i * 3; k < (i+1) * 3; k++){

                System.out.print(s.charAt(k) + " ");

            }
            System.out.println("|");
        }
        System.out.println("---------");
    }

    // check if input is numerical
    private static boolean isNumerical(String s){
        try {
            int sInt = Integer.parseInt(s);
        } catch (NumberFormatException nfe) {
            return false;
        }
        return true;
    }


    // self descriptive
    private static char[][] generateEmptyBoard() {
//        Scanner scanner = new Scanner(System.in);
//        System.out.print("Enter cells: ");
//
//        String s = scanner.nextLine();
        String s = "_________";

        // creating memory for array
        char[][] positions = new char[3][3];

        // index to access letters
        int index = 0;

        //populating array
        for (int i = 0; i < 3; i++){
            for (int k = 0; k < 3; k++){
                positions[i][k] = s.charAt(index);
                index++;
            }
        }

        return positions;

    }

    // check if there are any winning sequence of "X" or "O"
    private static boolean check3(char symbol, char[][] positions) {
        return (symbol == positions[0][0]) && symbol == positions[0][2] && symbol == positions[0][1]
                || symbol == positions[0][0] && symbol == positions[2][0] && symbol == positions[1][0]
                || symbol == positions[0][0] && symbol == positions[2][2] && symbol == positions[1][1]
                || symbol == positions[2][0] && symbol == positions[0][2] && symbol == positions[1][1]
                || symbol == positions[2][2] && symbol == positions[2][0] && symbol == positions[2][1]
                || symbol == positions[2][2] && symbol == positions[0][2] && symbol == positions[1][2]
                || symbol == positions[0][1] && symbol == positions[1][1] && symbol == positions[2][1]
                || (symbol == positions[1][1] && (symbol == positions[0][0] && symbol == positions[2][2]
                || symbol == positions[2][0] && symbol == positions[0][2]
                || symbol == positions[1][0] && symbol == positions[1][2]
                || symbol == positions[0][1] && symbol == positions[2][1]));
    }

    // check if there are empty cells
    private static boolean checkEmpty(char[][] positions) {
        for (char[] row : positions) {
            for (char position : row) {
                if (position != 'X' && position != 'O'){
                    return true;
                }
            }
        }
        return false;
    }

    // count the number o "X" on the board
    private static int countX(char[][] positions) {
        int count = 0;
        for (char[] row : positions) {
            for (char position : row) {
                if (position == 'X'){
                    count++;
                }
            }
        }
        return count;
    }

    // count the number of "O" on the board
    private static int countO(char[][] positions) {
        int count = 0;
        for (char[] row : positions) {
            for (char position : row) {
                if (position == 'O'){
                    count++;
                }
            }
        }
        return count;
    }

    // return game state
    private static String returnState(char[][] positions) {
        boolean isEmpty = checkEmpty(positions);
        boolean xWins = check3('X', positions);
        boolean oWins = check3('O', positions);
        int countX = countX(positions);
        int countO = countO(positions);
        int diff = Math.abs(countO - countX);

        if (!xWins && !oWins && isEmpty && diff < 2) {
            return "Game not finished";
        } else if (!xWins && !oWins && !isEmpty) {
            return "Draw";
        } else if ((xWins && oWins) || diff >= 2) {
            return "Impossible";
        } else if (xWins) {
            return "X wins";
        } else if (oWins) {
            return "O wins";
        } else {
            return "";
        }
    }

    // main program
    public static void main(String[] args) {
        // -----content for previous levels----------------
        // get input and print game screen
        // char[][] positions = getInput();
        // -----------------------------------------------------


        // declare variable that will hold the game state
        String state;

        // declare variable that will hold the current turn
        char turn = 'X';

        // get empty board
        char[][] positions = generateEmptyBoard();

        // print board
        printBoard(positions);

        // while the game is not over
        do {

            // get move choice
            int[] choice = getChoice(positions);

            // apply move

            positions = applyChoice(choice, positions, turn);

            // change player turn
            turn = turn == 'O' ? 'X' : 'O';

            // print board again
            printBoard(positions);

            // get state
            state = returnState(positions);

            // check if game ended
        } while (state != "X wins" && state != "O wins" && state != "Draw");

        // print final game state
        System.out.println(state);
    }




}
