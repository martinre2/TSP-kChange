package Solve

import (
	_ "fmt"
	p "github.com/martinre2/TSP-kChange/Parse"
	"math/rand"
	"time"
)

type Solver struct {
	Problems   []*p.TSPParser
	MaxIter    int
	TypeChange int
}

func NewSolver(problems []*p.TSPParser, maxiter int, typeChange int) *Solver {
	solver := new(Solver)
	solver.Problems = problems
	solver.MaxIter = maxiter
	solver.TypeChange = typeChange
	return solver
}

func (s *Solver) RandTour(problem_indx int) []int {
	tour := make([]int, s.Problems[problem_indx].Data.Size)
	for i := 0; i < s.Problems[problem_indx].Data.Size; i++ {
		tour[i] = i
	}

	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	for i := len(tour) - 1; i > 0; i-- {
		j := rand.Intn(i)
		tour[i], tour[j] = tour[j], tour[i]
	}
	return tour
}

func (s *Solver) CalcWeights(problem_indx int, tour []int) float32 {
	var total_weights float32
	total_weights = 0.0
	for i, e := range tour {
		if i < s.Problems[problem_indx].Data.Size-1 {
			total_weights += s.Problems[problem_indx].Data.Weights[e][tour[i+1]]
		} else {
			total_weights += s.Problems[problem_indx].Data.Weights[e][tour[0]]
		}
	}
	return total_weights
}

func (s *Solver) Change(tour []int, problem_indx int) []int {
	var rs_tour []int
	switch s.TypeChange {
	case 0:
		rs_tour = s.RandTour(problem_indx)
		break
	case 2:
		rs_tour = s.TwoChange(tour)
		break
	case 3:
		rs_tour = s.ThreeChange(tour)
		break
	}
	return rs_tour
}

func (s *Solver) TwoChange(tour []int) []int {
	//fmt.Println(tour)
	changeTour := make([]int, len(tour))
	var i, k int
	for {
		t := time.Now()
		rand.Seed(int64(t.Nanosecond()))
		i = rand.Intn(len(tour))
		k = rand.Intn(len(tour))
		if i != k &&
			i < k &&
			(k-i) > 1 {
			break
		}
	}
	//fmt.Println(i, k)

	copy(changeTour[0:i], tour[0:i])
	revTour := reverseInts(tour[i:k])
	copy(changeTour[i:k], revTour)
	copy(changeTour[k:len(changeTour)], tour[k:len(tour)])
	//fmt.Println(changeTour)
	return changeTour
}

func (s *Solver) ThreeChange(tour []int) []int {
	changeTour := make([]int, len(tour))
	var i, k, j int
	for {
		t := time.Now()
		rand.Seed(int64(t.Nanosecond()))
		i = rand.Intn(len(tour))
		k = rand.Intn(len(tour))
		j = rand.Intn(len(tour))
		if (i != k) && (k != j) && (i < k) && (k < j) {
			break
		}
	}

	copy(changeTour[0:i], tour[0:i])
	revTour1 := reverseInts(tour[i:k])
	copy(changeTour[i:k], revTour1)
	revTour2 := reverseInts(tour[k:j])
	copy(changeTour[k:j], revTour2)
	copy(changeTour[j:len(changeTour)], tour[j:len(tour)])

	return changeTour
}

func reverseInts(parray []int) []int {
	input := make([]int, len(parray))
	copy(input, parray)
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
