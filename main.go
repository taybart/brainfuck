package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/taybart/log"
)

const (
	DONE = 0
)

type Fucked struct {
	Stack []uint8
	ptr   int
	loops []int
}

var code = `
+++++ +++++             initialize counter (cell #0) to 10
[                       use loop to set 70/100/30/10
    > +++++ ++              add  7 to cell #1
    > +++++ +++++           add 10 to cell #2
    > +++                   add  3 to cell #3
    > +                     add  1 to cell #4
<<<< -                  decrement counter (cell #0)
]
> ++ .                  print 'H'
> + .                   print 'e'
+++++ ++ .              print 'l'
.                       print 'l'
+++ .                   print 'o'
> ++ .                  print ' '
<< +++++ +++++ +++++ .  print 'W'
> .                     print 'o'
+++ .                   print 'r'
----- - .               print 'l'
----- --- .             print 'd'
> + .                   print '!'
> .                     print '\n'
`

func main() {
	// log.SetLevel(log.DEBUG)
	b := Fucked{
		Stack: make([]byte, 20),
		ptr:   0,
		loops: []int{},
	}
	b.Process(code)
	fmt.Println(b.String())
}

func NewBeFucked(stackSize int) Fucked {
	return Fucked{
		Stack: make([]byte, stackSize),
		ptr:   0,
		loops: []int{},
	}
}

func (b *Fucked) Process(code string) {
	cleaned := b.clean(code)
	i := 0
	for i != len(cleaned) {
		fmt.Println(b.String())
		op := cleaned[i]
		switch op {
		case '<':
			// log.Debugf("sub pointer %d -> %d\n", b.ptr, b.ptr-1)
			b.ptr--
		case '>':
			// log.Debugf("add pointer %d -> %d\n", b.ptr, b.ptr+1)
			b.ptr++
		case '+':
			v := b.Stack[b.ptr]
			log.Debugf("add [%d] %d -> %d\n", b.ptr, v, v+1)
			b.Stack[b.ptr] = v + 1
		case '-':
			v := b.Stack[b.ptr]
			log.Debugf("sub [%d] %d -> %d\n", b.ptr, v, v-1)
			b.Stack[b.ptr] = v - 1
		case '[':
			// take note of location in code to jump back to
			log.Debug("loop start")
			if !b.hasLoop(i) {
				b.loops = append(b.loops, i)
			}
			// b.looplines = append([]int{i}, b.looplines...)
		case ']':
			if b.Stack[b.ptr] == DONE {
				i = b.loops[len(b.loops)-1]
				log.Debugf("loop to %d\n", i)
				b.loops = b.loops[:len(b.loops)-1]
			}
			// 	log.Debugf("loop done, left %d\n", len(b.loops))
			// b.loops = b.loops[:len(b.loops)-1]
			// 	if len(b.loops) == 2 {
			// 		os.Exit(0)
			// 	}
			// } else {
			// }
		case ',':
			b.Stack[b.ptr] = b.read()
			fmt.Printf("%+v\n", b.Stack)
		case '.':
			fmt.Print("out", string(b.Stack[b.ptr]))
		}
		i++
	}
}

func (b *Fucked) clean(dirty string) string {
	cleaned := ""
	for _, op := range dirty {
		switch op {
		case '<', '>', '+', '-', '[', ']', ',', '.':
			cleaned += string(op)
		}
	}
	return cleaned
}

func (b *Fucked) String() string {
	output := "["
	for i, p := range b.Stack {
		if i == b.ptr {
			output += log.BoldYellow
		}
		output += fmt.Sprintf(" %d ", p)
		if i == b.ptr {
			output += log.Reset
		}
	}
	output += "]"
	return output
}

func (b *Fucked) hasLoop(i int) bool {
	for _, l := range b.loops {
		if i == l {
			return true
		}
	}
	return false
}

func (b *Fucked) read() byte {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return []byte(input)[0]
}
