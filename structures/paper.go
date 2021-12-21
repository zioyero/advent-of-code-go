package structures

type Paper struct {
	dots  []Position
	paper map[Position]bool
}

type Position struct {
	x int
	y int
}

type Fold struct {
	X   bool
	Y   bool
	Val int
}

func NewPosition(x int, y int) Position {
	return Position{x, y}
}

func NewPaper() Paper {
	return Paper{
		dots:  make([]Position, 0),
		paper: make(map[Position]bool),
	}
}

func (p *Paper) Dots() []Position {
	return p.dots
}

func (p *Paper) AddDots(positions []Position) {
	for _, pos := range positions {
		p.AddDot(pos)
	}
}

func (p *Paper) AddDot(pos Position) {
	dot, ok := p.paper[pos]
	if ok && dot {
		// dot already on the paper
		return
	}
	p.dots = append(p.dots, pos)
	p.paper[pos] = true
}

func (p *Paper) DotCount() int {
	return len(p.dots)
}

func (p *Paper) Size() (int, int) {
	maxX := 0
	maxY := 0
	for _, dot := range p.dots {
		if dot.x > maxX {
			maxX = dot.x
		}
		if dot.y > maxY {
			maxY = dot.y
		}
	}
	return maxX, maxY
}

func (p *Paper) Fold(fold Fold) Paper {
	if fold.X {
		return FoldX(*p, fold.Val)
	} else {
		return FoldY(*p, fold.Val)
	}
}

func FoldX(p Paper, x int) Paper {
	folded := NewPaper()
	for _, dot := range p.dots {
		if dot.x > x {
			folded.AddDot(Position{x: x - (dot.x - x), y: dot.y})
		} else if dot.x == x {
			// drop
		} else {
			folded.AddDot(dot)
		}
	}
	return folded
}

func FoldY(p Paper, y int) Paper {
	folded := NewPaper()
	for _, dot := range p.dots {
		if dot.y > y {
			folded.AddDot(Position{x: dot.x, y: y - (dot.y - y)})
		} else if dot.y == y {
			// drop
		} else {
			folded.AddDot(dot)
		}
	}
	return folded
}

func (p *Paper) String() string {
	sizeX, sizeY := p.Size()
	str := "" // fmt.Sprintf("Dots: %v\n", p.dots)
	for y := 0; y <= sizeY; y++ {
		for x := 0; x <= sizeX; x++ {
			if p.paper[Position{x, y}] {
				str += "#"
			} else {
				str += "."
			}
		}
		str += "\n"
	}
	return str
}
