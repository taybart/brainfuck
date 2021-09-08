package main

import "fmt"

// char array[30000] = {0}; char *ptr = array;
// >	++ptr;
// <	--ptr;
// +	++*ptr;
// -	--*ptr;
// .	putchar(*ptr);
// ,	*ptr = getchar();
// [	while (*ptr) {
// ]	}

type fucked struct {
	stack     []byte
	ptr       int
	looplines []int
}

func main() {
	code := []rune{'+', '+', '>', '+'}
	b := fucked{
		stack:     make([]byte, 20),
		ptr:       0,
		looplines: []int{},
	}
	for _, op := range code {
		b.process(op)
	}
	fmt.Printf("%+v\n", b.stack)
}

func (b *fucked) process(in rune) int {
	switch in {
	case '<':
		b.ptr--
	case '>':
		b.ptr++
	case '+':
		b.stack[b.ptr]++
	case '-':
		b.stack[b.ptr]--
	case '[':
		b.looplines = append(b.looplines, b.ptr)
	case ']':
		if b.stack[b.ptr] != 0 {
			b.ptr = b.looplines[len(b.looplines)-1]
			return b.ptr
		}
	case '.':
		fmt.Println(b.stack[b.ptr])
	}
	return -1
}
