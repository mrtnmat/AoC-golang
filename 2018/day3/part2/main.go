package main

import (
	"fmt"
	"strings"
)

const (
	fbsize int = 1024
)

type (
	fabric [fbsize][fbsize]rune
	coor   struct{ x, y int }
	size   struct{ x, y int }
	patch  struct {
    id      int
		coor
		size
    overlap bool
	}
	state struct {
		fb      fabric
		overlap int
	}
)

func main() {
	var st state
	var fb fabric
	var lines []string
  var patches []patch

	for i, _ := range fb {
		for j, _ := range fb[i] {
			fb[i][j] = '.'
		}
	}

	st = state{fb, 0}

	lines = strings.Split(input, "\n")
	for i, _ := range lines {
    patches = append(patches, patch{})
		words := strings.Split(lines[i], " ")
		fmt.Sscanf(words[0], "#%d", &patches[i].id)
		fmt.Sscanf(words[2], "%d,%d:", &patches[i].coor.x, &patches[i].coor.y)
		fmt.Sscanf(words[3], "%dx%d", &patches[i].size.x, &patches[i].size.y)
		st.claim_patch(patches[i])
	}

	//st.fb.print_all()

	fmt.Printf("total overlap: %v\n", st.overlap)
}

func (st *state) claim_patch(ptc patch) {
	var ySlice []rune
	var xSlice [][fbsize]rune

	xSlice = st.fb[ptc.coor.x : ptc.coor.x+ptc.size.x]
	for i, _ := range xSlice {
		ySlice = xSlice[i][ptc.coor.y : ptc.coor.y+ptc.size.y]
		for j, _ := range ySlice {
			switch ySlice[j] {
			case '.':
				ySlice[j] = '/'
			case '/':
				ySlice[j] = '#'
				st.overlap++
        ptc.overlap = true
			}
		}
	}
}

func (fb *fabric) print_all() {
	fb.print(coor{0, 0}, size{fbsize, fbsize})
}

func (fb *fabric) print(c coor, s size) {
	var ySlice []rune
	var xSlice [][fbsize]rune

	xSlice = fb[c.x : c.x+s.x]
	for i, _ := range xSlice {
		ySlice = xSlice[i][c.y : c.y+s.y]
		fmt.Printf("%s\n", string(ySlice))
	}
}
