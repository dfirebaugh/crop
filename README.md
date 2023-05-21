# crop
I needed a tool to crop a bunch of png images.  So, this is a very simple cli tool to crop images down.

> note: at the moment this will only crop from the top left corner to a certain width and height -- it also defaults to 128x128 if you don't pass in `-x` and `-y` flags

## Usage

#### Basic usage
```bash
crop

Usage: go run crop.go [-i=<dir>] [-o=<outputDir>] [-x=<pixels to crop>] [-y=<pixels to crop>] <image>
```

You can crop one image:

```bash
crop <image filepath>
```

... or crop an entire directory

```bash
crop -i <path to input directory> -o <path to output directory>
```

```bash
crop -i ./assets/original/ -o ./assets/cropped/ -x 128 -y 128
```
