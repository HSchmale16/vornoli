# Vornoli Generator

Generates a Vornoli diagram given an input file structed like so:

![Example Diagram](example.png)

Did it with linear algebra cause that's neat.

## Input Format

```
width height numPoints
x1 y1
x2 y2
...
xN yN
```

# Run it and test it

```shell
go build ; time ./vornoli < test.txt > t.ppm ; display t.ppm
```

Requires golang to be installed, and ImageMagick for display command otherwise will work on any system.