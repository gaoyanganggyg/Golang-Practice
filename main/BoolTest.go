package main

import (
	"fmt"
	"bytes"
	"strconv"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	fmt.Println(word, bit)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	fmt.Println(word, bit)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	d := IntSet{[]uint64{1, 2, 3}}
	fmt.Println(d.String())

	var s []int
	for x := 195465; x>0; x=x/2{
		d := x%2
		s = append(s, d)
	}
	fmt.Printf("%b\n", 195465)
	l := len(s)
	for i:=0; i<l/2; i++{
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
	var buf bytes.Buffer
	for i:=0; i<l; i++{
		buf.WriteString(strconv.Itoa(s[i]))
	}
	fmt.Println(buf.String())
}