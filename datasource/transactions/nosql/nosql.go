package nosql

import (
	"context"
	"reflect"

	"github.com/couchbase/gocb"

	"github.com/go-kivik/kivik"
)

// Transaction interface for nosql
type Transaction interface {
	Begin()                               // Begin transaction
	Commit()                              // Commit transaction
	Rollback()                            // Rollback transaction
	DataSource(dbName string) interface{} // Get lowlevel db functions
	Client() interface{}
	Queries() map[string]string     // Queries loaded from files
	LookupQuery(name string) string // Get by name loaded from file query
}

type transaction struct {
	Transaction
	client  interface{}
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
	if reflect.TypeOf(t.client) == reflect.TypeOf(&kivik.Client{}) {
		db, err := t.client.(*kivik.Client).DB(context.TODO(), dbName)
		if err != nil {
			return nil
		}
		return db
	}
	if reflect.TypeOf(t.client) == reflect.TypeOf(&gocb.Cluster{}) {
		db, err := t.client.(*gocb.Cluster).OpenBucket(dbName, "")
		if err != nil {
			return nil
		}
		return db
	}

	return nil
}

func (t *transaction) Client() interface{} {
	return t.client
}

func (t *transaction) Queries() map[string]string {
	return t.queries
}
func (t *transaction) LookupQuery(name string) string {
	query, _ := t.queries[name]
	return query
}

// TransactionFactory interface
type TransactionFactory interface {
	BeginNewTransaction() Transaction
	Close()
}

type transactionFactory struct {
	client  interface{}
	queries map[string]string
}

// NewTransactionFactory create transaction factory
func NewTransactionFactory(client interface{}, queries map[string]string) TransactionFactory {
	return &transactionFactory{client: client, queries: queries}
}

func (t transactionFactory) Close() {

}

func (t transactionFactory) BeginNewTransaction() Transaction {
	tx := new(transaction)
	tx.queries = t.queries
	tx.client = t.client
	tx.Begin()
	return tx
}
