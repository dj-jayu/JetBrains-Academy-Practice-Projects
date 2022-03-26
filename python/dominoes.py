# Write your code here
# Write your code here
import random


class InvalidMoveException(Exception):
    pass


class DominoSet:
    def __init__(self):
        self.domino_set = []

    def add_domino_end_reverse_if_need(self, domino):
        if domino[0] != self.domino_set[-1][1]:
            domino.reverse()
        self.add_domino_end(domino)

    def add_domino_end(self, domino):
        self.domino_set.append(domino)

    def add_domino_start_reverse_if_need(self, domino):
        if domino[1] != self.domino_set[0][0]:
            domino.reverse()
        self.add_domino_start(domino)

    def add_domino_start(self, domino):
        self.domino_set.insert(0, domino)

    def add_dominos(self, domino_list):
        self.domino_set.extend(domino_list)

    def get_dominos_generator(self):
        return (x for x in self.domino_set)

    def __len__(self):
        return len(self.domino_set)

    def __getitem__(self, item):
        return self.domino_set[item]

    def __str__(self):
        return f"[{self.domino_set[0]}, {self.domino_set[1]}"

    def __iter__(self):
        return (x for x in self.domino_set)

    def remove(self, domino):
        self.domino_set.remove(domino)

    def get_dominos(self):
        return self.domino_set

    def pop(self):
        return self.domino_set.pop()


class FullSet(DominoSet):
    def __init__(self):
        super().__init__()
        self.create_set()

    def create_set(self):
        for i in range(7):
            for j in range(i, 7):
                self.add_domino_end(Piece([i, j]))
        random.shuffle(self.domino_set)


class ScoresDict:
    def __init__(self):
        self.base_list = [0, 1, 2, 3, 4, 5, 6, 7]
        self.base_scores_dict = dict.fromkeys(self.base_list, 0)
        self.dominos_scores_dict = {}

    def update_base_scores_dict(self, snake, hand):
        snake_list = [domino.to_list() for domino in snake.get_dominos_list()]
        hand_list = [domino.to_list() for domino in hand.get_dominos()]
        merged_list = snake_list + hand_list
        sum_dict = {i: merged_list.count(i) for i in self.base_list}
        self.base_scores_dict.update(sum_dict)

    def update_dominos_scores_dict(self, hand):
        updated_dict = {}
        for domino in hand:
            first_value = domino.to_list()[0]
            second_value = domino.to_list()[1]
            first_score = self.base_scores_dict[first_value]
            second_score = self.base_scores_dict[second_value]
            score_sum = first_score + second_score
            updated_dict[domino] = score_sum
        self.dominos_scores_dict = updated_dict

    def update_scores(self, snake, hand):
        self.update_base_scores_dict(snake, hand)
        self.update_dominos_scores_dict(hand)

    def get_domino_score(self, domino_side_list):
        return self.dominos_scores_dict[domino_side_list[0]]


class Player:
    def __init__(self):
        self.domino_set = DominoSet()
        self.scores_dict = ScoresDict()

    def pop(self):
        return self.domino_set.pop()

    def get_domino(self, index):
        return self.domino_set[index]

    def set_dominos(self, domino_list):
        self.domino_set.add_dominos(domino_list)

    def get_dominos_generator(self):
        return self.domino_set.get_dominos_generator()

    def get_dominos(self):
        return self.domino_set.get_dominos()

    def remove_domino(self, domino):
        self.domino_set.remove(domino)

    def get_dominos_size(self):
        return len(self.domino_set)

    def generate_valid_pieces_list(self, domino_snake):
        first_piece = domino_snake[0]
        last_piece = domino_snake[-1]

        # percorre lista de pecas checando qual eh valida, exclui as outras
        valid_pieces = []
        for piece in self.get_dominos():
            side = piece.is_valid(first_piece, last_piece)
            if side == "left" or side == "right":
                valid_pieces.append([piece, side])
            elif side == "both":
                valid_pieces.extend([[piece, "left"], [piece, "right"]])
        return valid_pieces

    def __str__(self):
        return self.description

    def __eq__(self, other):
        return self.description == other


class HumanPlayer(Player):
    def __init__(self):
        super().__init__()
        self.description = "player"

    def pieces_message(self):
        message = ""
        for i, domino in enumerate(self.get_dominos_generator()):
            message += f"{i + 1}:[{domino[0]}, {domino[1]}]\n"
        return message


class ComputerPlayer(Player):
    def __init__(self):
        super().__init__()
        self.description = "computer"

    def update_domino_scores(self, snake):
        self.scores_dict.update_scores(snake, self.domino_set)


class StockPieces(Player):
    pass


class Status:
    def __init__(self):
        self.next_player = random.choice(["computer", "player"])
        self.winner = None
        self.is_running = True

    def set_next_player(self, next_player):
        self.next_player = next_player

    def next(self):
        self.next_player = "player" if self.next_player == "computer" else "computer"

    def get_next_player(self):
        return self.next_player

    def update_status(self):
        pass

    def set_winner(self, winner):
        self.winner = winner


class Piece:
    def __init__(self, value_list):
        self.value_list = value_list

    def __hash__(self):
        return hash((self.value_list[0], self.value_list[1]))

    def reverse(self):
        self.value_list.reverse()

    def __repr__(self):
        return f"{self.value_list}"

    def __str__(self):
        return f"[{self.value_list[0]},{self.value_list[1]}]"

    def __eq__(self, other):
        return self.value_list == other.value_list

    def __getitem__(self, item):
        return self.value_list[item]

    def is_valid(self, first_piece, last_piece):
        if (self.value_list[0] == last_piece[1] or self.value_list[1] == last_piece[1]) and (
                self.value_list[1] == first_piece[0] or self.value_list[0] == first_piece[0]):
            return "both"
        elif self.value_list[0] == last_piece[1] or self.value_list[1] == last_piece[1]:
            return "right"
        elif self.value_list[1] == first_piece[0] or self.value_list[0] == first_piece[0]:
            return "left"
        return False

    def to_list(self):
        return self.value_list


class DominoSnake(DominoSet):
    def get_dominos(self):
        if len(self.domino_set) >= 7:
            return_str = ""
            for x in self.domino_set[0:3]:
                return_str += f"[{str(x[0])}, {str(x[1])}]"
            return_str += "..."
            for x in self.domino_set[-3:]:
                return_str += f"[{str(x[0])}, {str(x[1])}]"
            return return_str
        else:
            return_str2 = ""
            for x in self.domino_set:
                return_str2 += f"[{str(x[0])}, {str(x[1])}]"
            return return_str2

    def get_snake_csv(self):
        return ",".join([str(x) for pair in self.domino_set for x in pair])

    def get_dominos_list(self):
        return self.domino_set

    def draw(self):
        if self.domino_set[0][0] == self.domino_set[-1][1] and self.get_snake_csv().count(
                str(self.domino_set[0][0])) == 8:
            return True
        return False


class Game:
    def __init__(self):
        self.take_extra_piece = False
        self.choice_signal = None
        self.next_domino = None
        self.human_player = HumanPlayer()
        self.computer_player = ComputerPlayer()
        self.stock_pieces = StockPieces()
        self.full_set = FullSet()
        self.domino_snake = DominoSnake()
        self.status = Status()
        self.distribute_pieces()

    def distribute_pieces(self):
        first_player = None
        first_piece = None
        while not (first_player and first_piece):
            self.stock_pieces.set_dominos(self.full_set[0:14])
            self.computer_player.set_dominos(self.full_set[14:21])
            self.human_player.set_dominos(self.full_set[21:28])
            first_player, first_piece = self.get_first_player()
        self.domino_snake.add_domino_end(first_piece)
        if first_player == "computer":
            self.human_player.remove_domino(first_piece)
        elif first_player == "player":
            self.computer_player.remove_domino(first_piece)
        self.status.set_next_player(first_player)
        self.computer_player.update_domino_scores(self.domino_snake)

    def get_first_player(self):
        beginning_pieces = [Piece([6, 6]), Piece([5, 5]), Piece([4, 4]), Piece([3, 3]), Piece([2, 2]), Piece([1, 1])]
        for piece in beginning_pieces:
            if piece in self.human_player.get_dominos():
                return ComputerPlayer(), piece
            elif piece in self.computer_player.get_dominos():
                return HumanPlayer(), piece
        return 0, 0

    def next(self):
        pass

    def next_move_message(self):
        if self.status.get_next_player() == "computer":
            return "Status: Computer is about to make a move. Press Enter to continue..."
        elif self.status.get_next_player() == "player":
            return "Status: It's your turn to make a move. Enter your command."

    def still_running(self):
        return self.define_winner()

    def define_winner(self):
        if self.computer_player.get_dominos_size() == 0:
            self.status.set_winner(self.computer_player)
            return False
        elif self.human_player.get_dominos_size() == 0:
            self.status.set_winner(self.human_player)
            return False
        elif self.domino_snake.draw():
            self.status.set_winner(None)
            return False
        return True

    def final_message(self):
        if self.status.winner == "player":
            return f"Status: The game is over. You won!"
        elif self.status.winner == "computer":
            return f"Status: The game is over. The computer won!"
        elif not self.status.winner:
            return f"Status: The game is over. It's a draw!"

    def get_valid_play(self):
        next_player = self.human_player if self.status.get_next_player() == "player" else self.computer_player
        valid_pieces_list = next_player.generate_valid_pieces_list(self.domino_snake)
        while True:
            if next_player == "computer":
                input()
                if valid_pieces_list:
                    valid_pieces_list.sort(key=self.computer_player.scores_dict.get_domino_score)
                    # chose random piece
                    return valid_pieces_list[-1]
                else:
                    # chose to get new piece
                    return None
            else:
                try:
                    choice = input()
                    choice = int(choice)
                    if choice == 0:
                        return None
                    abs_choice = abs(choice)
                    choice_signal = choice // abs_choice
                    choice_side = "left" if choice_signal == -1 else "right"
                    chosen_piece = [next_player.get_domino(abs_choice - 1), choice_side]
                    if chosen_piece not in valid_pieces_list:
                        raise InvalidMoveException
                    return chosen_piece
                except (TypeError, ValueError, IndexError):
                    print("Invalid input. Please try again.")
                    continue
                except InvalidMoveException:
                    print("Illegal move. Please try again.")
                    continue

    def apply_next_move(self, move):
        if not move:
            if self.status.get_next_player() == "computer" and self.stock_pieces.domino_set:
                self.computer_player.domino_set.add_domino_end(self.stock_pieces.domino_set.pop())
            elif self.status.get_next_player() == "player" and self.stock_pieces.domino_set:
                self.human_player.domino_set.add_domino_end(self.stock_pieces.domino_set.pop())
        else:
            if self.status.get_next_player() == "computer":
                self.computer_player.remove_domino(move[0])
            else:
                self.human_player.remove_domino(move[0])

            if move[1] == "left":
                self.domino_snake.add_domino_start_reverse_if_need(move[0])
            else:
                self.domino_snake.add_domino_end_reverse_if_need(move[0])

        self.status.next()
        if self.status.get_next_player() == "computer":
            self.computer_player.update_domino_scores(self.domino_snake)


game = Game()

while game.still_running():
    print("=" * 70)
    print(f"Stock size: {game.stock_pieces.get_dominos_size()}")
    print(f"Computer pieces: {game.computer_player.get_dominos_size()}")
    print(f"\n{game.domino_snake.get_dominos()}\n")
    print("Your pieces:")
    print(game.human_player.pieces_message())
    print(game.next_move_message())
    next_move = game.get_valid_play()
    game.apply_next_move(next_move)

print("=" * 70)
print(f"Stock size: {game.stock_pieces.get_dominos_size()}")
print(f"Computer pieces: {game.computer_player.get_dominos_size()}")
print(f"\n{game.domino_snake.get_dominos()}\n")
print("Your pieces:")
print(game.human_player.pieces_message())
print(game.final_message())
