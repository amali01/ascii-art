```bash
                                                __          __         _                                           _                  
                                                \ \        / /        | |                                         | |                 
                                                 \ \  /\  / /    ___  | |   ___    ___    _ __ ___     ___        | |_    ___         
                                                  \ \/  \/ /    / _ \ | |  / __|  / _ \  | '_ ` _ \   / _ \       | __|  / _ \        
                                                   \  /\  /    |  __/ | | | (__  | (_) | | | | | | | |  __/       \ |_  | (_) |       
                                                    \/  \/      \___| |_|  \___|  \___/  |_| |_| |_|  \___|        \__|  \___/        
                                                                                                                                      
                                                                                                                                      
                                           _____    _____   _____   _____                    _____    _______         _____    _____     ____   
                                  /\      / ____|  / ____| |_   _| |_   _|           /\     |  __ \  |__   __|       |  __ \  |  __ \   / __ \  
                                 /  \    | (___   | |        | |     | |            /  \    | |__) |    | |          | |__) | | |__) | | |  | | 
                                / /\ \    \___ \  | |        | |     | |           / /\ \   |  _  /     | |          |  ___/  |  _  /  | |  | | 
                               / ____ \   ____) | | |____   _| |_   _| |_         / ____ \  | | \ \     | |          | |      | | \ \  | |__| | 
                              /_/    \_\ |_____/   \_____| |_____| |_____|       /_/    \_\ |_|  \_\    |_|          |_|      |_|  \_\  \____/  
                                                                                                                                                
                                                                                                                                                
                                                            _                _     _               _               _                              _   _                   _   _  
                                                           | |              | |   | |             | |             | |                            (_) | |                 | | | | 
             ___    _ __     ___          ___    ___     __| |   ___        | |_  | |__     __ _  | |_          __| |   ___    ___    ___         _  | |_          __ _  | | | | 
            / _ \  | '_ \   / _ \        / __|  / _ \   / _` |  / _ \       | __| |  _ \   / _` | | __|        / _` |  / _ \  / __|  / _ \       | | | __|        / _` | | | | | 
           | (_) | | | | | |  __/       | (__  | (_) | | (_| | |  __/       \ |_  | | | | | (_| | \ |_        | (_| | | (_) | \__ \ |  __/       | | \ |_        | (_| | | | | | 
            \___/  |_| |_|  \___|        \___|  \___/   \__,_|  \___|        \__| |_| |_|  \__,_|  \__|        \__,_|  \___/  |___/  \___|       |_|  \__|        \__,_| |_| |_| 
                                                                                                                                                                                 
                                                                                                                                                                                                                                                                                                     
                _                                              _               _       ______                                                  _                          
               | |                       /\                   (_)             | |     |  ____|                                     /\         | |                         
               | |__    _   _           /  \     _ __ ___      _    __ _    __| |     | |__     _ __ ___     __ _   _ __          /  \      __| |  _ __     __ _   _ __   
               | '_ \  | | | |         / /\ \   | '_ ` _ \    | |  / _` |  / _` |     |  __|   | '_ ` _ \   / _` | | '_ \        / /\ \    / _` | | '_ \   / _` | | '_ \  
               | |_) | | |_| |        / ____ \  | | | | | |   | | | (_| | | (_| |  _  | |____  | | | | | | | (_| | | | | |  _   / ____ \  | (_| | | | | | | (_| | | | | | 
               |_.__/   \__, |       /_/    \_\ |_| |_| |_|   | |  \__,_|  \__,_| ( ) |______| |_| |_| |_|  \__,_| |_| |_| ( ) /_/    \_\  \__,_| |_| |_|  \__,_| |_| |_| 
                        __/ /                                _/ |                 |/                                       |/                                             
                       |___/                                |__/                                                                                                          
               
               
```
# Ascii-Art

- This is a Go program that allow users to draw ASCII Art from the input.
- Users are given these options, including banners, colors, alignment, output which save the generated art to a file and reversing text from a file.
- Users are supposed to input only printable ASCII characters.


## Run Locally

Clone the project

```bash
git clone https://learn.reboot01.com/git/amali/ascii-art-reverse.git
```

Go to the project directory

```bash
  cd ascii-art-reverse
```

**Example how to run the program:**

_Make sure you are in project directory_

```bash
go run .

Usage: go run main.go followed by one of these options

*in the same order*
1- [STRING]
2- [STRING] [BANNER]
3- --color=<color> [STRING]
4- --color=<color> [STRING] [BANNER]
5- --color=<color> <letters to be colored> [STRING]
6- --color=<color> <letters to be colored> [STRING] [BANNER]
*in any order*
7- --output=<fileName.txt> 
8- --align=<type>  ( left right center justify )
**********************************
Can only be activeted alone
9- --reverse=<fileName> 


Those are the available colors:
red, green, yellow, blue, purple, cyan, white, gray, darkred, orange, pink,maram
gold, teal, lavender, brown, lightblue, magenta, olive, salmon, skyblue, darkpurple, lime


```
- For a faster audit, you can use 'testall.sh' to cheak all the examples in the test folder.

#### Usefull linkes
uni code https://web.itu.edu.tr/~sgunduz/courses/mikroisl/ascii.html



#### Creator

 emahfood - amali - ajaberi
