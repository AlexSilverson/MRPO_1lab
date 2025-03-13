package games

import (
	"math/rand"
	"strconv"
	"time"
)

type progressionGame struct {
	question []int
	answer   int
	rules    string
}

func (cur *progressionGame) Generate() error {
	rand.NewSource(int64(time.Nanosecond))
	len := rand.Intn(6) + 5
	cur.question = make([]int, len)
	q := rand.Intn(7) + 1
	cur.question[0] = rand.Intn(5) + 1
	for i := 1; i < len; i++ {
		cur.question[i] = cur.question[i-1] * q
	}
	cur.answer = rand.Intn(len)
	return nil
}
func (cur *progressionGame) GetAnswer() int {
	return cur.question[cur.answer]
}
func (cur *progressionGame) GetQuestion() []string {
	question := make([]string, len(cur.question))
	for i := 0; i < len(cur.question); i++ {
		if i == cur.answer {
			question[i] = ".."
			continue
		}
		question[i] = strconv.Itoa(cur.question[i])
	}
	return question
}
func (cur *progressionGame) GetRules() string {
	return cur.rules
}
func ProgressionGame() Game {
	return &progressionGame{rules: "What number is missing in the progression?"}
}
