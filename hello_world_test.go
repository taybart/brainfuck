package main

import "testing"

func TestHelloWorld(t *testing.T) {
	b := NewBeFucked(6)

	code := `
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

	b.Process(code)
	stack := []uint8{0, 87, 100, 33, 10, 0}
	for i, s := range stack {
		if b.Stack[i] != s {
			t.Fatalf("mismatched stacks:\n%+v\n!=\n%+v\n", stack, b.Stack)
		}
	}
}
