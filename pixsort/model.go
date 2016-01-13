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
	Score  int
}

func NewModel(points []Point) *Model {
	model := Model{}
	model.Points = points
	model.Seen = make([]bool, len(points))
	model.Score = 0
	for i := 0; i < len(points)-1; i++ {
		model.Score += points[i].DistanceTo(points[i+1])
	}
	return &model
}

func (m *Model) Energy() float64 {
	return float64(m.Score)
}

func (m *Model) DoMove() interface{} {
	n := len(m.Points)
	i := rand.Intn(n)
	j := rand.Intn(n)
	m.Update(i, j, -1)
	m.Move(i, j)
	m.Update(i, j, 1)
	return Undo{i, j}
}

func (m *Model) UndoMove(undo interface{}) {
	u := undo.(Undo)
	m.Update(u.J, u.I, -1)
	m.Move(u.J, u.I)
	m.Update(u.J, u.I, 1)
}

func (m *Model) Copy() Annealable {
	points := make([]Point, len(m.Points))
	copy(points, m.Points)
	return &Model{points, m.Seen, m.Score}
}

func (m *Model) Move(i, j int) {
	p := m.Points
	v := p[i]
	for i < j {
		p[i] = p[i+1]
		i++
	}
	for i > j {
		p[i] = p[i-1]
		i--
	}
	p[j] = v
}

func (m *Model) Update(i, j, sign int) {
	var indexes []int
	if i == j {
		return
	} else if sign < 0 && i < j {
		indexes = []int{i - 1, i, j}
	} else if sign > 0 && i < j {
		indexes = []int{i - 1, j - 1, j}
	} else if sign < 0 && i > j {
		indexes = []int{i - 1, i, j - 1}
	} else if sign > 0 && i > j {
		indexes = []int{i, j - 1, j}
	}
	for _, a := range indexes {
		if a < 0 || a >= len(m.Seen)-1 || m.Seen[a] {
			continue
		}
		m.Seen[a] = true
		m.Score += sign * m.Points[a].DistanceTo(m.Points[a+1])
	}
	for _, a := range indexes {
		if a < 0 || a >= len(m.Seen)-1 {
			continue
		}
		m.Seen[a] = false
	}
}
