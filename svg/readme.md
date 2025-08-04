# Floating Point Graphics Computation, SVG

Scalable Vector Graphics is used to plot two variables x, y as a 3-D mesh.
    `z = f(x, y)`

## Assignment
We are showing the output for `function sin(r)/r`, where r is sqrt(x*x+y*y).
## Building

The entry point is orchestrated as a web server that computes surfaces and writes SVG data to client. Therefore, we have the main represent this and the svg file for render functionality.
    `go build main.go svg.go`
## Running
To run enter the following on your terminal:
    `./main`
After enter the localhost with :3000 as port as follows:
    - `http://127.0.0.1:3000?color=blue`
