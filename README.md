# Terminal Space Invaders

![project demo](./demo.gif)

This project is a low-fidelity implementation of Space Invaders, playable directly in the terminal!
The primary purpose of this was for me to learn `Golang` and explore how games may implement an event loop and player input.
It also features dynamic resolution scaling where images are
resized then transformed into ascii characters, which are then drawn to the screen.

### Dynamic resolution scaling

![resolution scaling](./dynamic-resolution-demo.gif)

## Running the game

This program depends on `go` and `cmake`.

Clone this repo

```
git clone git@github.com:spoonk/space-invaders.git
```

Then run the project with:

```
make run
```
