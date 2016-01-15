package pixsort

import "math/rand"

type Point struct {
	X, Y    int
	R, G, B byte
}

func (a Point) DistanceTo(b Point) int {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dr := int(a.R - b.R)
	dg := int(a.G - b.G)
	db := int(a.B - b.B)
	c := dr*dr + dg*dg + db*db
	return dx*dx + dy*dy + c/256
}

type Undo struct {
	I, J, Score int
}

type Model struct {
	Points []Point
	Score  int
}

func NewModel(points []Point) *Model {
	model := Model{}
	model.Points = points
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
	for m.Closest(i, j) > 50 {
		j = rand.Intn(n)
	}
	s := m.Score
	m.Update(i, j, -1)
	m.Move(i, j)
	m.Update(i, j, 1)
	return Undo{i, j, s}
}

func (m *Model) UndoMove(undo interface{}) {
	u := undo.(Undo)
	m.Move(u.J, u.I)
	m.Score = u.Score
}

func (m *Model) Copy() Annealable {
	points := make([]Point, len(m.Points))
	copy(points, m.Points)
	return &Model{points, m.Score}
}

func (m *Model) Closest(i, j int) int {
	p := m.Points
	a := p[i].DistanceTo(p[j])
	b := a
	if i < j {
		if j < len(p)-1 {
			b = p[i].DistanceTo(p[j+1])
		}
	} else {
		if j > 0 {
			b = p[i].DistanceTo(p[j-1])
		}
	}
	if a <= b {
		return a
	} else {
		return b
	}
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
	if i == j {
		return
	}
	var indexes []int
	if sign < 0 {
		if i < j {
			indexes = []int{i - 1, i, j}
		} else {
			indexes = []int{i - 1, i, j - 1}
		}
	} else {
		if i < j {
			indexes = []int{i - 1, j - 1, j}
		} else {
			indexes = []int{i, j - 1, j}
		}
	}
	n := len(m.Points) - 1
	for _, a := range indexes {
		if a < 0 || a >= n {
			continue
		}
		m.Score += sign * m.Points[a].DistanceTo(m.Points[a+1])
	}
}
