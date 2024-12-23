package lem

func Search() [][]string {
	visited[Start] = true
	if contains(Ways[Start]) {
		solutions = append(solutions, []string{Start, End})
	}
	for i := 0; i < len(Ways[Start])-1; i++ {
		if len(solutions) == Ants-1 {
			break
		}
		Bfs(Start, true)
	}
	Close()
	for i := 0; i < len(Ways[Start]); i++ {

		Bfs(Start, false)

		if len(solutions) == Ants {
			break
		}
	}
	sort1()
	return solutions
}

func contains(s []string) bool {
	for i, v := range s {
		if v == End {
			Ways[Start] = append(Ways[Start][:i], Ways[Start][i+1:]...)
			return true
		}
	}
	return false
}

func compare(s1, s2 []string) bool {
	if len(s1) != len(s2) || len(s1) == 0 || len(s2) == 0 {
		return false
	}

	for i, v := range s1 {
		if s2[i] != v {
			return false
		}
	}
	return true
}

func sort1() {
	for i := 0; i < len(solutions)-1; i++ {
		for j := i + 1; j < len(solutions); j++ {
			if len(solutions[j]) < len(solutions[i]) {
				solutions[i], solutions[j] = solutions[j], solutions[i]
			}
			if compare(solutions[j][1:len(solutions[i])-1], solutions[i][1:len(solutions[i])-1]) {
				solutions = append(solutions[:i], solutions[i+1:]...)
			}
		}
	}
}

func Bfs(s string, ignor bool) bool {
	parent := make(map[string]string)
	parent[s] = Start
	if s == End {

		solutions = append(solutions, findway(parent))
		return false
	}

	visited[Start] = true
	queu := []string{s}
	visited[s] = true

	for i := 0; i < len(queu); i++ {
		visiting := queu[i]
		for _, neighbour := range Ways[visiting] {
			if ccc(neighbour) && ignor {
				continue
			}
			if !visited[neighbour] {
				visited[neighbour] = true
				parent[neighbour] = visiting
				queu = append(queu, neighbour)
			}
			if neighbour == End {
				solutions = append(solutions, findway(parent))
				Close()
				return false

			}
		}
	}
	return true
}

func ccc(s string) bool {
	for _, v := range badrooms {
		if v == s {
			return true
		}
	}
	return false
}

func findway(parent map[string]string) []string {
	curent := End
	visited = make(map[string]bool)
	way := []string{curent}
	for curent != Start {
		way = append(way, parent[curent])
		curent = parent[curent]
	}
	return flip(way)
}

func Close() {
	visited = make(map[string]bool)
	for _, s := range solutions {
		for _, v := range s[1 : len(s)-1] {
			visited[v] = true
		}
	}
}

func flip(s []string) []string {
	r := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}

