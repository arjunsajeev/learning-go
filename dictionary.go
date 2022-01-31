package main

import "errors"

type Dictionary map[string]string

var ErrorNotFound = errors.New("not found")

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if ok {
		return value, nil
	}
	return "", ErrorNotFound
}
