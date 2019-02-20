package list

type List struct {
	Front *Elem
	Back  *Elem
}
type Elem struct {
	Value interface{}
	Next  *Elem
	Prev  *Elem
}

func NewList() *List {
	return &List{}
}

func (s *List) PushBack(val interface{}) *Elem {
	e := &Elem{Value: val}
	e.Next = nil
	if s.Front == nil {
		s.Front = e
		e.Prev = nil
	} else {
		s.Back.Next = e
		e.Prev = s.Back
	}

	s.Back = e

	return e
}

func (s *List) Remove(e *Elem) interface{} {
	for pe := s.Front; pe != nil; pe = pe.Next {
		if pe == e {
			if pe == s.Front {
				if s.Front == s.Back {
					s.Front = nil
					s.Back = nil
				} else {
					s.Front = s.Front.Next
					if s.Front != nil {
						s.Front.Prev = nil
					}
				}
			} else if pe == s.Back {
				s.Back = s.Back.Prev
				if s.Back != nil {
					s.Back.Next = nil
				}
			} else {
				prev := pe.Prev
				next := pe.Next
				prev.Next = next
				next.Prev = prev
			}

			return pe.Value
		}
	}

	return nil
}
