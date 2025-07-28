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

Simply `go build lissajous`. Then write output to a sample .gif file like one available in same directory.

## Write GIfs to output file

    `./lissajous >out.gif`
