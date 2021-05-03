package robotname

import (
	"errors"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var names = map[string]bool{}

func (robot *Robot) Name() (string, error) {
	if robot.name != "" {
		return robot.name, nil
	}

	if len(names) >= 10*10*10*26*26 {
		return "", errors.New("overflow")
	}

	for {
		robot.name = nameGenerator()
		if !names[robot.name] {
			names[robot.name] = true
			break
		}
	}

	return robot.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func nameGenerator() (name string) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(25) + 65
	name += string(r)
	r = rand.Intn(25) + 65
	name += string(r)

	r = rand.Intn(9) + 48
	name += string(r)
	r = rand.Intn(9) + 48
	name += string(r)
	r = rand.Intn(9) + 48
	name += string(r)

	return name
}
