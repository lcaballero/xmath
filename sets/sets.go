package sets

type Hasher func() uint64

func (h Hasher) HashCode() uint64 {
	return h()
}

type Item interface {
	HashCode() uint64
}

type Set map[uint64]Item

func NewSet(items ...Item) Set {
	s := make(Set)
	return s.Add(items...)
}

func (s Set) Union(b Set) Set {
	c := make(Set)
	sets := []Set{ s, b }
	for _,set := range sets {
		for k, v := range set {
			_, ok := c[k]
			if !ok {
				c[k] = v
			}
		}
	}
	return c
}

func (a Set) Intersection(b Set) Set {
	c := make(Set)
	for k, v := range a {
		if b.Has(v) {
			c[k] = v
		}
	}
	return c
}


func (s Set) Has(a Item) bool {
	_, ok := s[a.HashCode()]
	return ok
}

func (s Set) Length() int {
	return len(s)
}

func (s Set) Add(items ...Item) Set {
	for _,a := range items {
		s[a.HashCode()] = a
	}
	return s
}
