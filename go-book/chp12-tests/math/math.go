package math

//symlinked into GOPATH via something like:
//ln -s $GOPATH/src/go-book/chp11-packages/math ~/workspace/go-scratchpad/go-book/chp11-packages/math/

//Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
} //capital letters = public for other packages/programs
