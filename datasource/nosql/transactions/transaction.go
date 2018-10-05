package nosql

import (
	"context"

	"github.com/go-kivik/kivik"
)

// Transaction interface for gorm
type Transaction interface {
	Begin()                               // Begin transaction
	Commit()                              // Commit transaction
	Rollback()                            // Rollback transaction
	DataSource(dbName string) interface{} // Get lowlevel db functions
	Client() *kivik.Client
	Queries() map[string]string     // Queries loaded from files
	LookupQuery(name string) string // Get by name loaded from file query
}

type transaction struct {
	Transaction
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

func (t *transaction) DataSource(dbName string) interface{} {
	db, err := t.client.DB(context.TODO(), dbName)
	if err != nil {
		return nil
	}
	return db
}

func (t *transaction) Client() *kivik.Client {
	return t.client
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
	client  *kivik.Client
	queries map[string]string
}

func NewTransactionFactory(client *kivik.Client, queries map[string]string) TransactionFactory {
	return &transactionFactory{client: client, queries: queries}
}

func (t transactionFactory) Close() {
	//t.db.Close()
}

func (t transactionFactory) BeginNewTransaction() Transaction {
	tx := new(transaction)
	tx.queries = t.queries
	tx.client = t.client
	tx.Begin()
	return tx
}
