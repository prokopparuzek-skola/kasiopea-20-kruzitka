package main

import "fmt"

func bfs(graph [][]int) {
	var ruler bool
	var Aqueue, Fqueue []int
	var bebef []bool
	var buyR []int
	var test bool = true

	Aqueue = make([]int, 0)
	Fqueue = make([]int, 0)
	bebef = make([]bool, len(graph))
	buyR = make([]int, 0)
	Aqueue = append(Aqueue, 0)
	for len(Aqueue) > 0 {
		for _, v := range Aqueue {
			for _, a := range graph[v] {
				if !bebef[a] {
					Fqueue = append(Fqueue, a)
					bebef[a] = true
					if ruler {
						buyR = append(buyR, a)
					}
				}
			}
		}
		Aqueue = Fqueue
		Fqueue = make([]int, 0)
		ruler = !ruler
	}
	for _, v := range bebef {
		if !v {
			test = false
			break
		}
	}
	if test {
		fmt.Println("ANO")
		fmt.Println(len(buyR))
		for _, v := range buyR {
			fmt.Println(v + 1)
		}
	} else {
		fmt.Println("NE")
	}
}

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var N, M int
		var graph [][]int
		fmt.Scan(&N, &M)
		graph = make([][]int, N)
		for j := range graph {
			graph[j] = make([]int, 0)
		}
		for j := 0; j < M; j++ {
			var u, v int
			fmt.Scan(&u, &v)
			u--
			v--
			graph[u] = append(graph[u], v)
			graph[v] = append(graph[v], u)
		}
		bfs(graph)
	}
}
