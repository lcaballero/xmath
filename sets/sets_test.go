package sets

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)


func TestSets(t *testing.T) {

	Convey("Intersection of {A,B,C,D} + {D,E,F,G} have length == 1 with only D.", t, func() {
		n := NewNotes()
		a := NewSet(n.A, n.B, n.C, n.D)
		b := NewSet(n.D, n.E, n.F, n.G)
		c := a.Intersection(b)
		So(c.Length(), ShouldEqual, 1)
	})

	Convey("Union of {A,B,C} + {D,E,F} should have length == 6.", t, func() {
		n := NewNotes()
		a := NewSet(n.A, n.B, n.C)
		b := NewSet(n.D, n.E, n.F)
		c := a.Union(b)
		So(c.Length(), ShouldEqual, a.Length() + b.Length())
	})

	Convey("Union of {Tail} + {Heads, Tails} should have length == 2.", t, func() {
		a := NewSet(Heads)
		b := NewSet(Tails, Heads)
		c := a.Union(b)
		So(c.Length(), ShouldEqual, b.Length())
	})

	Convey("Size after adding item should increment size.", t, func() {
		s := make(Set)
		s.Add(Heads)
		So(s.Length(), ShouldEqual, 1)
		s.Add(Tails)
		So(s.Length(), ShouldEqual, 2)
	})

	Convey("Empty set should have size 0", t, func() {
		s := make(Set)
		So(s.Length(), ShouldEqual, 0)
	})
}

func NewHasher(a int) Hasher {
	return func() uint64 {
		return uint64(a)
	}
}


var Heads Hasher = NewHasher(2)
var Tails Hasher = NewHasher(1)

type Notes struct {
	A, B, C, D, E, F, G Hasher
}
func NewNotes() *Notes {
	return &Notes{
		A: NewHasher(int('a')),
		B: NewHasher(int('b')),
		C: NewHasher(int('c')),
		D: NewHasher(int('d')),
		E: NewHasher(int('e')),
		F: NewHasher(int('f')),
		G: NewHasher(int('g')),
	}
}
