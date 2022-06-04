package DB

type Tx struct {
	db *DB
	data map[string]string
}

func (tx *Tx) AddRecord(key string, value string){
	tx.data[key] = value
}

func (tx *Tx) GetRecord(key string) string {
	return tx.data[key]
}

func (tx *Tx) Commit() {
	for key, value := range tx.data {
		tx.db.AddRecord(key, value)
	}
}

func (tx *Tx) NumeToEqual(value string) int {
	count := 0
	for _, record := range tx.data {
		if record == value {
			count++
		} 
	}
	return count
}

func (tx *Tx) UnSet(key string)  {
	delete(tx.data, key)
}