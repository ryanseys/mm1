package main

import "fmt"
import "math"
import "math/rand"

// Number of iterations
const ITERATIONS = 10

func main() {
	var lambda = 1.0
	var mu = 1 / 0.6
	var queue = []float64{}
	var nextArrival = getExpRandNum(lambda)
	var nextDeparture = math.Inf(1)
	// var expectedWait = 1.0 / (mu - lambda)
	var totalWait = 0.0
	var customersServiced = 0
	// var waitTimes = []float64{}
	var numPacketsInSystems = []int{}

	for {
		if nextArrival <= nextDeparture {
			if len(queue) == 0 {
				nextDeparture = nextArrival + getExpRandNum(mu)
			}
			queue = append(queue, nextArrival)
			nextArrival += getExpRandNum(lambda)
			numPacketsInSystems = append(numPacketsInSystems, len(queue))
		} else {
			wait := nextDeparture - popSlice(&queue)
			totalWait += wait
			customersServiced++

			if customersServiced == 1000 {
				fmt.Println("Done!")
				fmt.Println("Total wait: ", totalWait)
				return
			}
			if len(queue) == 0 {
				nextDeparture = math.Inf(1)
			} else {
				nextDeparture += getExpRandNum(mu)
			}
		}
	}
}

func getExpRandNum(lambda float64) float64 {
	return math.Log(1-rand.Float64()) / (-lambda)
}

func getUniformRanNum() float64 {
	return rand.Float64()
}

func getUniformRandBetween(min, max int) float64 {
	return rand.Float64()*float64(max-min) + float64(min)
}

func getUniformRandIntBetween(min, max int) int {
	return int(math.Floor(rand.Float64()*float64(max-min))) + min
}

func popSlice(a *[]float64) float64 {
	i := len(*a) - 1
	e := (*a)[i]
	*a = append((*a)[:i], (*a)[i+1:]...)
	return e
}
