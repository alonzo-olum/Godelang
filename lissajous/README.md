# Animated GIFs

Using go's image package, to create a sequence of bit-mapped images and encode the sequence as a GIF animation.
The image forms a lissajous figure ( staple visual effect of the 1960s sci-fi).

## Concept
Oscillating a circle in one revolution equates a 2PI radians (360 degrees).
A lissajous wave forms a parametric curve produced by harmonic oscillations of multiple sine waves. We are going to simulate that.

## What we know and want to apply
Tracing the x components while traversing the circle gives a sine wave, therefore we are going to plot each x coordinate for which a sin(t) exists.
For the y components we are going to plot a sped up x points using a random frequency and using phase to accentuate 3-D effect.

We want roughly five cycles required to complete a resolution.

## Running

Simply `go build main.go lissajous.go`. Then write output to a sample .gif file like one available in same directory.
 Running the main script in the directory starts a web service at `port :3000` on the localhost.
A request is submitted as `http://127.0.0.1:3000?cycles=7`.

## Write GIfs to output file

    `./main >out.gif`

### NB: Run the main file as above otherwise the screen will be used as stdout if you do not provide a .gif file.
