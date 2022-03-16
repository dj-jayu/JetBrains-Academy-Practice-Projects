# Write your code here
import random


def get_score(name):
    with open(file="rating.txt", mode="rt") as f:
        for line in f:
            if name in line:
                return int(line.split()[1])
    return 0


class Game:
    def __init__(self):
        self.possible_commands = ["!exit", "!rating"]
        self.possible_choices = ["rock", "paper", "scissors"]
        self.score = 0

    def start_game(self):
        name = input("Enter your name: ")
        print(f"Hello, {name}")
        self.score = get_score(name)
        possible_choices = input()
        if possible_choices != "":
            self.possible_choices = possible_choices.split(",")
        self.possible_commands.extend(self.possible_choices)
        print("Okay, let's start")

        while True:
            user_choice = input()
            if user_choice not in self.possible_commands:
                print("Invalid input")
            elif user_choice == "!exit":
                print("Bye!")
                break
            elif user_choice == "!rating":
                print(f"Your rating: {self.score}")
            else:
                self.evaluate_play(user_choice)

    def evaluate_play(self, user_choice):
        computer_choice = random.choice(self.possible_choices)
        index_of_choice = self.possible_choices.index(user_choice)
        absent_list = self.possible_choices[index_of_choice:]
        absent_list.extend(self.possible_choices[:index_of_choice])
        absent_list = absent_list[1:]
        # print(f"absent list: {absent_list}")
        winning_list = absent_list[:len(absent_list)//2]
        losing_list = absent_list[len(absent_list)//2:]
        # print(f"winning_list: {winning_list}\nlosing_list: {losing_list}")

        if computer_choice in losing_list:
            print(f"Well done. The computer chose {computer_choice} and failed")
            self.score += 100
        elif computer_choice in winning_list:
            print(f"Sorry, but the computer chose {computer_choice}")
            self.score += 0
        else:
            print(f"There is a draw {computer_choice}")
            self.score += 50


game = Game()
game.start_game()
