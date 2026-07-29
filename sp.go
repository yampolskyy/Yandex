package main

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"strings"
	"time"
)

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
}

func GetTasks(ctx context.Context, r io.Reader, w io.Writer, user, status *string, timeout time.Duration) error {
	var Tickets = []Ticket{}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var elements = strings.Split(line, "_")
		if len(elements) < 4 {
			continue
		}

		if strings.Contains(elements[0], "TICKET") {
			if elements[1] != "" {
				for _, i := range []string{"Готово", "В работе", "Не будет сделано"} {
					if elements[2] == i {
						parsedTime, err := time.Parse("2006-01-02", elements[3])
						if err == nil {
							if user == nil && status == nil {
								Tickets = append(Tickets, Ticket{elements[0], elements[1], elements[2], parsedTime})

							} else if user != nil && status == nil {
								if elements[1] == *user {
									Tickets = append(Tickets, Ticket{elements[0], elements[1], elements[2], parsedTime})
								}

							} else if user == nil && status != nil {
								if elements[2] == *status {
									Tickets = append(Tickets, Ticket{elements[0], elements[1], elements[2], parsedTime})
								}

							} else if user != nil && status != nil {
								if elements[1] == *user && elements[2] == *status {
									Tickets = append(Tickets, Ticket{elements[0], elements[1], elements[2], parsedTime})
								}
							}
						}
					}
				}
			}
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
	}
	var encoder = json.NewEncoder(w)
	return encoder.Encode(Tickets)
}
