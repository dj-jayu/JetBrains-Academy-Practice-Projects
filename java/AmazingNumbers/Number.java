package numbers;

import java.util.ArrayList;
import java.util.List;

// class representing one number and it's properties
public class Number {
    // declare properties
    String formattedOneNumber;
    long number;
    String numberStr;
    String formattedTwoNumbers;
    boolean even;
    boolean odd;
    boolean buzz;
    boolean duck;
    boolean palindromic;
    boolean gapful;
    boolean spy;
    boolean sunny;
    boolean square;
    boolean jumping;
    boolean happy;
    boolean sad;

    // getters

    public boolean isHappy() {
        return this.happy;
    }

    public boolean isSad() {
        return this.sad;
    }

    public boolean isJumping() {
        return jumping;
    }

    public String getFormattedOneNumber() {
        return formattedOneNumber;
    }

    public String getFormattedTwoNumbers() {
        return formattedTwoNumbers;
    }

    public long getNumber() {
        return number;
    }

    public String getNumberStr() {
        return numberStr;
    }

    public boolean isEven() {
        return even;
    }

    public boolean isOdd() {
        return odd;
    }

    public boolean isBuzz() {
        return buzz;
    }

    public boolean isDuck() {
        return duck;
    }

    public boolean isPalindromic() {
        return palindromic;
    }

    public boolean isGapful() {
        return gapful;
    }

    public boolean isSpy() {
        return spy;
    }

    public boolean isSunny() {return sunny;}
    
    public boolean isSquare() {
        return this.square;
    }

    // setters
    public void setHappy() {
        // number to be analysed
        long numberBeingTested = this.number;

        // starts looping and stop when arriving at one or the initial number
        Long sum;
        List<Long> numbersTestedArray = new ArrayList<>();
        while (true) {
            // array to hold numbers already tested


            // make number a string
            String numberBeingTestedStr = String.valueOf(numberBeingTested);

            // make number have the value of sum, reset sum to zero
            sum = 0L;

            // sum alg from number, squared, to number
            for (int i = 0; i < numberBeingTestedStr.length(); i++) {
                sum += (long) Math.pow(Long.parseLong(String.valueOf(numberBeingTestedStr.charAt(i))), 2);
            }

            // check if number is one, or this.number
            if (sum == 1) {
                this.happy = true;
                break;
            }
            if (numbersTestedArray.contains(sum)) {
                this.happy = false;
                break;
            } else {
                numbersTestedArray.add(sum);
            }
            numberBeingTested = sum;
//            System.out.println("numberBeingTestedStr = " + numberBeingTestedStr);
        }
    }

    public void setSad() {
        this.sad = !this.happy;
    }

    public void setJumping() {
        // assumes number is jumping
        this.jumping = true;

        // loop though numbers looking for an indication it is not jumping
        for(int i = 0; i < this.numberStr.length() - 1; i++) {
            // create ints from the two numbers to compare
            int firstNumber = Integer.parseInt(numberStr.substring(i, i+1));
            int secondNumber = Integer.parseInt(numberStr.substring(i+1, i+2));

            // subtract the two numbers
            int subtractionValue = Math.abs(firstNumber - secondNumber);

            // if subtraction value bigger than one, set jumping to false
            if(subtractionValue != 1) {
                this.jumping = false;
                break;
            }
        }
    }

    private void setSquare() {
        this.square = Math.sqrt((this.number)) % 1 == 0;
    }

    private void setSunny() {
        this.sunny = Math.sqrt((this.number + 1)) % 1 == 0;
    }

    private void setDuck() {
        this.duck = this.numberStr.substring(1).contains("0");
    }

    private void setSpy() {
        // sum of digits
        long sum = 0;

        // product of digits
        long product = 1;

        // calculate sum and product
        for(String s : this.numberStr.split("")) {
            sum += Long.parseLong(s);
            product *= Long.parseLong(s);
        }

        // if they are equal or not set value of setspy
        this.spy = sum == product;

    }

    private void setGapful() {
        // variable to hold number divisor
        long divisor = 0;

        // if number has 0, 1, or 2 alg, it can't be a gapful number
        if(this.numberStr.length() < 3) {
            this.gapful = false;
        } else {
            // construct divisor
            divisor = Long.valueOf(this.numberStr.substring(0, 1) + this.numberStr.substring(this.numberStr.length() - 1, this.numberStr.length()));
            // check if number is gapful
            if(this.number % divisor == 0) {
                this.gapful = true;
            } else {
                this.gapful = false;
            }
        }
    }

    private void setPalindromic() {
        StringBuilder numberStrBuilder = new StringBuilder(this.numberStr);
        this.palindromic = numberStrBuilder.reverse().toString().equals(this.numberStr);
    }

    private void setBuzz() {
        if (this.number % 7 == 0 && this.number % 10 == 7) {
            this.buzz = true;
        } else if (this.number % 7 == 0) {
            this.buzz = true;
        } else if (this.number % 10 == 7) {
            this.buzz = true;
        } else {
            this.buzz = false;
        }
    }

    private void setEven() {
        this.even = this.number % 2 == 0;
        this.odd = !this.even;
    }

    // the string that will be printed when user enters 1 number
    public void setFormattedOneNumber() {
        this.formattedOneNumber = String.format("Properties of %d\n\t\tbuzz: %b\n\t\tduck: %b\n\t\tpalindromic: %b\n\t\tgapful: %b\n\t\tspy: %b\n\t\tsquare: %b\n\t\tsunny: %b\n\t\teven: %b\n\t\todd: %b\n\t\tjumping: %b\n\t\thappy: %b\n\t\tsad: %b\n", this.getNumber(), this.isBuzz(), this.isDuck(), this.isPalindromic(), this.isGapful(),
        this.isSpy(), this.isSquare(), this.isSunny(), this.isEven(), this.isOdd(), this.isJumping(), this.isHappy(), this.isSad());
    }

    // the string that will be printed when user enters 2 numbers
    private void setFormattedTwoNumbers() {
        // base string to print
        StringBuilder baseString = new StringBuilder(String.format("%,d", this.number) + " is");

        // variable to control correct placement of comma
        String comma = "";
        // loop and add to base string
        if(this.isBuzz()) {
            baseString.append(comma + " buzz");
            comma = ",";
        }
        if(this.isDuck()) {
            baseString.append(comma + " duck");
            comma = ",";
        }
        if(this.isGapful()) {
            baseString.append(comma + " gapful");
            comma = ",";
        }
        if(this.isPalindromic()) {
            baseString.append(comma + " palindrome");
            comma = ",";
        }
        if(this.isEven()) {
            baseString.append(comma + " even");
            comma = ",";
        }
        if(this.isOdd()) {
            baseString.append(comma + " odd");
            comma = ",";
        }
        if(this.isSpy()) {
            baseString.append(comma + " spy");
            comma = ",";
        }
        if(this.isSunny()) {
            baseString.append(comma + " sunny");
            comma = ",";
        }
        if(this.isSquare()) {
            baseString.append(comma + " square");
            comma = ",";
        }
        if(this.isJumping()) {
            baseString.append(comma + " jumping");
            comma = ",";
        }
        if(this.isHappy()) {
            baseString.append((comma + " happy"));
            comma = ",";
        }
        if(this.isSad()) {
            baseString.append((comma + " sad"));
            comma = ",";
        }
        this.formattedTwoNumbers = baseString.toString();
    }

    // constructor
    public Number(long number) {
        this.number = number;
        this.numberStr = String.valueOf(this.number);
        // set properties
        this.setEven();
        this.setBuzz();
        this.setDuck();
        this.setPalindromic();
        this.setGapful();
        this.setSpy();
        this.setSunny();
        this.setSquare();
        this.setJumping();
        this.setHappy();
        this.setSad();
        // set formated text to print
        this.setFormattedTwoNumbers();
        this.setFormattedOneNumber();
    }

    // receive a string with the property name and returns the property boolean value
    public boolean checkProperty(String property) {
        switch(property) {
            case "square":
                return this.isSquare();
            case "sunny":
                return this.isSunny();
            case "even":
                return this.isEven();
            case "odd":
                return this.isOdd();
            case "buzz":
                return this.isBuzz();
            case "duck":
                return this.isDuck();
            case "spy":
                return this.isSpy();
            case "gapful":
                return this.isGapful();
            case "palindromic":
                return this.isPalindromic();
            case "jumping":
                return this.isJumping();
            case "happy":
                return this.isHappy();
            case "sad":
                return this.isSad();
        }
        return false;
    }
}
