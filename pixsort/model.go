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
}

func NewModel(points []Point) *Model {
	model := Model{}
	model.Points = points
	return &model
}

func (m *Model) Energy() float64 {
	total := 0.0
	for i := 0; i < len(m.Points)-1; i++ {
		total += float64(m.Points[i].DistanceTo(m.Points[i+1]))
	}
	return total
}

func (m *Model) DoMove() interface{} {
	n := len(m.Points)
	i := rand.Intn(n)
	j := rand.Intn(n)
	points := m.Points
	p := points[i]
	points = append(points[:i], points[i+1:]...)
	points = append(points[:j], append([]Point{p}, points[j:]...)...)
	return Undo{i, j}
}

func (m *Model) UndoMove(undo interface{}) {
	u := undo.(Undo)
	j := u.I
	i := u.J
	points := m.Points
	p := points[i]
	points = append(points[:i], points[i+1:]...)
	points = append(points[:j], append([]Point{p}, points[j:]...)...)
}

func (m *Model) Copy() Annealable {
	points := make([]Point, len(m.Points))
	copy(points, m.Points)
	return &Model{points}
}
