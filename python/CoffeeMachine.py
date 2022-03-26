def get_input():
    print("Write action (buy, fill, take, remaining, exit):")
    action = input()
    return action


class Machine:
    def __init__(self):

        self.water = 400
        self.milk = 540
        self.coffee = 120
        self.cups = 9
        self.money = 550
        while True:
            action = get_input()
            if action == "buy":
                self.buy()
            elif action == "fill":
                self.fill()
            elif action == "take":
                self.take()
            elif action == "remaining":
                self.show_contents()
            elif action == "exit":
                break

    def check_resources(self, option):
        self.test_remove_ing(option)
        if self.cups < 0:
            self.reverse_remove_ing(option)
            return False, "cups"
        elif self.water < 0:
            self.reverse_remove_ing(option)
            return False, "water"
        elif self.coffee < 0:
            self.reverse_remove_ing(option)
            return False, "coffee"
        elif self.milk < 0:
            self.reverse_remove_ing(option)
            return False, "milk"
        else:
            self.reverse_remove_ing(option)
            return True, ""

    def show_contents(self):
        print(f"""The coffee machine has:
    {self.water} ml of self.water
    {self.milk} ml of self.milk
    {self.coffee} g of self.coffee beans
    {self.cups} disposable self.cups
    ${self.money} of self.money""")

    def fill(self):
        print("Write how many ml of water you want to add:")
        option = input()
        self.water += int(option)
        print("Write how many ml of milk you want to add:")
        option = input()
        self.milk += int(option)
        print("Write how many grams of coffee beans you want to add:")
        option = input()
        self.coffee += int(option)
        print("Write how many disposable cups of coffee you want to add:")
        option = input()
        self.cups += int(option)

    def buy(self):
        print("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
        option = input()
        if option != "back":
            has_resources, missing = self.check_resources(option)
            if has_resources:
                print("I have enough resources, making you a coffee!")
                self.remove_ing(option)
            else:
                print(f"Sorry, not enough {missing}!")

    def reverse_remove_ing(self, option):
        if option == "1":
            self.water += 250
            self.coffee += 16
            self.money -= 4
            self.cups += 1
        elif option == "2":
            self.water += 350
            self.milk += 75
            self.coffee += 20
            self.money -= 7
            self.cups += 1
        elif option == "3":
            self.water += 200
            self.milk += 100
            self.coffee += 12
            self.money -= 6
            self.cups += 1

    def remove_ing(self, option):
        if option == "1":
            self.water = max(self.water - 250, 0)
            self.coffee = max(self.coffee - 16, 0)
            self.money += 4
            self.cups = max(self.cups - 1, 0)
        elif option == "2":
            self.water = max(self.water - 350, 0)
            self.milk = max(self.milk - 75, 0)
            self.coffee = max(self.coffee - 20, 0)
            self.money += 7
            self.cups = max(self.cups - 1, 0)
        elif option == "3":
            self.water = max(self.water - 200, 0)
            self.milk = max(self.milk - 100, 0)
            self.coffee = max(self.coffee - 12, 0)
            self.money += 6
            self.cups = max(self.cups - 1, 0)

    def test_remove_ing(self, option):
        if option == "1":
            self.water -= 250
            self.coffee -= 16
            self.money += 4
            self.cups -= 1
        elif option == "2":
            self.water -= 350
            self.milk -= 75
            self.coffee -= 20
            self.money += 7
            self.cups -= 1
        elif option == "3":
            self.water -= 200
            self.milk -= 100
            self.coffee -= 12
            self.money += 6
            self.cups -= 1

    def take(self):
        print(f"I gave you ${self.money}")
        self.money = 0


m = Machine()
