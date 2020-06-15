# Game of Life in Golang

## Design decisions
While a common approach involves arrays of arrays, we wanted to try a flat approach.  A board contains a slice of cell objects each of which contains a slice of pointers to each neighbor cell.

## Testing
We used the blinker oscillator for most of our testing
010      
010  
010  

=>

000
111
000

## Next steps
glider is buggy
