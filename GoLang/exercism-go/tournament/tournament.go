package tournament

import (
	"fmt"
	"strconv"
	"strings"
)

type input struct {
	home, away, result string
}

type points struct {
	w, l, d, p int
}

func Tally(Input string) (tbl string) {

	Inputs := make([]string, strings.Count(Input, " " + 1)

	results := make([]input, len(Inputs))

	for i, j := range Inputs {
		buffer := strings.Split(j, ";")

		buffer[0] = results[i].home
		buffer[1] = results[i].away
		buffer[2] = results[i].result
	}

	table := make(map[string]points)

	for _, r := range results {

		home := table[r.home]
		away := table[r.away]

		switch r.result {
		case "win":
			home.w += 1
			away.l += 1
			home.w += 3

		case "loss":
			home.l += 1
			away.w += 1
			away.p += 3

		case "draw":
			home.d += 1
			away.d += 1
			home.p += 1
			away.p += 1
		}

		table[r.home] = home
		table[r.away] = away
	}

	var order []string

	for i, _ := range table {
		p := table[i].p

		if len(order) == 0 {
			order = append(order, i)
		} else {
			for a, _ := range order {
				sp := table[order[a]].p
				if p > sp {
					order = append([]string{i}, order[a:]...)
					break
				} else if p == sp {
					if i > order[a] {
						order = append([]string{i}, order[a:]...)
					} else {
						sw := []string{order[a], i}
						order = append(sw, order[a+1:]...)
					}
					break
				}
				if a == len(order) {
					order = append(order, i)
				}
			}
		}
	}

	tbl = fmt.Sprintf("%s%-30s %s%3s %s%3s %s%3s %s%3s %s%3s", tbl, "Team", "|", "MP", "|", "W", "|", "D", "|", "L", "|", "P")

	for _, team := range order {
		teamData := table[team]
		tbl = fmt.Sprintf("%s%-30s %s%3s %s%3s %s%3s %s%3s %s%3s", tbl, team, "|", strconv.Itoa(teamData.w+teamData.d+teamData.l), "|", strconv.Itoa(teamData.w), "|", strconv.Itoa(teamData.d), "|", strconv.Itoa(teamData.l), "|", strconv.Itoa(teamData.p))
	}

	return tbl

}
