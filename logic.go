package main

import "strconv"

type model struct {
	ID   int
	Name string
}

func businessLogic(p Path) ([]*model, error) {
	var result []*model
	for i := 0; i < 10; i++ {
		result = append(result, &model{
			ID:   i,
			Name: "sample" + strconv.Itoa(i),
		})
	}
	return result, nil
}
