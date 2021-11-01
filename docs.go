/*
Package gokv contains a simple key-value store abstraction in the form of a Go interface.
Implementations of the gokv.Store interface can be found in the sub-packages.

Usage

Example code for using Redis:

	package main

	import (
		"fmt"

		"github.com/gmaisto/kvstore
		"github.com/gmaisto/kvstore/badgerdb"
	)

	type foo struct {
		Bar string
	}

	func main() {
		options := badgerdb.DefaultOptions

		// Create client
		client, err := badgerdb.NewClient(options)
		if err != nil {
			panic(err)
		}
		defer client.Close()

		// Store, retrieve, print and delete a value
		interactWithStore(client)
	}

	// interactWithStore stores, retrieves, prints and deletes a value.
	// It's completely independent of the store implementation.
	func interactWithStore(store kvstore.Store) {
		// Store value
		val := foo{
			Bar: "baz",
		}
		err := store.Set("foo123", val)
		if err != nil {
			panic(err)
		}

		// Retrieve value
		retrievedVal := new(foo)
		found, err := store.Get("foo123", retrievedVal)
		if err != nil {
			panic(err)
		}
		if !found {
			panic("Value not found")
		}

		fmt.Printf("foo: %+v", *retrievedVal) // Prints `foo: {Bar:baz}`

		// Delete value
		err = store.Delete("foo123")
		if err != nil {
			panic(err)
		}
	}

More details can be found on https://github.com/gmaisto/kvstore.
*/
package kvstore
