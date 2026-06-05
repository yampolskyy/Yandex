package main

import (
	"strings"
	"time"
)

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
}

var Tickets []Ticket

func GetTasks(text string, user, status *string) []Ticket {
	Tickets = []Ticket{}
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		els := strings.Split(line, "_")

		if len(els) < 4 {
			continue
		}

		if strings.Contains(els[0], "TICKET") {
			if els[1] != "" {
				for _, i := range []string{"Готово", "В работе", "Не будет сделано"} {
					if els[2] == i {
						parsedTime, err := time.Parse("2006-01-02", els[3])
						if err == nil {
							if user == nil && status == nil {
								Tickets = append(Tickets, Ticket{els[0], els[1], els[2], parsedTime})

							} else if user != nil && status == nil {
								if els[1] == *user {
									Tickets = append(Tickets, Ticket{els[0], els[1], els[2], parsedTime})
								}

							} else if user == nil && status != nil {
								if els[2] == *status {
									Tickets = append(Tickets, Ticket{els[0], els[1], els[2], parsedTime})
								}

							} else if user != nil && status != nil {
								if els[1] == *user && els[2] == *status {
									Tickets = append(Tickets, Ticket{els[0], els[1], els[2], parsedTime})
								}
							}
						}
					}
				}
			}
		}
	}
	return Tickets
}
