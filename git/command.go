package git

import (
	"os/exec"
	"strings"
	"time"
)

const GIT_COMMAND string = "git"

type Commit struct {
	Hash        string
	ParentHash  string
	Author      string
	AuthorEmail string
	AuthorDate  time.Time
}

func Command(args ...string) string {
	out, _ := exec.Command(GIT_COMMAND, args...).Output()

	return string(out)
}

func Log() []Commit {
	out := Command("log", "--pretty=format:%H::%P::%an::%ae::%ad", "--date=short")

	var lines []string = strings.Split(out, "\n")

	var commits []Commit = make([]Commit, len(lines))
	for i, line := range lines {
		c := strings.Split(line, "::")
		t, _ := time.Parse("2006-01-02", c[4])

		commit := Commit{
			Hash:        c[0],
			ParentHash:  c[1],
			Author:      c[2],
			AuthorEmail: c[3],
			AuthorDate:  t,
		}

		commits[i] = commit
	}

	return commits
}
