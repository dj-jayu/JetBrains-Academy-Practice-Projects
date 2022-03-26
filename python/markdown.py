# write your code here
markdown_commands = ["plain", "bold", "italic", "header", "link", "inline-code", "ordered-list", "unordered-list",
                     "new-line", "ordered-list", "unordered-list"]
special_commands = ["!help", "!done"]
possible_commands = markdown_commands + special_commands
final_text = ""


def write_ol(text_to_add, final_text):
    text_list = text_to_add.rstrip(",").split(",")
    for index, text in enumerate(text_list):
        final_text += f"{index + 1}. {text}\n"
    return final_text


def write_uol(text_to_add, final_text):
    text_list = text_to_add.rstrip(",").split(",")
    for text in text_list:
        final_text += f"* {text}\n"
    return final_text


def generate_list(rows, command, final_text):
    text_to_add = ""
    for i in range(rows):
        text_to_add += input(f"Row #{i + 1}: ") + ","
    if command == "ordered-list":
        return write_ol(text_to_add, final_text)

    elif command == "unordered-list":
        return write_uol(text_to_add, final_text)


def process_command(command, final_text):
    if command == "header":
        while True:
            level = int(input("Level: "))
            if level not in range(1, 7):
                print("The level should be within the range of 1 to 6")
            else:
                break
        text = input("Text: ")
        return final_text + "#" * level + " " + text + "\n"

    elif command in ["plain", "bold", "italic", "inline-code"]:
        text = input("Text: ")
        final_text = generate_text(command, text, final_text)
        return final_text
    elif command == "link":
        label = input("Label: ")
        url = input("URL: ")
        return final_text + f"[{label}]({url})"
    elif command == "new-line":
        return final_text + "\n"
    elif command == "ordered-list" or command == "unordered-list":
        while True:
            rows = int(input("Number of rows: "))
            if rows > 0:
                break
            else:
                print("The number of rows should be greater than zero")
        final_text = generate_list(rows, command, final_text)
        return final_text


def generate_text(command, text, final_text):
    if command == "plain":
        return final_text + text
    elif command == "bold":
        return final_text + "**" + text + "**"
    elif command == "italic":
        return final_text + "*" + text + "*"
    elif command == "inline-code":
        return final_text + "`" + text + "`"


while True:
    command = input("Choose a formatter: ")
    if command not in possible_commands:
        print("Unknown formatting type or command")
    elif command == "!help":
        print("""Available formatters: plain bold italic header link inline-code ordered-list unordered-list new-line
Special commands: !help !done""")
    elif command == "!done":
        with open("output.md", "w") as f:
            f.write(final_text)
        break
    else:
        final_text = process_command(command, final_text)
        print(final_text)
