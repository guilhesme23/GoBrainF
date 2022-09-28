package compiler

type Tape struct {
	data   []int
	cursor int
	size   int
}

func CreateTape(size int) Tape {
	cursor := 0
	data := make([]int, size)
	return Tape{data, cursor, size}
}

func (t *Tape) ChangeCellValue(val int) {
	limit := 256
	sanitizedValue := val % limit
	tmp := t.data[t.cursor] + sanitizedValue
	if tmp < 0 {
		t.data[t.cursor] = limit + tmp
	} else {
		t.data[t.cursor] = tmp
	}
}

func (t *Tape) MoveCursor(dir int) {
	idx := (t.cursor + dir) % t.size

	if idx < 0 {
		t.cursor = t.size + idx
	} else {
		t.cursor = idx
	}
}

func (t Tape) GetTape() []int {
	return t.data
}
