package DB

import "fmt"

type command func([]string)

func(r *Runner) hasTransactions() bool {
	return len(r.transactions) > 0
}

func (r *Runner) hasCurrentTransaction() bool {
	return r.currTransaction != nil
}

func isIndexOnArray(arrayLen int, index int) bool {
	if (arrayLen -1) == index {
		return true
	}

	return false
}

func isValidInstruction(instructions []string, validLen int, errMsg string, fn command) {
	if len(instructions) != validLen {
		fmt.Println(errMsg)
		return
	}

	fn(instructions)
}