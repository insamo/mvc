package nosql

import (
	"context"

	"github.com/go-kivik/kivik"
)

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
	db      *kivik.DB
	client  *kivik.Client
	queries map[string]string
}

func (t *transaction) Begin() {

}

func (t *transaction) Commit() {
	//t.tx.Commit()
}

func (t *transaction) Rollback() {
	//t.tx.Rollback()
}

func (t *transaction) DataSource() interface{} {
	return t.db
}

func (t *transaction) Queries() map[string]string {
	return t.queries
}
func (t *transaction) LookupQuery(name string) string {
	query, _ := t.queries[name]
	return query
}

type TransactionFactory interface {
	BeginNewTransaction(dbName string) Transaction
	Close()
}

type transactionFactory struct {
	client  *kivik.Client
	queries map[string]string
}

func NewTransactionFactory(client *kivik.Client, queries map[string]string) TransactionFactory {
	return &transactionFactory{client: client, queries: queries}
}

func (t transactionFactory) Close() {
	//t.db.Close()
}

func (t transactionFactory) BeginNewTransaction(dbName string) Transaction {
	tx := new(transaction)

	db, err := tx.client.DB(context.TODO(), dbName)
	if err != nil {
		return nil
	}
	tx.db = db
	tx.queries = t.queries
	tx.Begin()
	return tx
}
