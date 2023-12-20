package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	prog, _ := cut(string(data), "\n\n")

	fmt.Println(possibilities(parseProg(prog)))
}

func possibilities(src map[string][]processor) int64 {
	p := constrainedPart{
		x:  constrainedValue{ok: true},
		m:  constrainedValue{ok: true},
		a:  constrainedValue{ok: true},
		s:  constrainedValue{ok: true},
		ok: true,
	}
	allParts := run(src, p, "in", "")
	fmt.Println(allParts)
	var total int64
	for _, p := range allParts {
		total += p.volume()
	}
	return total
}

func run(src map[string][]processor, p constrainedPart, name string, depth string) []constrainedPart {
	fmt.Println(depth, name, p)
	var reply []constrainedPart
	for _, proc := range src[name] {
		pass, next, fail := proc.next(p)
		if pass.ok {
			switch next {
			case "R":
			case "A":
				reply = append(reply, pass)
			default:
				reply = append(reply, run(src, pass, next, depth+"  ")...)
			}
		}
		if fail.ok {
			p = fail
		} else {
			break
		}
	}
	fmt.Println(depth, reply)
	return reply
}

type constraint struct {
	isSet bool
	val   int
}
type constrainedValue struct {
	lt, gt constraint
	ok     bool
}

func (c constrainedValue) width() int64 {
	if c.lt.isSet && c.gt.isSet {
		return int64(c.lt.val-c.gt.val) - 1
	}
	if c.lt.isSet {
		return int64(c.lt.val) - 1
	}
	if c.gt.isSet {
		return 4000 - int64(c.gt.val)
	}
	return 4000
}

type constrainedPart struct {
	x, m, a, s constrainedValue
	ok         bool
}

func (p constrainedPart) volume() int64 {
	return p.x.width() * p.m.width() * p.a.width() * p.s.width()
}

func (cv constrainedValue) String(name string) string {
	if !cv.ok {
		return "!OK"
	}
	if cv.lt.isSet && cv.gt.isSet {
		return fmt.Sprintf("%d<%s<%d", cv.gt.val, name, cv.lt.val)
	}
	if cv.lt.isSet {
		return fmt.Sprintf("%s<%d", name, cv.lt.val)
	}
	if cv.gt.isSet {
		return fmt.Sprintf("%d<%s", cv.gt.val, name)
	}
	return ""
}

func (c constrainedPart) String() string {
	if !c.ok {
		return "!OK !!!"
	}
	var ss []string
	for _, e := range []struct {
		k string
		v constrainedValue
	}{
		{"x", c.x},
		{"m", c.m},
		{"a", c.a},
		{"s", c.s},
	} {
		if s := e.v.String(e.k); s != "" {
			ss = append(ss, s)
		}
	}
	return "{" + strings.Join(ss, ",") + "}"
}

type constrainer struct {
	key        byte
	comparator byte
	value      int
	nextLabel  string
}

func newConstrainer(ss []string) constrainer {
	return constrainer{
		key:        ss[0][0],
		comparator: ss[1][0],
		value:      atoi(ss[2]),
		nextLabel:  ss[3],
	}
}

type processor interface {
	next(cp constrainedPart) (constrainedPart, string, constrainedPart)
}

type passthru string

func (p passthru) next(cp constrainedPart) (constrainedPart, string, constrainedPart) {
	fail := cp
	fail.ok = false
	return cp, string(p), fail
}

func (c constrainer) next(p constrainedPart) (constrainedPart, string, constrainedPart) {
	var cv *constrainedValue
	switch c.key {
	case 'x':
		cv = &p.x
	case 'm':
		cv = &p.m
	case 'a':
		cv = &p.a
	case 's':
		cv = &p.s
	default:
		panic(c.key)
	}
	passValue, failValue := *cv, *cv
	passValue.constrain(c.comparator, c.value)
	failValue.constrain(reverse(c.comparator, c.value))
	*cv = passValue
	p.ok = passValue.ok
	passPart := p

	*cv = failValue
	p.ok = failValue.ok
	failPart := p
	return passPart, c.nextLabel, failPart
}

func (cv *constrainedValue) constrain(comparator byte, v int) {
	switch comparator {
	case '<':
		if cv.lt.isSet {
			if v < cv.lt.val {
				cv.lt.val = v
			}
		} else {
			cv.lt.isSet = true
			cv.lt.val = v
		}
	case '>':
		if cv.gt.isSet {
			if v > cv.gt.val {
				cv.gt.val = v
			}
		} else {
			cv.gt.isSet = true
			cv.gt.val = v
		}
	default:
		panic(comparator)
	}
	if cv.lt.isSet && cv.gt.isSet {
		cv.ok = cv.lt.val > cv.gt.val
	}
}

func reverse(comparator byte, val int) (byte, int) {
	switch comparator {
	case '<':
		return '>', val - 1
	case '>':
		return '<', val + 1
	default:
		panic(comparator)
	}
}

// ([a-z]+)([<>]):([a-z]+|[AR])
var conditionalRegexp = regexp.MustCompile(`([amsx])([<>])([0-9]+):([a-z]+|[AR]).*`)

type piece struct {
	x, m, a, s int
}

func parseProg(s string) map[string][]processor {
	p := make(map[string][]processor)
	for _, line := range strings.Fields(s) {
		name, content := cut(line, "{")
		content = strings.TrimSuffix(content, "}")
		var wf []processor
		for _, section := range strings.Split(content, ",") {
			m := conditionalRegexp.FindStringSubmatch(section)
			if len(m) > 0 {
				wf = append(wf, newConstrainer(m[1:]))
			} else {
				wf = append(wf, passthru(section))
			}
		}
		p[name] = wf
	}
	return p
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func cut(s, tok string) (string, string) {
	a, b, ok := strings.Cut(s, tok)
	if !ok {
		panic(s)
	}
	return a, b
}
