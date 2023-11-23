package life

import (
	"math/rand"
	"time"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) (*World, error) {
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}, nil
}

func (w *World) Next(x, y int) bool {
	n := w.Neighbours(x, y)
	alive := w.Cells[y][x]
	if n < 4 && n > 1 && alive {
		return true
	}
	if n == 3 && !alive {
		return true
	}

	return false
}

func (w *World) Neighbours(x, y int) int {
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	count := 0
	for i := 0; i < 8; i++ {
		newX := x + dx[i]
		newY := y + dy[i]
		if newX >= 0 && newX < w.Width && newY >= 0 && newY < w.Height {
			if w.Cells[newY][newX] {
				count++
			}
		}
	}
	return count
}

func NextState(oldWorld, newWorld *World) {
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			newWorld.Cells[i][j] = oldWorld.Next(j, i)
		}
	}
}

func (w *World) RandInit(percentage int) {
	numAlive := percentage * w.Height * w.Width / 100
	w.fillAlive(numAlive)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < w.Height*w.Width; i++ {
		randRowLeft := r.Intn(w.Width)
		randColLeft := r.Intn(w.Height)
		randRowRight := r.Intn(w.Width)
		randColRight := r.Intn(w.Height)

		w.Cells[randRowLeft][randColLeft] = w.Cells[randRowRight][randColRight]
	}
}

func (w *World) fillAlive(num int) {
	aliveCount := 0
	for j, row := range w.Cells {
		for k := range row {
			w.Cells[j][k] = true
			aliveCount++
			if aliveCount == num {
				return
			}
		}
	}
}
