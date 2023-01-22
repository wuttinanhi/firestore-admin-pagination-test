package main

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
)

func main() {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		fmt.Printf("error initializing app: %v", err)
		return
	}

	// initialize a Firestore client
	ctx = context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Printf("error initializing Firestore client: %v", err)
		return
	}

	// check if collection exists
	if !IsCollectionExists(client, ctx, "mocks") {
		// if not then mock
		fmt.Println("creating mock data...")
		CreateMock(client, ctx)
	}

	p := NewPagination(client, ctx, "mocks")
	docs, err := p.Paginate(1, 5, "isbn", "asc")
	if err != nil {
		fmt.Printf("error paginating: %v", err)
		return
	}

	fmt.Printf("total documents: %d\n", len(docs))
	for _, doc := range docs {
		fmt.Println(doc.Data())
	}
}
