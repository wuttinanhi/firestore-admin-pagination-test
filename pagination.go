package main

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Pagination struct {
	client     *firestore.Client
	ctx        context.Context
	collection string
}

func NewPagination(client *firestore.Client, ctx context.Context, collection string) *Pagination {
	return &Pagination{
		client:     client,
		ctx:        ctx,
		collection: collection,
	}
}

func (p *Pagination) GetCollection() *firestore.CollectionRef {
	return p.client.Collection(p.collection)
}

func (p *Pagination) Paginate(page int, size int, orderBy string, sort string) ([]*firestore.DocumentSnapshot, error) {
	var docs []*firestore.DocumentSnapshot

	// set query
	col := p.GetCollection()
	if col == nil {
		return nil, errors.New("collection not found")
	}

	// create new query
	q := col.Query

	// set order
	if orderBy != "" {
		if sort == "asc" {
			q = q.OrderBy(orderBy, firestore.Asc)
		} else {
			q = q.OrderBy(orderBy, firestore.Desc)
		}
	}

	// set limit
	q = q.Limit(size)

	// set offset
	if page > 1 {
		q = q.Offset((page - 1) * size)
	}

	// get documents
	di := q.Documents(p.ctx)
	for {
		doc, err := di.Next()
		if err != nil {
			// check if error is "no more items"
			if err == iterator.Done {
				break
			}

			// else return error
			return nil, err
		}

		// check if done
		if doc == nil {
			break
		}

		docs = append(docs, doc)
	}

	return docs, nil
}
