## Golang implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life)

This is my first non trivial Golang program, implementing 1d and 2d Conway's Game of Life.

It makes use of Go Modules and runs under Go 1.17+ (at least).

### 1D

This "unofficial" version of Conway's GoF was implemented as a first step in this project and follows (by default) the following rules:

#### State diagram

<?xml version="1.0" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "https://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">

<svg width="800" height="600" version="1.1" xmlns="http://www.w3.org/2000/svg">
	<ellipse stroke="black" stroke-width="1" fill="none" cx="237.5" cy="267.5" rx="30" ry="30"/>
	<text x="212.5" y="273.5" font-family="Times New Roman" font-size="20">Alive</text>
	<ellipse stroke="black" stroke-width="1" fill="none" cx="401.5" cy="267.5" rx="30" ry="30"/>
	<text x="375.5" y="273.5" font-family="Times New Roman" font-size="20">Dead</text>
	<path stroke="black" stroke-width="1" fill="none" d="M 210.703,280.725 A 22.5,22.5 0 1 1 210.703,254.275"/>
	<text x="105.5" y="273.5" font-family="Times New Roman" font-size="20">N = 1</text>
	<polygon fill="black" stroke-width="1" points="210.703,254.275 207.17,245.527 201.292,253.618"/>
	<path stroke="black" stroke-width="1" fill="none" d="M 428.297,254.275 A 22.5,22.5 0 1 1 428.297,280.725"/>
	<text x="474.5" y="273.5" font-family="Times New Roman" font-size="20">N != 1</text>
	<polygon fill="black" stroke-width="1" points="428.297,280.725 431.83,289.473 437.708,281.382"/>
	<path stroke="black" stroke-width="1" fill="none" d="M 374.945,281.35 A 149.361,149.361 0 0 1 264.055,281.35"/>
	<polygon fill="black" stroke-width="1" points="374.945,281.35 365.661,279.677 369.373,288.962"/>
	<text x="285.5" y="313.5" font-family="Times New Roman" font-size="20">N != 1</text>
	<path stroke="black" stroke-width="1" fill="none" d="M 265.222,256.125 A 180.118,180.118 0 0 1 373.778,256.125"/>
	<polygon fill="black" stroke-width="1" points="265.222,256.125 274.357,258.481 271.344,248.946"/>
	<text x="289.5" y="238.5" font-family="Times New Roman" font-size="20">N = 1</text>
</svg>
