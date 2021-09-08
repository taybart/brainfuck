package main

import (
	"fmt"

	"github.com/taybart/log"
)

const (
	DONE = 0
)

type fucked struct {
	code      []rune
	stack     []byte
	ptr       int
	looplines []int
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
	b := fucked{
		stack:     make([]byte, 20),
		ptr:       0,
		looplines: []int{},
	}
	b.process(code)
	fmt.Printf("%+v\n", b.stack)
}

func (b *fucked) process(code string) {
	i := 0
	for i != len(code) {
		op := code[i]
		switch op {
		case '<':
			log.Debugf("sub pointer %d -> %d\n", b.ptr, b.ptr-1)
			b.ptr--
		case '>':
			log.Debugf("add pointer %d -> %d\n", b.ptr, b.ptr+1)
			b.ptr++
		case '+':
			v := b.stack[b.ptr]
			log.Debugf("add [%d] %d -> %d\n", b.ptr, v, v+1)
			b.stack[b.ptr] = v + 1
		case '-':
			v := b.stack[b.ptr]
			log.Debugf("sub [%d] %d -> %d\n", b.ptr, v, v-1)
			b.stack[b.ptr] = v - 1
		case '[':
			// take note of location in code to jump back to
			log.Debug("loop start")
			b.looplines = append(b.looplines, i)
		case ']':
			if b.stack[b.ptr] == DONE {
				log.Debug("loop done")
				b.looplines = b.looplines[:len(b.looplines)-1]
			} else {
				i = b.looplines[len(b.looplines)-1]
				log.Debugf("loop to %d\n", i)
				continue
			}
		case '.':
			fmt.Print(string(b.stack[b.ptr]))
		}
		i++
	}
}
