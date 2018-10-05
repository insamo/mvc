package transactions

import "github.com/jinzhu/gorm"

// Transaction interface for gorm
type Transaction interface {
	Begin()                         // Begin transaction
	Commit()                        // Commit transaction
	Rollback()                      // Rollback transaction
	DataSource() interface{}        // Get lowlevel db functions
	Queries() map[string]string     // Queries loaded from files
	LookupQuery(name string) string // Get by name loaded from file query
}

type transaction struct {
	Transaction
	db      *gorm.DB
	tx      *gorm.DB
	queries map[string]string
}

func (t *transaction) Begin() {
	t.tx = t.db.Begin()
	if t.tx.Error != nil {
		panic(t.tx.Error)
	}
}

func (t *transaction) Commit() {
	t.tx.Commit()
}

func (t *transaction) Rollback() {
	t.tx.Rollback()
}

func (t *transaction) DataSource() interface{} {
	return t.tx
}

func (t *transaction) Queries() map[string]string {
	return t.queries
}
func (t *transaction) LookupQuery(name string) string {
	query, _ := t.queries[name]
	return query
}

type TransactionFactory interface {
	BeginNewTransaction() Transaction
	Close()
}

type transactionFactory struct {
	db      *gorm.DB
	queries map[string]string
}

func NewTransactionFactory(db *gorm.DB, queries map[string]string) TransactionFactory {
	return &transactionFactory{db: db, queries: queries}
}

func (t transactionFactory) Close() {
	t.db.Close()
}

func (t transactionFactory) BeginNewTransaction() Transaction {
	tx := new(transaction)
	tx.db = t.db
	tx.queries = t.queries
	tx.Begin()
	return tx
}
