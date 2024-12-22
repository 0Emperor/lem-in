package lem

var (
	Ways       = make(map[string][]string)
	Emptyroom  = make(map[string]bool)
	rating     = make(map[int]int)
	Rooms      = []string{}
	Start, End string
	Ants       int
	seen       = make(map[string]bool)
	max        = 0
	badrooms   = []string{}
)

type ppp struct {
	rating int
	index  int
}

var (
	visited   = make(map[string]bool)
	solutions [][]string
	stack     []string
)
