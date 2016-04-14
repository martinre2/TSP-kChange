package main

import (
	"fmt"
	p "github.com/martinre2/TSP-kChange/Parse"
	s "github.com/martinre2/TSP-kChange/Solve"
	"time"
)

func main() {
	fmt.Println("Prueba TSP K-Change")

	var problems []*p.TSPParser

	parser := new(p.TSPParser)
	parser.ParseFile("./tsp/ftv33.txt")
	problems = append(problems, parser)
	parser2 := new(p.TSPParser)
	parser2.ParseFile("./tsp/ftv33.txt")
	problems = append(problems, parser2)
	parser3 := new(p.TSPParser)
	parser3.ParseFile("./tsp/ftv35.txt")
	problems = append(problems, parser3)
	parser4 := new(p.TSPParser)
	parser4.ParseFile("./tsp/ftv38.txt")
	problems = append(problems, parser4)
	parser5 := new(p.TSPParser)
	parser5.ParseFile("./tsp/ftv44.txt")
	problems = append(problems, parser5)
	parser6 := new(p.TSPParser)
	parser6.ParseFile("./tsp/ftv47.txt")
	problems = append(problems, parser6)
	parser7 := new(p.TSPParser)
	parser7.ParseFile("./tsp/ftv55.txt")
	problems = append(problems, parser7)
	parser8 := new(p.TSPParser)
	parser8.ParseFile("./tsp/ftv64.txt")
	problems = append(problems, parser8)
	parser9 := new(p.TSPParser)
	parser9.ParseFile("./tsp/ftv70.txt")
	problems = append(problems, parser9)

	solver := s.NewSolver(problems, 30, 3)
	var best, worts float32
	fmt.Println("Problema", "\t", "Nombre", "\t", "Ciudades", "\t", "Iteraciones", "\t", "Tiempo", "\t", "# Best", "\t", "Best", "\t", "Worts", "\t", "Optimo")
	for i, p := range problems {
		start_process := time.Now()

		tour := solver.RandTour(i)
		best = 9999999999
		worts = 0.0
		best_find := 0
		for it := solver.MaxIter; it > 0; it-- {
			weight := solver.CalcWeights(i, tour)

			if weight < best {
				best = weight
			}
			if weight > worts {
				worts = weight
			}
			if weight == float32(p.Data.BestKnow) {
				best_find++
			}

			tour = solver.Change(tour, i)
		}
		end_process := time.Now()
		fmt.Println(i, "\t", p.Data.Name, "\t", p.Data.Dimension, "\t", solver.MaxIter, "\t", end_process.Sub(start_process), "\t", best_find, "\t", best, "\t", worts, "\t", p.Data.BestKnow)
	}

}
