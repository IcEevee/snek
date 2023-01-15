# import keyboard
import os
from colorama import Fore, Back

TERM_X = os.get_terminal_size().lines
TERM_Y = os.get_terminal_size().columns

cells = [(x, y) for x in range(TERM_X) for y in range(TERM_Y)]

for cell in cells:

    # Status Lines
    if cell[0] == TERM_X-2:
        # SNEK Title
        match str(cell[1]):
            case "0":
                print(f"{Back.GREEN} ", end="")
            case other:
                print(f"{Back.RESET} ", end="")

    else:
        print(f"{Back.RESET} ", end="")
