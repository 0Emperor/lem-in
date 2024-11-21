package lem

func Search() [][]string {
	Dfs(Start)
	sort(solutions)
	set()
	rate()
	return combine(gettrials())
}

func set() {
	for v := range solutions {
		rating[v] = len(solutions[v])
	}
}

var rating = make(map[int]int)

func cop(p1, p2 []string) bool {
	for i, v := range p1 {
		if p2[i] != v {
			return false
		}
	}
	return true
}

func hhhh(paths ...[]string) bool {
	for _, path := range paths {
		for _, path0 := range paths {
			if cc(path0, path) {
				return true
			}
		}
	}
	return false
}

func cc(p1, p2 []string) bool {
	for i := 1; i < len(p1)-1; i++ {
		v := p1[i]
		for j := 1; j < len(p2)-1; j++ {
			r := p2[j]
			if !cop(p1, p2) && v == r {
				return true
			}
		}
	}
	return false
}

func rate() {
	for i, v := range solutions {
		for t, r := range solutions {
			if i != t && cc(r, v) {
				rating[i]++
			}
		}
	}
}

type ppp struct {
	rating int
	index  int
}

func sortby(ss []ppp) []ppp {
	for i := 0; i < len(ss)-1; i++ {
		for j := i + 1; j < len(ss); j++ {
			if ss[i].rating > ss[j].rating {
				ss[i], ss[j] = ss[j], ss[i]
			}
		}
	}
	return ss
}

func re() []ppp {
	ss := []ppp{}
	for i, v := range rating {
		s := ppp{
			rating: v,
			index:  i,
		}
		ss = append(ss, s)
	}
	return sortby(ss)
}

func combine(n int) [][]string {
	sss := [][]string{}
	rrr := re()
	max := 0
	for u := 0; u < len(rrr); u++ {
		e := rrr[u]
		i := e.index
		if len(solutions[i]) > max {
			max = i
		}
		supp := [][]string{}
		supp = append(sss, solutions[i])
		if !hhhh(supp...) {
			sss = append(sss, solutions[i])
		} else if len(solutions[max]) > 2*len(solutions[i]) {
			supp = [][]string{}
			supp = append(sss, solutions[i])
			max = i
			u = 0
		}
		sort(sss)
		if len(sss) == n {
			break
		}
	}
	return sss
}

func sort(unsorted [][]string) [][]string {
	for i := 0; i < len(unsorted); i++ {
		for j := i + 1; j < len(unsorted); j++ {
			if len(unsorted[i]) >= len(unsorted[j]) {
				unsorted[i], unsorted[j] = unsorted[j], unsorted[i]
			}
		}
	}
	return unsorted
}

func gettrials() int {
	fromstart := len(Ways[Start])
	fromend := len(Ways[End])
	if fromend <= fromstart {
		return fromend
	} else {
		return fromstart
	}
}

var (
	visited   = make(map[string]bool)
	solutions [][]string
	stack     []string
)

func Dfs(current string) {
	stack = append(stack, current)

	if visited[current] { //|| (len(Ways[current]) == 1 && current != End && current != Start)
		stack = stack[:len(stack)-1]
		return
	}

	if current == End {
		supp := []string{}
		supp = append(supp, stack...)
		solutions = append(solutions, supp)
		stack = stack[:len(stack)-1]
		return
	}
	visited[current] = true
	for _, v := range Ways[current] {
		Dfs(v)
	}
	stack = stack[:len(stack)-1]
	visited[current] = false
}
