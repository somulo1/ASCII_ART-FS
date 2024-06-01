# ASCII Art Generator

## Description
This is a simple command-line tool written in Go that generates ASCII art based on standard text. It reads characters from a standard text file and converts them into ASCII art representation.

## Usage
1. **Clone Repository**: Clone this repository to your local machine.
2. **Install go**: Ensure you have Go installed on your machine.
3. **Run Program**: Execute the program by providing the input string as a command-line argument. For example:
    ```bash
    go run . "Hello"
    ```
    This will generate ASCII art representing the word "Hello" as shown below.
    ```bash
     _              _   _          
    | |            | | | |         
    | |__     ___  | | | |   ___   
    |  _ \   / _ \ | | | |  / _ \  
    | | | | |  __/ | | | | | (_) | 
    |_| |_|  \___| |_| |_|  \___/  
                                
                                                          
    ```                                
    or you can have an additional argument to represent the banner file you'll use.

    Usage: go run . [STRING] [BANNER]

    EX: go run . hello standard

    ```console
    $ go run . "hello" standard | cat -e
     _              _   _          $
    | |            | | | |         $
    | |__     ___  | | | |   ___   $
    |  _ \   / _ \ | | | |  / _ \  $
    | | | | |  __/ | | | | | (_) | $
    |_| |_|  \___| |_| |_|  \___/  $
                                   $
                                   $
    $ go run . "Hello There!" shadow | cat -e
                                                                                             $
    _|    _|          _| _|                _|_|_|_|_| _|                                  _| $
    _|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
    _|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
    _|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
    _|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                             $
                                                                                             $
    $ go run . "Hello There!" thinkertoy | cat -e
                                                    $
    o  o     o o           o-O-o o                o $
    |  |     | |             |   |                | $
    O--O o-o | | o-o         |   O--o o-o o-o o-o o $
    |  | |-' | | | |         |   |  | |-' |   |-'   $
    o  o o-o o o o-o         o   o  o o-o o   o-o O $
                                                    $
                                                    $
    ```
    
4. **Line Breaks**: You can provide a newline character (`\n`) to create a line break in the ASCII art. For example:
    ```bash
    go run . "Hello\nThere"
     _    _          _   _          
    | |  | |        | | | |         
    | |__| |   ___  | | | |   ___   
    |  __  |  / _ \ | | | |  / _ \  
    | |  | | |  __/ | | | | | (_) | 
    |_|  |_|  \___| |_| |_|  \___/  
                                    
                                    
     _______   _                           
    |__   __| | |                          
       | |    | |__     ___   _ __    ___  
       | |    |  _ \   / _ \ | '__|  / _ \ 
       | |    | | | | |  __/ | |    |  __/ 
       |_|    |_| |_|  \___| |_|     \___| 
                                        
                                        

    ```
    This will generate ASCII art for "Hello" followed by a newline, and then ASCII art for "World".

## Arguments
- **Input Text**: The program takes one argument, which is the input text for which ASCII art is to be generated. This text can contain alphanumeric characters, special characters, and newline characters (`\n`).

## Files
- **main.go**: Contains the main program logic, including argument parsing and reading from the standard text file.
- **ascii_art.go**: Contains the function for generating ASCII art based on the input text.

## Standard Text File
The ASCII art characters are stored in a standard text file named `standard.txt`. Each character's representation occupies 8 lines in the file.

## Errors
- Only one or two arguments is allowed to be input by the user
- If too many or too few arguments are provided, the program will display an error message and exit.
- If there is an error reading the banner file, the program will display an error message and exit.

## Dependencies
- This program does not have any external dependencies.

