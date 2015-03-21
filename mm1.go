package main

import (
	"fmt"
	"github.com/aybabtme/uniplot/histogram"
	"math"
	"math/rand"
	"os"
	"time"
)

// Queue represents a queue in a M/M/1 Queue simulation
type Queue struct {
	vec []float64
}

// Pop pops the last value off the Queue
// and returns it.
func (q *Queue) Pop() float64 {
	i := len(q.vec) - 1
	elem := q.vec[i]
	q.vec = q.vec[:i]
	return elem
}

// Push appends the value to the queue.
// No value is returned.
func (q *Queue) Push(x float64) {
	q.vec = append(q.vec, x)
}

// Size returns the number of elements in the Queue.
func (q *Queue) Size() int {
	return len(q.vec)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lambda := 1.0
	mu := 1 / 0.6
	queue := Queue{[]float64{}}
	nextArrival := getExpRandNum(lambda)
	nextDeparture := math.Inf(1)
	// var expectedWait = 1.0 / (mu - lambda)
	totalWait := 0.0
	customersServiced := 0
	// var waitTimes = []float64{}
	numPacketsInSystems := []float64{}

	for {
		if nextArrival <= nextDeparture {
			if queue.Size() == 0 {
				nextDeparture = nextArrival + getExpRandNum(mu)
			}
			queue.Push(nextArrival)
			nextArrival += getExpRandNum(lambda)
			numPacketsInSystems = append(numPacketsInSystems, float64(queue.Size()))
		} else {
			wait := nextDeparture - queue.Pop()
			totalWait += wait
			customersServiced++

			if customersServiced == 1000 {
				fmt.Println("Done!")
				fmt.Println("Total wait: ", totalWait)

				hist := histogram.Hist(10, numPacketsInSystems)
				histogram.Fprint(os.Stdout, hist, histogram.Linear(10))
				return
			}
			if queue.Size() == 0 {
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

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func popSlice(a *[]float64) float64 {
	i := len(*a) - 1
	e := (*a)[i]
	*a = append((*a)[:i], (*a)[i+1:]...)
	return e
}
