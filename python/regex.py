#this code is a little mess, but it was enought to pass all the jetbrainsacademy tests
import sys

sys.setrecursionlimit(10000)


def check_regex(reg, letters):
    if reg and reg[0] == "^":
        return check_regex_equal_length(reg[1:], letters)
    return check_regex_without_symbols(reg, letters)

def get_pure_length(letters_without_symbols):
    interrogation_number = letters_without_symbols.count("?")
    asterix_count = letters_without_symbols.count("*")
    plus_count = letters_without_symbols.count("+")
    e_count = letters_without_symbols.count("$")
    barra_count = letters_without_symbols.count("\\")
    return len(letters_without_symbols) - ((interrogation_number * 2) + (asterix_count * 2) + plus_count) - e_count - barra_count


def check_regex_without_symbols(reg_without_symbols, letters_without_symbols):
    # if regex > input
    #print(reg_without_symbols, letters_without_symbols)
    if get_pure_length(reg_without_symbols) > len(letters_without_symbols):
        #print(letters_without_symbols)
        #print("tamanho")
        return False
    if reg_without_symbols != "" and reg_without_symbols[-1] != "$":
        len_without_money_symbol = len(reg_without_symbols)
    else:
        len_without_money_symbol = len(reg_without_symbols) - 1

    #print(f"chamando check_regex_eq_leng com {reg_without_symbols=} e {letters_without_symbols[:len_without_money_symbol]}")
    if check_regex_equal_length(reg_without_symbols, letters_without_symbols[:len_without_money_symbol]):
        if reg_without_symbols != "" and reg_without_symbols[-1] != "$":
            #print("testando se regex acaba com $")
            #print("regex nao acaba com $")
            #print("retornando VERDADE!")
            return True
        else:
            if letters_without_symbols.endswith(letters_without_symbols[:len_without_money_symbol]):
                #print("estamos no final da string e recebemos verdadeiro porque sobrou $")
                #print("retornando VERDADE")
                return True
            else:
                #print("nao estamos no final da string apesar de termos recebido verdadeiro porque sobrou $")
                #print(f"chamando e retornando check_regerx_sem_simbolo com {reg_without_symbols}, {letters_without_symbols[1:]}")
                return check_regex_without_symbols(reg_without_symbols, letters_without_symbols[1:])
    else:
        #print(f"a comparacao do regex com o pedaco da string {letters_without_symbols[:len_without_money_symbol]} deu FALSO")
        #print(f"chamando e retornando check_regex_sem_simbolo com {reg_without_symbols}, {letters_without_symbols[1:]}")
        return check_regex_without_symbols(reg_without_symbols, letters_without_symbols[1:])


def check_regex_equal_length(reg_equal, letters_equal):
    #print(reg_equal, letters_equal)
    if not reg_equal:
        #print("regex passado esta vazio")
        #print("retornando TRUE")
        return True
    if (reg_equal == "." or reg_equal == ".*" or reg_equal == ".+" or reg_equal == ".?") and not letters_equal:
        return True
    if reg_equal != "$" and not letters_equal:
        #print(f"string passada esta vazia, regex passado {reg_equal} nao termina em $")
        #print("retornando False")
        return False
    if reg_equal == "$" and not letters_equal:
        #print("regex passado termina em $, e string esta vazia")
        #print("retornando True")
        return True


    if reg_equal[0] == "\\":
        #print(f"{reg_equal[0]} eh \\")
        if check_one_letter(reg_equal[1], letters_equal[0]):
            #print(f"checou se {reg_equal[1]} == {letters_equal[0]} e eh")
            #print(f"retornando check equal length com {reg_equal[2:]} e {letters_equal[1:]}")
            return check_regex_equal_length(reg_equal[2:], letters_equal[1:])
        #print(f"checou se {reg_equal[1]} == {letters_equal[0]} e nao eh")
        #print("retornando falso")
        return False


    if len(reg_equal) > 1 and reg_equal[1] in ["?", "*", "+"] and not check_one_letter(reg_equal[0], letters_equal[0]):
        #print("aqui4")
        return check_regex_equal_length(reg_equal[2:], letters_equal[0:])
    # second ?, first +
    if len(reg_equal) > 1 and reg_equal[1] in ["?"] and check_one_letter(reg_equal[0], letters_equal[0]):
        #print("aqui3")
        return check_regex_equal_length(reg_equal[2:], letters_equal[1:])
    # second *, second +
    if len(reg_equal) > 2 and reg_equal[1] in ["*", "+"] and check_one_letter(reg_equal[0], letters_equal[0]) \
            and len(letters_equal) > 1 and check_one_letter(reg_equal[2], letters_equal[1]) and reg_equal[-1] != "$":
        #print("aqui2.12")
        return check_regex_equal_length(reg_equal[3:], letters_equal[2:])
    if len(reg_equal) > 2 and reg_equal[1] in ["*", "+"] and check_one_letter(reg_equal[0], letters_equal[0]) \
            and len(letters_equal) > 1 and reg_equal[2:-1] == letters_equal[1:]:
        #print("aqui2.12")
        return check_regex_equal_length(reg_equal[3:], letters_equal[2:])
    if len(reg_equal) > 2 and reg_equal[1] in ["*", "+"] and check_one_letter(reg_equal[0], letters_equal[0]) \
            and len(letters_equal) == 1:
        #print("aqui2.1")
        return check_regex_equal_length(reg_equal[1:], letters_equal[1:])
    if len(reg_equal) > 1 and reg_equal[1] in ["*", "+"] and check_one_letter(reg_equal[0], letters_equal[0]):
        #print("aqui2")
        return check_regex_equal_length(reg_equal, letters_equal[1:])

    if check_one_letter(reg_equal[0], letters_equal[0]):
        #print(f"as letras {reg_equal[0]} e {letters_equal[0]} sao iguais")
        #print(f"chamando e retornando check_regex_equal_length com {reg_equal[1:]}, {letters_equal[1:]}")
        return check_regex_equal_length(reg_equal[1:], letters_equal[1:])

    # first ?, first *, first +
    return False


def check_one_letter(reg_letter, input_letter):
    if not reg_letter:
        return True
    elif not input_letter:
        return False
    return reg_letter == input_letter or reg_letter == "."


# if __name__ == "__main__":
reg, letters = input().split("|")
print(check_regex(reg, letters))
