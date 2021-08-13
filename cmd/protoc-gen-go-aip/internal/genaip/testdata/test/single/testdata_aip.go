// Code generated by protoc-gen-go-aip. DO NOT EDIT.
//
// versions:
// 	protoc-gen-go-aip development
// 	protoc v3.17.3
// source: test/single/testdata.proto

package single

import (
	fmt "fmt"
	resourcename "go.einride.tech/aip/resourcename"
	strings "strings"
)

type BookResourceName struct {
	Shelf string
	Book  string
}

func (n BookResourceName) Validate() error {
	if n.Shelf == "" {
		return fmt.Errorf("shelf: empty")
	}
	if strings.IndexByte(n.Shelf, '/') != -1 {
		return fmt.Errorf("shelf: contains illegal character '/'")
	}
	if n.Book == "" {
		return fmt.Errorf("book: empty")
	}
	if strings.IndexByte(n.Book, '/') != -1 {
		return fmt.Errorf("book: contains illegal character '/'")
	}
	return nil
}

func (n BookResourceName) ContainsWildcard() bool {
	return false || n.Shelf == "-" || n.Book == "-"
}

func (n BookResourceName) String() string {
	return resourcename.Sprint(
		"shelves/{shelf}/books/{book}",
		n.Shelf,
		n.Book,
	)
}

func (n BookResourceName) MarshalString() (string, error) {
	if err := n.Validate(); err != nil {
		return "", err
	}
	return n.String(), nil
}

func (n *BookResourceName) UnmarshalString(name string) error {
	return resourcename.Sscan(
		name,
		"shelves/{shelf}/books/{book}",
		&n.Shelf,
		&n.Book,
	)
}