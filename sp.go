package main

import (
	"strings"
	"time"
)

func QuizRunner(questions, answers []string, answerCh chan string) int {
	correct := 0

	for i := range questions {
		timeout := make(chan bool, 1)
		go func() {
			time.Sleep(1 * time.Second)
			timeout <- true
		}()

		select {
		case <-timeout:
			continue
		case answer := <-answerCh:
			userAnswer := strings.TrimSpace(strings.ToLower(answer))
			correctAnswer := strings.TrimSpace(strings.ToLower(answers[i]))

			if userAnswer == correctAnswer {
				correct++
			}
		}
	}

	return correct
}
