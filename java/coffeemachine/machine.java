package machine;

import java.util.Scanner;

public class Machine {
    // class variables
    // create scanner to get user input
    public final Scanner scanner = new Scanner(System.in);

    // water in machine
    private  int waterInMachine = 400;

    // milk in machine
    private  int milkInMachine = 540;

    // beans in machine
    private  int beansInMachine = 120;

    // disposable cups in machine
    private  int cupsInMachine = 9;

    // money in machine
    private  int moneyInMachine = 550;

    // ing quantities needed for 1 of each type of coffee
     int waterForEspresso = 250;
     int milkForEspresso = 0;
     int beansForEspresso = 16;
     int moneyForEspresso = 4;
     int waterForLatte = 350;
     int milkForLatte = 75;
     int beansForLatte = 20;
     int moneyForLatte = 7;
     int waterForCapp = 200;
     int milkForCapp = 100;
     int beansForCapp = 12;
     int moneyForCapp = 6;

    // exit status
    public boolean exit = false;


    // class methods
    // receive user action and act on it
    public void processUserAction(String userAction) {
        switch (userAction) {
            case "buy":
                buyCoffee();
                break;
            case "fill":
                fillMachine();
                break;
            case "take":
                takeMoney();
                break;
            case "remaining":
                printIng();
                System.out.println();
                break;
            case "exit":
                exit = true;
                break;
        }
    }

    private  void takeMoney() {
        System.out.printf("I gave you $%d\n", moneyInMachine);
        moneyInMachine = 0;
    }

    // fill machine with ing
    private  void fillMachine() {
        // get ing quantities
        int waterToAdd = getWaterToAdd();
        int milkToAdd = getMilkToAdd();
        int beansToAdd = getBeansToAdd();
        int cupsToAdd = getCupsToAdd();

        // add quantities to machine
        waterInMachine += waterToAdd;
        milkInMachine += milkToAdd;
        beansInMachine += beansToAdd;
        cupsInMachine += cupsToAdd;
    }

    // get cups to add in machine
    private  int getCupsToAdd() {
        System.out.println("Write how many disposable cups of coffee you want to add:");
        return scanner.nextInt();
    }

    // get beans to add in machine
    private  int getBeansToAdd() {
        System.out.println("Write how many grams of coffee beans you want to add:");
        return scanner.nextInt();
    }

    // get milk to add in machine
    private  int getMilkToAdd() {
        System.out.println("Write how many ml of milk you want to add:");
        return scanner.nextInt();
    }

    // get water to add in machine
    private  int getWaterToAdd() {
        System.out.println("Write how many ml of water you want to add:");
        return scanner.nextInt();
    }

    // buy coffee
    private  void buyCoffee() {

        // declare int variable for choice
        int coffeeToBuyChoiceInt;

        // get what user wants to buy
        String coffeeToBuyChoice = getCoffeeToBuyChoice();

        // if user didn't choose back, buy coffee, else does nothing
        // convert choice to int
        if (!"back".equals(coffeeToBuyChoice)) {
            coffeeToBuyChoiceInt = Integer.parseInt(coffeeToBuyChoice);

            // if there are ing, make coffee, else warn the user
            if (enoughIngCheckAndPrint(coffeeToBuyChoiceInt)) {

                // remove ingredients from chosen coffee
                removeIngredients(coffeeToBuyChoiceInt);
                System.out.println("I have enough resources, making you a coffee!");
            }
        }
        // blank line
        System.out.println();
    }

    // check for enough ing, if not, print warning
    private  boolean enoughIngCheckAndPrint(int coffeeToBuyChoice) {
        // number of espressos that can be made
        String missingIng = calculateMissingIng(coffeeToBuyChoice);

        // if no ing is missing, return true
        if ("none".equals(missingIng)) {
            return true;
        } else {

            // print warning and return false
            System.out.printf("Sorry, not enough %s!\n", missingIng);
            return false;
        }
    }

    // remove ingredients from machine according to chosen coffee
    private  void removeIngredients(int coffeeToBuyChoice) {

        switch (coffeeToBuyChoice) {
            // espresso
            case 1:
                waterInMachine -= waterForEspresso;
                beansInMachine -= beansForEspresso;
                moneyInMachine += moneyForEspresso;
                cupsInMachine--;
                break;

            // latte
            case 2:
                waterInMachine -= waterForLatte;
                milkInMachine -= milkForLatte;
                beansInMachine -= beansForLatte;
                moneyInMachine += moneyForLatte;
                cupsInMachine--;
                break;

            // cappuccino
            case 3:
                waterInMachine -= waterForCapp;
                milkInMachine -= milkForCapp;
                beansInMachine -= beansForCapp;
                moneyInMachine += moneyForCapp;
                cupsInMachine--;
                break;
        }
    }

    // get the coffee user wants to buy
    private  String getCoffeeToBuyChoice() {
        System.out.println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:");
        return scanner.nextLine();
    }

    // get user choice of action
//    private  String getUserAction() {
//        System.out.println("Write action (buy, fill, take, remaining, exit):");
//        return scanner.nextLine();
//    }

    // print current ing
    private  void printIng() {
        System.out.printf("The coffee machine has:\n" +
                "%d ml of water\n" +
                "%d ml of milk\n" +
                "%d g of coffee beans\n" +
                "%d disposable cups\n" +
                "$%d of money\n", waterInMachine, milkInMachine, beansInMachine, cupsInMachine, moneyInMachine);
    }

    // print needed ingredients
    private  void printIngNeeded(int cups, int mlMilk, int mlWater, int gramsBeans){
        System.out.printf("For %d cups of coffee you will need:\n" +
                "%d ml of water\n" +
                "%d ml of milk\n" +
                "%d g of coffee beans\n", cups, mlMilk, mlWater, gramsBeans);
    }

    // return number of cups of coffee user wants
    private  int getCups(){
        System.out.println("Write how many cups of coffee you will need: ");
        return scanner.nextInt();
    }


    // return grams of beans needed for cups
    private  int calculateBeans(int cups){
        return 15 * cups;
    }

    // return ml of milk needed for cups
    private  int calculateMilk(int cups){
        return 200 * cups;
    }

    // return ml of water needed for cups
    private  int calculateWater(int cups) {
        return 50 * cups;
    }

    private  void printCupsConclusion(int cups, int cupsCanBeMade) {
        // check if number of cups user wanted can be made, not, or in excess
        if (cupsCanBeMade == cups) {
            System.out.println("Yes, I can make that amount of coffee");
        } else if (cupsCanBeMade > cups) {
            System.out.printf("Yes, I can make that amount of coffee (and even %d more than that)\n", cupsCanBeMade - cups);
        } else {
            System.out.printf("No, I can make only %d cup(s) of coffee\n", cupsCanBeMade);
        }
    }

    // return missing ing for choice
    private  String calculateMissingIng(int coffeeToBuyChoice) {

        // declare water, milk, and beans needed;
        int waterNeeded, milkNeeded, beansNeeded;

        // give values according to coffee choice
        switch (coffeeToBuyChoice) {
            case 1:
                waterNeeded = waterForEspresso;
                milkNeeded = milkForEspresso;
                beansNeeded = beansForEspresso;
                break;
            case 2:
                waterNeeded = waterForLatte;
                beansNeeded = beansForLatte;
                milkNeeded = milkForLatte;
                break;
            case 3:
                waterNeeded = waterForCapp;
                milkNeeded = milkForCapp;
                beansNeeded = beansForCapp;
                break;
            default:
                throw new IllegalStateException("Unexpected value: " + coffeeToBuyChoice);
        }

        // how many cups with this amount of water
        int cupsPossibleWater = waterInMachine / waterNeeded;

        // if zero possible
        if (cupsPossibleWater == 0) {
            return "water";
        }

        // how many with this amount of milk
        int cupsPossibleMilk = milkNeeded == 0 ? Integer.MAX_VALUE : milkInMachine / milkNeeded;

        // if zero possible
        if (cupsPossibleMilk == 0) {
            return "milk";
        }

        // how many with this amount of beans
        int cupsPossibleBeans = beansInMachine / beansNeeded;

        // if zero possible
        if (cupsPossibleBeans == 0) {
            return "beans";
        }

        // how many with this amount of cups
        int cupsPossibleCups = cupsInMachine;

        // if zero possible
        if (cupsPossibleCups == 0) {
            return "cups";
        }

        // if no ing missing
        return "none";
    }

    // return beans in machine
    private  int getBeansInMachine() {
        System.out.println("Write how many grams of coffee beans the coffee machine has:");
        return scanner.nextInt();
    }

    // return milk in machine
    private  int getMilkInMachine() {
        System.out.println("Write how many ml of milk the coffee machine has:");
        return scanner.nextInt();
    }

    // return water in machine
    private  int getWaterInMachine() {
        System.out.println("Write how many ml of water the coffee machine has:");
        return scanner.nextInt();
    }

}
