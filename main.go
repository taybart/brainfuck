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

type fucked struct {
	code  []rune
	stack []byte
	ptr   int
	loops []int
}

// var code = `
// +++++ +++++             initialize counter (cell #0) to 10
// [                       use loop to set 70/100/30/10
//     > +++++ ++              add  7 to cell #1
//     > +++++ +++++           add 10 to cell #2
//     > +++                   add  3 to cell #3
//     > +                     add  1 to cell #4
// <<<< -                  decrement counter (cell #0)
// ]
// > ++ .                  print 'H'
// > + .                   print 'e'
// +++++ ++ .              print 'l'
// .                       print 'l'
// +++ .                   print 'o'
// > ++ .                  print ' '
// << +++++ +++++ +++++ .  print 'W'
// > .                     print 'o'
// +++ .                   print 'r'
// ----- - .               print 'l'
// ----- --- .             print 'd'
// > + .                   print '!'
// > .                     print '\n'
// `

var code = `
-,+[                         Read first character and start outer character reading loop
    -[                       Skip forward if character is 0
        >>++++[>++++++++<-]  Set up divisor (32) for division loop
                               (MEMORY LAYOUT: dividend copy remainder divisor quotient zero zero)
        <+<-[                Set up dividend (x minus 1) and enter division loop
            >+>+>-[>>>]      Increase copy and remainder / reduce divisor / Normal case: skip forward
            <[[>+<-]>>+>]    Special case: move remainder back to divisor and increase quotient
            <<<<<-           Decrement dividend
        ]                    End division loop
    ]>>>[-]+                 End skip loop; zero former divisor and reuse space for a flag
    >--[-[<->+++[-]]]<[         Zero that flag unless quotient was 2 or 3; zero quotient; check flag
        ++++++++++++<[       If flag then set up divisor (13) for second division loop
                               (MEMORY LAYOUT: zero copy dividend divisor remainder quotient zero zero)
            >-[>+>>]         Reduce divisor; Normal case: increase remainder
            >[+[<+>-]>+>>]   Special case: increase remainder / move it back to divisor / increase quotient
            <<<<<-           Decrease dividend
        ]                    End division loop
        >>[<+>-]             Add remainder back to divisor to get a useful 13
        >[                   Skip forward if quotient was 0
            -[               Decrement quotient and skip forward if quotient was 1
                -<<[-]>>     Zero quotient and divisor if quotient was 2
            ]<<[<<->>-]>>    Zero divisor and subtract 13 from copy if quotient was 1
        ]<<[<<+>>-]          Zero divisor and add 13 to copy if quotient was 0
    ]                        End outer skip loop (jump to here if ((character minus 1)/32) was not 2 or 3)
    <[-]                     Clear remainder from first division if second division was skipped
    <.[-]                    Output ROT13ed character from copy and clear it
    <-,+                     Read next character
]                            End character reading loop
`

func main() {
	// log.SetLevel(log.DEBUG)
	b := fucked{
		stack: make([]byte, 20),
		ptr:   0,
		loops: []int{},
	}
	b.process(code)
	// fmt.Printf("%+v\n", b.stack)
}

func (b *fucked) process(code string) {
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
			if !b.hasLoop(i) {
				b.loops = append(b.loops, i)
			}
			// b.looplines = append([]int{i}, b.looplines...)
		case ']':
			if b.stack[b.ptr] != DONE {
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
			b.stack[b.ptr] = b.read()
			fmt.Printf("%+v\n", b.stack)
		case '.':
			fmt.Print(string(b.stack[b.ptr]))
		}
		i++
	}
}

func (b *fucked) clean(dirty string) string {
	cleaned := ""
	for _, op := range dirty {
		switch op {
		case '<', '>', '+', '-', '[', ']', ',', '.':
			cleaned += string(op)
		}
	}
	return cleaned
}

func (b *fucked) String() string {
	output := "["
	for i, p := range b.stack {
		if i == b.ptr {
			output += log.BoldYellow
		}
		output += fmt.Sprintf(" %d ", p)
		if i == b.ptr {
			output += log.Reset
		}
	}
	return output
}

func (b *fucked) hasLoop(i int) bool {
	for _, l := range b.loops {
		if i == l {
			return true
		}
	}
	return false
}

func (b *fucked) read() byte {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return []byte(input)[0]
}
