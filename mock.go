package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func ListCollections(client *firestore.Client, ctx context.Context) {
	ci := client.Collections(ctx)
	for {
		c, err := ci.Next()
		if err != nil {
			break
		}

		fmt.Println(c.ID, c.Path)
	}
}

func IsCollectionExists(client *firestore.Client, ctx context.Context, collection string) bool {
	ci := client.Collections(ctx)
	for {
		c, err := ci.Next()
		if err != nil {
			break
		}

		if c.ID == collection {
			return true
		}
	}

	return false
}

func CreateMock(client *firestore.Client, ctx context.Context) {
	// mock array
	mockArr := []string{"AAA", "BBB", "CCC", "DDD", "EEE", "FFF", "GGG"}

	for i, v := range mockArr {
		// create a document
		_, _, err := client.Collection("mocks").Add(ctx, map[string]interface{}{
			"isbn": i,
			"name": v,
		})
		if err != nil {
			fmt.Printf("error adding document: %v", err)
			return
		}
	}
}

func DestroyMock(client *firestore.Client, ctx context.Context) {
	client.RunTransaction(ctx, func(ctx context.Context, t *firestore.Transaction) error {
		// get all documents
		docs := client.Collection("mocks").Documents(ctx)
		for {
			doc, err := docs.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			// delete document
			t.Delete(doc.Ref)
		}

		return nil
	})
}
