package lem

func Search() [][]string {
	visited[Start] = true
	sol1 := SearchHelper(Ants, false)
	visited = make(map[string]bool)
	visited[Start] = true
	sol2 := SearchHelper(Ants, true)
	choose(sol1, sol2)
	sort1()
	return solutions
	
}

func choose(sol1, sol2 [][]string) {
	if len(sol1) < len(sol2) {
		solutions = sol2
		return
	}
	if len(sol1) > len(sol2) {
		solutions = sol1
		return
	}
	avg1 := avg(sol1)
	avg2 := avg(sol2)
	if avg1 > avg2 {
		solutions = sol2
		return
	}
	solutions = sol1
}

func avg(sli [][]string) int {
	avg := 0
	for _, v := range sli {
		avg += len(v)
	}
	return avg / len(sli)
}

func SearchHelper(ant int, rev bool) [][]string {
	sol := [][]string{}
	if !rev {
		for i := 0; i < len(Ways[Start]); i++ {
			Bfs(Ways[Start][i], &sol, &ant)
			Close(sol)
		}
	} else {
		for i := len(Ways[Start]) - 1; i >= 0; i-- {
			Bfs(Ways[Start][i], &sol, &ant)
			Close(sol)
		}
	}
	return sol
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

func Bfs(s string, solution1 *[][]string, ant *int) bool {
	solution := *solution1
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
			if !visited[neighbour] {
				visited[neighbour] = true
				parent[neighbour] = visiting
				queu = append(queu, neighbour)
			}
			if neighbour == End {
				newPath := findway(parent)
				if len(solution) > 0 {
					if len(solution) > Ants {
						if len(solution[len(solution)-1]) > len(newPath) {
							solution[len(solution)-1] = newPath
							*solution1 = solution
						}
						return false
					} else if len(newPath) < *ant {
						*ant -= len(solution[len(solution)-1])
						solution = append(solution, newPath)
						*solution1 = solution
						return false
					}
				}
				solution = append(solution, newPath)
				*solution1 = solution
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

func Close(sol [][]string) {
	visited = make(map[string]bool)
	for _, s := range sol {
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
