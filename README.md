### Snake Game
A console-based version of the beloved [game of Snake](https://www.google.com/search?q=play+snake) using OOP in Java.

## Play the game

The user can move the snake using the following commands:
- `up | right | down | left [n]` changes the snake orientation and moves the snake `n` square in the specified direction. If no parameter is given, the snake is moved 1 square.
- When the snake eats an apple, its tail grows by `1 square` and a new apple is added to the game area.

Game stops when the snake eats itself or hits the wall.

# Settings

You can modify game settings in the `snake.properties` file. You can set the number of rows and columns on the board (`rows`,`columns`),
and the starting coordinates for the snake (`startRow`,`startColumn`).
If not set, the defaults are 6 for rows and columns and 4 for starting row and column.
