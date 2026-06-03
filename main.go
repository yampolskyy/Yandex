package main

import (
	"fmt"
	"sort"
)

type Worker struct {
	Name            string
	Position        string
	Salary          uint
	ExperienceYears uint
}

type Company struct {
	workers []Worker
}

type CompanyInterface interface {
	AddWorkerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	if name != "" && position != "" && salary != 0 && experience != 0 {
		c.workers = append(c.workers, Worker{name, position, salary, experience})
		return nil
	} else {
		return fmt.Errorf("! EMPTY VALUE !")
	}
}

func (c *Company) SortWorkers() ([]string, error) {
	var answer []string
	sort.Slice(c.workers, func(i, j int) bool {
		return c.workers[i].Salary*c.workers[i].ExperienceYears > c.workers[j].Salary*c.workers[j].ExperienceYears
	})

	for i := 0; i < len(c.workers); i++ {
		build := fmt.Sprintf("%s — %d — %s", c.workers[i].Name, c.workers[i].Salary*c.workers[i].ExperienceYears, c.workers[i].Position)
		answer = append(answer, build)
	}
	if len(answer) != 0 {
		return answer, nil
	} else {
		return nil, fmt.Errorf("! EMPTY VALUE !")
	}
}
