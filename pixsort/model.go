package pixsort

import "math/rand"

type Point struct {
	X, Y int
}

func (a Point) DistanceTo(b Point) int {
	dx := a.X - b.X
	dy := a.Y - b.Y
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

type Undo struct {
	I, J int
}

type Model struct {
	Points []Point
	Seen   []bool
	Total  int
}

func NewModel(points []Point) *Model {
	model := Model{}
	model.Points = points
	model.Seen = make([]bool, len(points))
	model.Total = 0
	for i := 0; i < len(points)-1; i++ {
		model.Total += points[i].DistanceTo(points[i+1])
	}
	return &model
}

func (m *Model) Energy() float64 {
	return float64(m.Total)
}

func (m *Model) DoMove() interface{} {
	n := len(m.Points)
	i := rand.Intn(n)
	j := rand.Intn(n)
	m.Update(i, j, -1)
	m.Points[i], m.Points[j] = m.Points[j], m.Points[i]
	m.Update(i, j, 1)
	return Undo{i, j}
}

func (m *Model) UndoMove(undo interface{}) {
	u := undo.(Undo)
	i := u.I
	j := u.J
	m.Update(i, j, -1)
	m.Points[i], m.Points[j] = m.Points[j], m.Points[i]
	m.Update(i, j, 1)
}

func (m *Model) Copy() Annealable {
	points := make([]Point, len(m.Points))
	copy(points, m.Points)
	seen := make([]bool, len(m.Seen))
	return &Model{points, seen, m.Total}
}

func (m *Model) Update(i, j, sign int) {
	indexes := []int{i - 1, i, j - 1, j}
	for _, a := range indexes {
		if a < 0 || a >= len(m.Seen)-1 {
			continue
		}
		if m.Seen[a] {
			continue
		}
		m.Seen[a] = true
		m.Total += sign * m.Points[a].DistanceTo(m.Points[a+1])
	}
	for _, a := range indexes {
		if a < 0 || a >= len(m.Seen)-1 {
			continue
		}
		m.Seen[a] = false
	}
}
