package math

//symlinked into GOPATH via something like:
//ln -s /Users/se/workspace/go-scratchpad/ $GOPATH/src
//note: always use absolute paths with symlinks

//Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
} //capital letters = public for other packages/programs
