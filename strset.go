// Package strset implements set functionality on top of map. It only supports
// sets of strings. Sets are not safe for concurrent use, just as maps on which
// they are based.
package strset

// Set represents a set of string values
type Set map[string]struct{}

// New returns a new set, if arguments given, they are added to the set
func New(elem ...string) Set {
	s := make(Set)
	s.Add(elem...)
	return s
}

// Add adds new values to the set
func (s Set) Add(elem ...string) {
	for _, e := range elem {
		s[e] = struct{}{}
	}
}

// Del removes specified values from the set
func (s Set) Del(elem ...string) {
	for _, e := range elem {
		delete(s, e)
	}
}

// Contains returns true if the set contains the argument
func (s Set) Contains(elem string) bool {
	_, ok := s[elem]
	return ok
}

// Intersect returns a new set that only contains values that are contained in
// both sets
func (s Set) Intersect(t Set) Set {
	r := New()
	for k := range s {
		if _, ok := t[k]; ok {
			r[k] = struct{}{}
		}
	}
	return r
}

// AsSlice returns slice with all the values from the set. The order of values
// is not defined.
func (s Set) AsSlice() []string {
	res := make([]string, len(s))
	var i int
	for k := range s {
		res[i] = k
		i++
	}
	return res
}
