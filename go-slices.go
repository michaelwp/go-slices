package go_slices

/*
author: michael.wenceslaus@gmail.com
*/

type Slices []interface{}

func (s Slices) Shift() Slices {
	if len(s) > 1 {
		s = s[1:]
	} else {
		s = nil
	}

	return s
}

func (s Slices) Pop() Slices {
	if len(s) > 1 {
		s = s[:len(s)-1]
	} else {
		s = nil
	}

	return s
}

func (s Slices) Remove(element int) Slices {
	if len(s) > 1 {
		switch element {
		case 0:
			s = s.Shift()
		case len(s) - 1:
			s = s.Pop()
		default:
			s = append(s[:element], s[element+1:]...)
		}
	} else {
		s = nil
	}

	return s
}
