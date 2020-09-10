package tournament

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
	"text/template"
)

type Matches struct {
	Player        string
	MatchesPlayed int
	MatchesWon    int
	MatchesDrawn  int
	MatchesLost   int
	Points        int
}

func newMatches(player string) *Matches {
	return &Matches{Player: player}
}

var outputTemplate string = `Team                           | MP |  W |  D |  L |  P
{{range $p := .}}{{$p.Player}} |  {{$p.MatchesPlayed}} |  {{$p.MatchesWon}} |  {{$p.MatchesDrawn}} |  {{$p.MatchesLost}} |  {{$p.Points}}
{{end}}`

func Tally(r io.Reader, w io.Writer) error {

	records, err := validationAndReading(r)
	if err != nil {
		return err
	}

	print(classification(rounds(records)), w)

	return nil
}

func validationAndReading(r io.Reader) ([][]string, error) {

	csvReander := csv.NewReader(r)
	csvReander.Comma = ';'
	csvReander.Comment = '#'

	matches, err := csvReander.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, j := range matches {
		if len(j) != 3 {
			return nil, errors.New("Error")
		} else if !(j[2] == "win" || j[2] == "loss" || j[2] == "draw") {
			return nil, errors.New("Error")
		}
	}
	return matches, nil
}

func rounds(records [][]string) map[string]*Matches {

	var team map[string]*Matches = make(map[string]*Matches)

	for _, record := range records {
		if record[2] == "win" {
			if _, ok := team[record[0]]; !ok {
				team[record[0]] = newMatches(record[0])
			}
			if _, ok := team[record[1]]; !ok {
				team[record[1]] = newMatches(record[1])
			}
			team[record[0]].MatchesPlayed++
			team[record[0]].MatchesWon++
			team[record[1]].MatchesPlayed++
			team[record[1]].MatchesLost++

		} else if record[2] == "loss" {
			if _, ok := team[record[0]]; !ok {
				team[record[0]] = newMatches(record[0])
			}
			if _, ok := team[record[1]]; !ok {
				team[record[1]] = newMatches(record[1])
			}
			team[record[1]].MatchesPlayed++
			team[record[1]].MatchesWon++
			team[record[0]].MatchesPlayed++
			team[record[0]].MatchesLost++
		} else if record[2] == "draw" {
			if _, ok := team[record[0]]; !ok {
				team[record[0]] = newMatches(record[0])
			}
			if _, ok := team[record[1]]; !ok {
				team[record[1]] = newMatches(record[1])
			}
			team[record[0]].MatchesPlayed++
			team[record[0]].MatchesDrawn++
			team[record[1]].MatchesPlayed++
			team[record[1]].MatchesDrawn++
		}
	}
	return team
}

func classification(team map[string]*Matches) []Matches {

	placing := []Matches{}

	for _, value := range team {
		value.Points = (value.MatchesWon * 3) + value.MatchesDrawn
		placing = append(placing, *value)
	}

	sort.SliceStable(placing, func(i, j int) bool {
		if placing[i].Points > placing[j].Points {
			return true
		} else if placing[i].Points < placing[j].Points {
			return false
		} else {
			s := []string{placing[i].Player, placing[j].Player}
			sort.Strings(s)

			return placing[i].Player == s[0]
		}
	})

	return placing
}

func print(player []Matches, w io.Writer) {
	t := template.Must(template.New("output").Parse(outputTemplate))

	for i, _ := range player {
		player[i].Player = fmt.Sprintf("%s%s", player[i].Player, strings.Repeat(" ", 30-len(player[i].Player)))
	}

	t.Execute(w, player)

}
