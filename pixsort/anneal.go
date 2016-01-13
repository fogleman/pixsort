package pixsort

import (
	"fmt"
	"math"
	"math/rand"
)

type Annealable interface {
	Energy() float64
	DoMove() interface{}
	UndoMove(interface{})
	Copy() Annealable
}

func Anneal(state Annealable, maxTemp, minTemp float64, steps int) Annealable {
	factor := -math.Log(maxTemp / minTemp)
	state = state.Copy()
	bestState := state.Copy()
	bestEnergy := state.Energy()
	previousEnergy := bestEnergy
	for step := 0; step < steps; step++ {
		if step%100000 == 0 {
			showProgress(step, steps)
		}
		pct := float64(step) / float64(steps-1)
		temp := maxTemp * math.Exp(factor*pct)
		undo := state.DoMove()
		energy := state.Energy()
		change := energy - previousEnergy
		if change > 0 && math.Exp(-change/temp) < rand.Float64() {
			state.UndoMove(undo)
		} else {
			previousEnergy = energy
			if energy < bestEnergy {
				// pct := float64(step*100) / float64(steps)
				// fmt.Printf("step: %d of %d (%.1f%%), temp: %.3f, energy: %.3f\n",
				//     step, steps, pct, temp, energy)
				bestEnergy = energy
				bestState = state.Copy()
			}
		}
	}
	showProgress(steps, steps)
	return bestState
}

func PreAnneal(state Annealable, iterations int) float64 {
	state = state.Copy()
	previous := state.Energy()
	var total float64
	for i := 0; i < iterations; i++ {
		state.DoMove()
		energy := state.Energy()
		total += math.Abs(energy - previous)
		previous = energy
	}
	return total / float64(iterations)
}

func showProgress(i, n int) {
	pct := int(100 * float64(i) / float64(n))
	fmt.Printf("  %3d%% [", pct)
	for p := 0; p < 100; p += 3 {
		if pct > p {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("]    \r")
}
