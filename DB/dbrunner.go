package DB

import(
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Runner struct {
	db *DB
	currTransaction *Tx
	transactions []*Tx
}

func NewRunner() *Runner {
	newDb := NewDB()
	newTransactions := []*Tx{}
	return &Runner{
		db: newDb,
		transactions: newTransactions,
	}
}

func (r *Runner) Run() {
	ended := true
	for ended {
		fmt.Println(ENTRANCE_MSG)
		reader := bufio.NewReader(os.Stdin)
		char, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error", err.Error())
		}
		char = strings.TrimSuffix(char, "\n")
		instructions := strings.Fields(char)
		countInstructions := len(instructions)
		if countInstructions <= 2 && countInstructions >= 3 {
			fmt.Println(VALID_OPERATION_MSG)
			continue
		}

		switch strings.ToUpper(instructions[0]) {
		case END:
			fmt.Println(FINISHED_MSG)
			ended = false
			break
		case SET:
			isValidInstruction(instructions, 3, fmt.Sprintf(INVALID_COMMAND_MSG, "SET x 10"), r.set)
		case UNSET:
			isValidInstruction(instructions, 2, fmt.Sprintf(INVALID_COMMAND_MSG, "GET x"), r.unSet)
		case NUMEQUALTO:
			isValidInstruction(instructions, 2, fmt.Sprintf(INVALID_COMMAND_MSG, "NUMEQUALTO 10"), r.numeToEqual)
		case GET:
			isValidInstruction(instructions, 2, fmt.Sprintf(INVALID_COMMAND_MSG, "GET x"), r.get)
		case BEGIN:
			r.begin()
		case ROLLBACK:
			r.rollback()
		case COMMIT:
			r.commit()
		default:
			fmt.Println(VALID_OPERATION_MSG)
		}
	}
}

func (r *Runner) numeToEqual(instructions [] string)  {
	value :=  instructions[1]
	if r.hasCurrentTransaction() {
		fmt.Println(r.currTransaction.NumeToEqual(value))
	} else {
		fmt.Println(r.db.NumeToEqual(value))
	}
}

func (r *Runner) unSet(instructions [] string)  {
	key :=  instructions[1]
	if r.hasCurrentTransaction() {
		r.currTransaction.UnSet(key)
	} else {
		r.db.UnSet(key)
	}
}

func (r *Runner) get(instructions []string)  {
	key := instructions[1]
	if r.hasCurrentTransaction() {
		fmt.Println(r.currTransaction.GetRecord(key))
	} else {
		value := r.db.GetRecord(key)
		fmt.Println(value)
	}
}

func (r *Runner) set(instructions []string)  {
	key, value := instructions[1], instructions[2]
	if r.hasCurrentTransaction() {
		r.currTransaction.AddRecord(key, value)
	} else {
		r.db.AddRecord(key, value)
	}
	fmt.Println("")
}

func (r *Runner) begin() {
	newTx := r.db.Begin()
	r.currTransaction = newTx
	r.transactions = append(r.transactions, r.currTransaction)
}

func (r *Runner) commit() {
	for _, value := range r.transactions {
		value.Commit()
	}
	r.currTransaction.Commit()
	r.transactions = []*Tx{}
	r.currTransaction = nil
}

func (r *Runner) rollback() {
	if r.hasTransactions() {
		prevTransactions := r.transactions[:len(r.transactions)-1]
		r.transactions = prevTransactions
		if len(prevTransactions) == 0 {
			r.currTransaction = nil
			return
		}
		r.currTransaction = prevTransactions[len(r.transactions)-1:len(r.transactions)][0]
	}  else {
		fmt.Println(NO_TRANSACTION)
	}
}