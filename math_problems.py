import random

# Generate math expression for level one
def generate_expression(possible_symbols):
    first_number = random.randint(2, 9)
    second_number = random.randint(2, 9)
    symbol = random.choice(possible_symbols)

    return f"{first_number} {symbol} {second_number}"

# Generate list of math expressions for level one
def generate_expression_list(n, possible_symbols):
    expression_list = []
    for i in range(n):
        expression_list.append(generate_expression(possible_symbols))
    return expression_list

# Validate user input for level one answer
def is_valid(n):
    if len(n) == 1:
        return n.isdigit()
    elif n[1:].isdigit() and (n[0] == "-" or n[0].isdigit()):
        return True
    else:
        return False

# Generate random list of numbers for level 2
def generate_numbers_list(n):
    eleven_third_nine = list(range(11, 30))
    random.shuffle(eleven_third_nine)
    five_numbers = eleven_third_nine[:n]
    return five_numbers

# Main program
def main():
    possible_symbols = ["+", "-", "*"]
    number_of_correct_answers = 0
    levels = ["1", "2"]
    level_descriptions = {"1": "simple operations with numbers 2-9", \
                          "2": "integral squares 11-29"}
    # User chooses level
    while True:
        print("""Which level do you want? Enter a number:
    1 - simple operations with numbers 2-9
    2 - integral squares of 11-29""")
        level = input()
        if level not in levels:
            print("Incorrect format.")
        else:
            break
    # Starts level 1
    if level == "1":
        for expression in generate_expression_list(5, possible_symbols):
            correct_answer = str(eval(expression))
            print(expression)
            while True:
                try:
                    answer = input()
                    if not is_valid(answer):
                        raise ValueError
                    elif answer != correct_answer:
                        print("Wrong!")
                        break
                    elif answer == correct_answer:
                        print("Right!")
                        number_of_correct_answers += 1
                        break
                except ValueError:
                    print("Incorrect format.")
    # Starts level 2
    elif level == "2":
        five_numbers = generate_numbers_list(5)
        for number in five_numbers:
            while True:
                print(number)
                user_square = input()
                correct_square = str(number ** 2)
                if not user_square.isdigit():
                    print("Wrong format! Try again.")
                else:
                    break
            if user_square == correct_square:
                print("Right!")
                number_of_correct_answers += 1
            else:
                print("Wrong!")

    # Show score and ask if user wants to save in file
    print(f"Your mark is {number_of_correct_answers}/5. Would you like to save the result? Enter yes or no.")
    save_file = input()
    # Save file with scores if chosen
    if save_file in ["yes", "YES", "y", "Yes"]:
        print("What is your name?")
        name = input()
        with open("results.txt", "a") as f:
            f.write(f"{name}: {number_of_correct_answers}/5 in level {level} ({level_descriptions[level]}).")
        print("The results are saved in \"results.txt\".")

# Starts main program
main()
