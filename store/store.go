package store

import (
	"fmt"

	"github.com/hashicorp/go-memdb"
)

type Storer interface {
	Add(string, any)
	Get(string)
}

type MemoryStore struct {
	*memdb.MemDB
}

func getMemSchema() *memdb.DBSchema {
	return &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"student": {
				Name: "student",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "ID"},
					},
				},
			},
			"company": {
				Name: "company",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}
}

func NewMemoryStore() MemoryStore {
	db, err := memdb.NewMemDB(getMemSchema())
	if err != nil {
		panic(err)
	}

	return MemoryStore{db}
}

func (ms MemoryStore) Add(table string, value any) {
	add(ms, table, value)
}

func add(ms MemoryStore, table string, value any) {
	txn := ms.Txn(true)
	defer txn.Abort()

	if err := txn.Insert(table, value); err != nil {
		panic(err)
	}

	txn.Commit()
}

func (ms MemoryStore) Get(table string) {
	txn := ms.Txn(false)

	itr, err := txn.Get(table, "id")
	if err != nil {
		panic(err)
	}

	for obj := itr.Next(); obj != nil; obj = itr.Next() {
		fmt.Println(obj)
	}
}
