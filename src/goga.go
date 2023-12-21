package main

import (
//    "C"
//    "encoding/json"
//    "encoding/csv"
//    "os"
   "fmt"
//    "strconv"
	"math/rand"
)

type PopType struct {
	P []int
	F int
}

const NUM_GENES = 4		// Even numbers work more evenly
const MAX_GENE = 99
const POP_SIZE = 100
const ITERATIONS = 1000

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
	// Sort pop based on fit_res
	for j:=0;j<len(fit_res);j++{
		key := fit_res[j]
		key_pop := pop[j]
		i := j-1
		for (i>0 && fit_res[i]>key) {
			fit_res[i+1] = fit_res[i]
			pop[i+1] = pop[i]
			i = i-1
		}
		fit_res[i+1] = key
		pop[i+1] = key_pop
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

	min_res := 999999999999.0
	best_pop := pop[0]
	
	for i:=0;i<ITERATIONS;i++{
		// Selection
		sorted_pop := sort_pop(pop, fit_res)
		// fmt.Println(sorted_pop)

		// Elitism
		var new_pop [][]int
		for i:=len(fit_res)-1;i>=len(fit_res)-10;i--{
			new_pop = append(new_pop, sorted_pop[i])
		}
		// fmt.Println(new_pop)

		// Crossover
		for i:=0;i<90;i++{
			parent1 := pop[50+rand.Intn(POP_SIZE/2)]
			parent2 := pop[50+rand.Intn(POP_SIZE/2)]
			new_pop = append(new_pop, mate(parent1, parent2))
		}

		// Calculate fitness
		pop = new_pop
		fit_res = get_fit(pop)
		fmt.Println("Iteration:", i, "| Best pop:", pop[0], "| Fitness:", -fit_res[0], "hours")
		if -fit_res[0] > min_res {
			best_pop = pop[0]
			min_res = -fit_res[0]
		}
		// Negatives replaced if maximising
	}

	fmt.Println(pop[len(pop)-1])
	fmt.Println(fit_res[len(pop)-1])
		
	// Final Selection

	fmt.Println("DONE")
	fmt.Println("-------------------")
	fmt.Println("Best pop:", best_pop, "| Fitness:", min_res, "hours")
	fmt.Println("-------------------")
}
