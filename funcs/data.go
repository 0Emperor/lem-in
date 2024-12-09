package lem

var (
	Ways      = make(map[string][]string)
	Emptyroom = make(map[string]bool)
	rating    = make(map[int]int)
	Rooms     = []string{}
	Start, End    string
	Ants          int
	Graphoverview []byte
	seen =make(map[string]bool)
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
