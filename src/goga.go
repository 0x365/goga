package main

import (
//    "C"
//    "encoding/json"
   "encoding/csv"
   "os"
   "fmt"
   "strconv"
	"math/rand"
)

type PopType struct {
	P []int
	F int
}

const NUM_GENES = 4		// Even numbers work more evenly
const MAX_GENE = 81
const POP_SIZE = 100
const ITERATIONS = 100
const PAR_RATIO = 0.4	// Ratio of sorted population to use as parents for next pop
const ELITISM = 0.1

const SAVE_FILE = "data/ga_results.csv"

func fitness(genetics []int) (summer float64) {
	for i:=0;i<len(genetics);i++{
		summer += float64(genetics[i])
	}
	return summer
}

func initial_pop() (pop [][]int) {
	for j:=0;j<POP_SIZE;j++{
		var genetics []int
		for i:=0;i<NUM_GENES;i++{
			genetics = append(genetics, rand.Intn(MAX_GENE+1))
		}
		pop = append(pop, genetics)
	}
	return pop
}

func get_fit(pop [][]int) (fit_res []float64) {
	for i:=0;i<POP_SIZE;i++{
		fit_res = append(fit_res, fitness(pop[i]))
	}
	return fit_res
}

func sort_pop(pop [][]int, fit_res []float64) ([][]int) {
	// Bubble sort (WIP --- Could be improved for speed)
	n := POP_SIZE
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if fit_res[j] < fit_res[j+1] {
				fit_res[j], fit_res[j+1] = fit_res[j+1], fit_res[j]
				pop[j], pop[j+1] = pop[j+1], pop[j]
			}
		}
	}
	return pop
}

func mate(genes1, genes2 []int) (child_genes []int) {
	for i:=0;i<NUM_GENES;i++{
		prob := rand.Intn(100)
		if prob < 45 {
			child_genes = append(child_genes, genes1[i])
		} else if prob < 90 {
			child_genes = append(child_genes, genes2[i])
		} else {
			child_genes = append(child_genes, rand.Intn(MAX_GENE+1))
		}
	}
	return child_genes
}

func main() {
	fmt.Println("Start")

	// Generate initial population
	pop := initial_pop()
	// Get initial fitness
	fit_res := get_fit(pop)

	// Set initial top results
	min_res := fit_res[0]
	best_pop := pop[0]

	var pop_timeline [][]int
	var fit_timeline []float64
	
	for i:=0;i<ITERATIONS;i++{
		// Selection
		sorted_pop := sort_pop(pop, fit_res)		

		// Elitism
		var new_pop [][]int
		new_pop = append(new_pop, sorted_pop[:(POP_SIZE*ELITISM)]...)

		// Crossover
		for i:=0;i<POP_SIZE*(1-ELITISM);i++{
			parent1 := sorted_pop[POP_SIZE*(1-PAR_RATIO)+rand.Intn(POP_SIZE*PAR_RATIO)]
			parent2 := sorted_pop[POP_SIZE*(1-PAR_RATIO)+rand.Intn(POP_SIZE*PAR_RATIO)]
			new_pop = append(new_pop, mate(parent1, parent2))
		}

		// Calculate fitness
		pop = new_pop
		fit_res = get_fit(pop)

		// Display results for this population
		fmt.Println("Iteration:", i, "| Best pop:", pop[0], "| Fitness:", fit_res[0], "hours")
		if fit_res[0] > min_res {
			best_pop = pop[0]
			min_res = fit_res[0]
		}

		pop_timeline = append(pop_timeline, pop[0])
		fit_timeline = append(fit_timeline, fit_res[0])
	}

	save_timelines(pop_timeline, fit_timeline)

	fmt.Println("DONE")
	fmt.Println("-------------------")
	fmt.Println("Best pop:", best_pop, "| Fitness:", min_res, "hours")
	fmt.Println("-------------------")
}



func save_timelines(pop_timeline [][]int, fit_timeline []float64) {
	file, _ := os.Create(SAVE_FILE)
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()
	for i:=0;i<len(pop_timeline);i++{
		var temp []string
		temp = append(temp,  fmt.Sprintf("%f",fit_timeline[i]))
		for j:=0;j<len(pop_timeline[i]);j++{
			temp = append(temp, strconv.Itoa(pop_timeline[i][j]))
		}
		w.Write(temp)
	}
}