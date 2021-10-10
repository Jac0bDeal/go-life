# go-life

This project aims to implement a simple Conway's Game of Life sim in Golang. The current implementation simply prints 
the sim in real time to the terminal until the user kills the process. 

Additional features such as saving runs as gifs, saving/loading initial world states, and concurrency to allow for larger
worlds to run with a decent render time.

## Building
In order to build this, you need Go 1.17+ and Make installed.

From the project root, run
```shell
make bin
```

## Running
Once built, run the binary with
```shell
./bin/go-life
```

The simulation will begin running, printing to the terminal session used to launch it. You may need to resize the
terminal session to fit the size of the world.

To stop the simulation, simply kill the process using `Ctrl+X` or killing in the task manager.
