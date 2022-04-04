package numbers;

// import scanner
import java.util.*;

// class for part of program that will control everything
public class NumberChecker {

    // create exit status to control exiting the program
    private boolean exit = false;

    // create number array to add numbers during the program
    List<Number> numberArray = new ArrayList<>();

    // create list of possible property values
    final List<String> properties = new ArrayList<String>(List.of(new String[]{"even", "odd", "buzz", "duck", "palindromic", "gapful", "spy", "sunny", "square", "jumping", "happy", "sad"}));

    // receives input
    // and call the appropriate methods
    public void processInput(String input) {
        // split string into arraylist
        ArrayList<String> splitString = new ArrayList<String>(List.of(input.split(" ")));

        // split the string to see if there are one or two or three words
        if(splitString.size() == 1) {
            // if one words, pass to process one number method
            processOneNumber(input);
        } else if (splitString.size() == 2) {
            // if two words, pass to process two method
            processTwoNumbers(splitString);
            }  else if (splitString.size() >= 3) {
           // if three or more words, pass to process four method
            processFourNumbers(splitString);
        }
    }

    private void processFourNumbers(ArrayList<String> splitString) {
        // separate input word. 1,2 are the numbers. 3 is the first property
        String firstInput = splitString.get(0);
        String secondInput = splitString.get(1);
        String thirdInput = splitString.get(2).toLowerCase();

        // create list array of all properties: wanted + rejected
        // the sublist is to remove the two numbers and let only the properties
        ArrayList<String> allPropertiesArray = new ArrayList<String>(splitString.subList(2, splitString.size()));

        // create list array of wanted properties
        List<String> propertiesArray = (List<String>) allPropertiesArray.clone();
        propertiesArray.removeIf(element -> element.startsWith("-"));

        // create list array of rejected properties
        List<String> propertiesRejectedArray = (List<String>) allPropertiesArray.clone();
        propertiesRejectedArray.removeIf(element -> !element.startsWith("-"));

        // check if first number is natural
        if(this.checkNatural(firstInput)) {
            System.out.println("The first parameter should be a natural number or zero.");
            System.out.println();
        }
        // check if second number is natural
        else if(this.checkNatural(secondInput)) {
            System.out.println("The second parameter should be a natural number.");
            System.out.println();
        }
        // if the numbers are valid
        else {
            // create array with wrong (not available) properties
            List<String> wrongProperties = checkWrongProperties(propertiesArray, propertiesRejectedArray);

            // create array with incompatible properties (e.g. even x odd)
            List<String> incompatiblePropertiesArray = checkMutuallyExclusiveProperties(allPropertiesArray);

            // if there are wrong properties, print them:
            // if only one property is wrong
            if(wrongProperties.size() == 1) {
                System.out.printf("The property [%s] is wrong.\n" +
                        "Available properties: %s", thirdInput, Arrays.toString(this.properties.toArray()));
                System.out.println();
            }
            // if 2 or more properties are wrong
            else if(wrongProperties.size() > 1) {
                System.out.printf("The properties %s are wrong.\n" +
                        "Available properties: %s", Arrays.toString(wrongProperties.toArray()), Arrays.toString(this.properties.toArray()));
                System.out.println();
            }
            // if there are incompatible properties, print them
            else if(incompatiblePropertiesArray.size() > 0) {
                System.out.printf("The request contains mutually exclusive properties: %s\n" +
                        "There are no numbers with these properties.", Arrays.toString(incompatiblePropertiesArray.toArray()));
                System.out.println();
            } else {
                // if everything is right with the input, print the numbers and their properties
                printTwoNumbersTwoOptions(Long.valueOf(firstInput), Long.valueOf(secondInput), propertiesArray, propertiesRejectedArray);
            }

        }

    }

    // filter array of properties and return only the unavailable ones
    private List<String> checkWrongProperties(List<String> propertiesArray, List<String> propertiesRejectedArray) {
        // create array of properties not in class property list and not in propertiesR
        List<String> wrongPropertiesArray = new ArrayList<>(propertiesArray);
        wrongPropertiesArray.addAll(propertiesRejectedArray);
        wrongPropertiesArray.removeIf(element -> this.properties.contains(element.replace("-", "")));
        return wrongPropertiesArray;
    }

    // print numbers accordingly to given properties
    private void printTwoNumbersTwoOptions(long firstInput, long secondInput, List<String> propertiesArray, List<String> propertiesRejectedArray) {
        // variable to count how many numbers have been printed
        int printedCount = 0;

        // number to start from printing
        long number = firstInput;

        // print until total printed numbers is the number defined by the second input
        while(printedCount < secondInput) {
            // initialize new number object (so properties are automatically calculated by number class constructor)
            Number newNumberObject = new Number(number);

            // print the number if it has asked properties and doesn't have rejected ones
            if (checkAllPropertiesInNumber(newNumberObject, propertiesArray) && checkAllRejectedPropertiesNotInNumber(newNumberObject, propertiesRejectedArray)) {
                System.out.println(newNumberObject.getFormattedTwoNumbers());
                // counts as one more printed number
                printedCount++;
            }
            // next number to be analised
            number++;
        }
        // blank line
        System.out.println();
    }

    // returns true if given number has no property rejected by the user
    private boolean checkAllRejectedPropertiesNotInNumber(Number newNumberObject, List<String> propertiesRejectedArray) {
        // try to find a property reject in the number
        for(String property : propertiesRejectedArray) {
            if(newNumberObject.checkProperty(property.replace("-", ""))) {
                return false;
            }
        }
        // if newNumberObject has no properties rejected, return true
        return true;
    }

    // filter the given array for incompatible properties (e.g. even and odd)
    private List<String> checkMutuallyExclusiveProperties(List<String> propertiesArray) {
        // array of incompatible properties
        List<String> incompatiblePropertiesArray = new ArrayList<>();
        // check every pair looking for exclusive properties
        for(int i = 0; i < propertiesArray.size() - 1; i++) {
            for(int j = i + 1; j < propertiesArray.size(); j++) {
                if(checkTwoPropertiesForIncompatibility(propertiesArray.get(i), propertiesArray.get(j))) {
                    // add incompatible property to array to be returned if it is not already there
                    if(!incompatiblePropertiesArray.contains(propertiesArray.get(i))) {
                        incompatiblePropertiesArray.add(propertiesArray.get(i));
                    }
                    // add incompatible property to array to be returned if it is not already there
                    if(!incompatiblePropertiesArray.contains(propertiesArray.get(j))) {
                        incompatiblePropertiesArray.add(propertiesArray.get(j));
                    }
                }
            }
        }
        // return array filled with incompatible properties
        return incompatiblePropertiesArray;
    }

    // returns true if given 2 properties are incompatible (can't exist at the same number)
    private boolean checkTwoPropertiesForIncompatibility(String thirdInput, String fourthInput) {
        return "even".equals(thirdInput) && "odd".equals(fourthInput) || "odd".equals(thirdInput) && "even".equals(fourthInput) ||
                "duck".equals(thirdInput) && "spy".equals(fourthInput) || "spy".equals(thirdInput) && "duck".equals(fourthInput)
                || "sunny".equals(thirdInput) && "square".equals(fourthInput) || "square".equals(thirdInput) && "sunny".equals(fourthInput)
                || thirdInput.equals("-" + fourthInput) || fourthInput.equals("-" + thirdInput)
                || "-even".equals(thirdInput) && "-odd".equals(fourthInput) || "-odd".equals(thirdInput) && "-even".equals(fourthInput);
    }

    // returns true if number given as argument has all asked properties
    private boolean checkAllPropertiesInNumber(Number newNumberObject, List<String> propertiesArray) {
        // look for a false property
        for(String property : propertiesArray) {
            if(!newNumberObject.checkProperty(property)) {
                return false;
            }
        }
        // if newNumberObject has all properties listed, return true
        return true;
    }


    // receives string with 2 numbers from processInput and coordinate actions
    public void processTwoNumbers(List<String> twoStringsArray) {
        // separate input
        String firstInput = twoStringsArray.get(0);
        String secondInput = twoStringsArray.get(1);

        // check if first number is valid
        if(this.checkNatural(firstInput)) {
            System.out.println("The first parameter should be a natural number or zero.");
            System.out.println();
        }
        // check if second number is valid
        else if(this.checkNatural(secondInput)) {
            System.out.println("The second parameter should be a natural number.");
            System.out.println();
        }
        // process the two natural numbers
        else {
            printPropertiesTwoNumbers(Long.valueOf(firstInput), Long.valueOf(secondInput));
        }
    }

    // receives 1 number from processInput and coordinate actions
    public void processOneNumber(String numberStr) {

        // declares number integer
        long numberLong;

        // Warn if number is not natural
        // else converts to int and print properties
        if (this.checkNatural(numberStr)) {
            System.out.println("The first parameter should be a natural number or zero.");
            System.out.println();
        } else if ("0".equals(numberStr)) {
            this.exit = true;
        }
        else {
            numberLong = Long.parseLong(numberStr);
            printPropertiesOneNumber(numberLong, numberStr);
        }
    }

    // receive two longs, and print their properties
    private void printPropertiesTwoNumbers(long firstInput, long secondInput) {
        // create array of the numbers
        for(long i = firstInput; i < firstInput + secondInput; i++) {
            this.numberArray.add(new Number(i));
        }
        // loop start in firstInput and ends before secondInput
        for(Number number : this.numberArray) {
            System.out.println(number.getFormattedTwoNumbers());
        }
        System.out.println();
    }

    // print numberLong properties
    private void printPropertiesOneNumber(long numberLong, String numberStr) {
        System.out.println(new Number(numberLong).getFormattedOneNumber());
        System.out.println();
    }

    // returns true if number is natural
    private boolean checkNatural(String numberStr) {
        if ((numberStr.contains(".") || numberStr.contains(",") || numberStr.matches("[^\\d,.]+") || Long.parseLong(numberStr) < 0)) {
            return true;
        } else {
            return false;
        }
    }

    // return exit status
    public boolean getExitStatus() {
        return this.exit;
    }
}
