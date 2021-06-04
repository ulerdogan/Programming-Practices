package tree

import (
	"errors"
)


type Record struct {
	ID, Parent int 
}

type Node struct {
	ID int
	Children []*Node
}


func Build(records []Record) (*Node, error) {
	lenRec := len(records)
	if lenRec == 0 {
		return nil, nil
	}

	match := make([]int, lenRec)

	for i, p := range records {
		if p.ID < 0 || p.ID >= lenRec {
			return nil, errors.New("Overflow detected...")
		}

		match[p.ID] = i
	}


	n := make([]Node, lenRec)
	for i, _:= range match {
		r := records[match[i]]

		if r.ID != i {
			return nil, errors.New("Node error...")
		}

		valid := (r.ID > r.Parent) || (r.ID == 0 && r.Parent == 0)

		if !valid {
			return nil, errors.New("Invalid parent...")
		}

		n[i].ID = i
		if i != 0 {
			p := &n[r.Parent]
			p.Children = append(p.Children, &n[i])
		}


	}
	
	
	return &n[0], nil
}