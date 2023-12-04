package utils

type Set[T comparable] struct {
	Map map[T]bool
}

func NewSet[T comparable]() (newSet Set[T]) {
	newSet.Map = make(map[T]bool)
	return
}

func (s *Set[T]) Add(elements ...T) {
	for _, element := range elements {
		s.Map[element] = true
	}
}

func (s *Set[T]) Remove(elements ...T) (ok bool) {
	for _, element := range elements {
		if !s.Exists(element) {
			ok = false
			return
		}
	}
	for _, element := range elements {
		delete(s.Map, element)
	}
	ok = true
	return
}

func (s *Set[T]) Exists(element T) bool {
	return s.Map[element]
}

func (s *Set[T]) Len() int {
	return len(s.Map)
}
