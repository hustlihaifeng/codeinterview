package main

import "fmt"

func main() {
	var cm CMap
	cm.Store(1, 1)
	rst, ok := cm.Load(1)
	fmt.Println(cm.Load(1))
	fmt.Printf("%#v %#v\n", rst, ok)
	fmt.Printf("%T %T\n", rst, ok)
}

type CMap struct {
	m map[interface{}]interface{}
}

func (s *CMap) Load(key interface{}) (interface{}, bool) {
	if s.m == nil {
		s.m = make(map[interface{}]interface{})
	}

	if p, ok := s.m[key]; ok {
		return *(p.(*interface{})), ok
	}
	return nil, false
}

func (s *CMap) Store(key, value interface{}) {
	if s.m == nil {
		s.m = make(map[interface{}]interface{})
	}

	s.m[key] = &value
}
