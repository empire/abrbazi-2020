package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	check_err_number(n, err)
	Execute(n)
}

type Repo struct {
	local  []string
	remote [][]string
}

func Execute(n int) {
	scanner := bufio.NewScanner(os.Stdin)
	repo := new(Repo)
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			panic(fmt.Errorf("Not enough intpus given"))
		}
		repo.ExecuteCommand(scanner.Text())
	}
}

func (repo *Repo) ExecuteCommand(s string) {
	out := strings.Split(s, " ")
	if len(out) < 2 {
		panic(fmt.Errorf("Invlid statement %s", s))
	}
	out = out[1:]
	switch out[0] {
	case "add":
		repo.Add(out[1:])
	case "clear":
		repo.Clear()
	case "del":
		repo.Del()
	case "pull":
		repo.Pull()
	case "checkout":
		repo.Checkout()
	case "commit":
		repo.Commit()
	default:
		panic(fmt.Errorf("unrecognized command %s", out[0]))
	}
}

func (repo *Repo) Add(out []string) {
	repo.local = append(repo.local, strings.Join(out, " "))
}

func (repo *Repo) Del() {
	if len(repo.local) == 0 {
		return
	}
	repo.local = repo.local[:len(repo.local)-1]
}

func (repo *Repo) Clear() {
	repo.local = nil
}

func (repo *Repo) Pull() {
	var latest []string
	if len(repo.remote) != 0 {
		latest = repo.remote[len(repo.remote)-1]
	}
	fmt.Println(len(latest))
	if len(repo.remote) != 0 {
		fmt.Println(strings.Join(latest, "\n"))
	}
}

func (repo *Repo) Checkout() {
	repo.remote = repo.remote[:len(repo.remote)-1]
}

func (repo *Repo) Commit() {
	repo.remote = append(repo.remote, repo.local)
	repo.local = nil
}

func check_err_number(n int, err error) {
	if err != nil {
		panic(err)
	}
	if n <= 0 {
		panic(fmt.Errorf("Invalid number given %d", n))
	}
}
