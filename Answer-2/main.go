package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	check_err_number(n, err)
	statements := ParseInput(n)

	// Caputure Output
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	inf_loop := Run(n, statements)

	// Release Output
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	// Check infinite loop
	if inf_loop {
		fmt.Println(-1)
		return
	}
	fmt.Printf("%s", out)
}

func ParseInput(n int) []Statement {
	scanner := bufio.NewScanner(os.Stdin)
	var statements []Statement
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			panic(fmt.Errorf("Not enough intpus given"))
		}
		statements = append(statements, Parse(scanner.Text()))
	}
	return statements
}

func Run(n int, statements []Statement) bool {
	v := make(Values)
	seen := make(map[int64]bool)
	v["IR"] = 0
	inf_loop := false
	for {
		if v["IR"] >= int64(n) {
			break
		}

		if seen[v["IR"]] {
			inf_loop = true
			break
		}
		seen[v["IR"]] = true
		s := statements[v["IR"]]
		v["IR"] += 1
		s.Eval(v)
	}

	return inf_loop
}

func Parse(s string) Statement {
	out := strings.Split(s, " ")
	if len(out) < 2 {
		panic(fmt.Errorf("Invlid statement %s", s))
	}
	switch out[0] {
	case "assign":
		return ParseAssign(out[1:])
	case "cout":
		return ParseCout(out[1])
	case "goto":
		return ParseGoto(out[1])
	default:
		panic(fmt.Errorf("unrecognized command %s", out[0]))
	}
}

func ParseAssign(args []string) *Assign {
	return &Assign{
		out: &Variable{name: args[0]},
		op1: ParseValue(args[2]),
		op2: ParseValue(args[4]),
	}
}

func ParseValue(s string) Statement {
	if s == "a" || s == "b" || s == "c" {
		return &Variable{name: s}
	}
	return ParseNumber(s)
}

func ParseNumber(s string) *Number {
	value, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return &Number{value: int64(value)}
}

func ParseCout(v string) Statement {
	return &Cout{out: ParseValue(v)}
}

func ParseGoto(l string) Statement {
	return &Goto{to: ParseNumber(l)}
}

func check_err_number(n int, err error) {
	if err != nil {
		panic(err)
	}
	if n <= 0 {
		panic(fmt.Errorf("Invalid number given %d", n))
	}
}

type Values map[string]int64

type Statement interface {
	Eval(v Values) int64
}

type Variable struct {
	name string
}

func (variable *Variable) Eval(v Values) int64 {
	return v[variable.name]
}

type Number struct {
	value int64
}

func (number *Number) Eval(v Values) int64 {
	return number.value
}

func (number *Number) String() string {
	return "Number"
}

type Assign struct {
	out      *Variable
	op1, op2 Statement
}

func (a *Assign) Eval(m Values) int64 {
	v := a.op1.Eval(m) + a.op2.Eval(m)
	m[a.out.name] = v
	return v
}

func (assign *Assign) String() string {
	return "Assign"
}

type Cout struct {
	out Statement
}

func (cout *Cout) Eval(v Values) int64 {
	out := cout.out.Eval(v)
	fmt.Println(out % (1e10 + 7))
	return out
}

func (cout *Cout) String() string {
	return "Cout"
}

type Goto struct {
	to *Number
}

func (g *Goto) Eval(v Values) int64 {
	v["IR"] = g.to.Eval(v) - 1
	return 0
}

func (g *Goto) String() string {
	return fmt.Sprintf("Goto %d", g.to.value)
}
