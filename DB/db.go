package DB

type DB struct {
	data map[string]string
}

func NewDB()(*DB) {
	return &DB{
		data: map[string]string{},
	}
}

func (db *DB) Begin() (*Tx) {
	newTransaction := &Tx{
		db: db,
		data: map[string]string{},
	}

	return newTransaction
}

func (db *DB) AddRecord(key string, value string) {
	db.data[key] = value
}

func (db *DB) GetRecord(key string) string {
	return db.data[key]
}

func (db *DB) NumeToEqual(value string) int {
	count := 0
	for _, record := range db.data {
		if record == value {
			count++
		} 
	}
	return count
}

func (db *DB) UnSet(key string) {
	delete(db.data, key)
}