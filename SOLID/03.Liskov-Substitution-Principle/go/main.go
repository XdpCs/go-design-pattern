package main

import (
	"errors"
	"fmt"
)

// @Title        main.go
// @Description
// @Create       XdpCs 2024-05-20 13:14
// @Update       XdpCs 2024-05-20 13:14

type Document struct {
	data     []byte
	fileName string
}

// Before Liskov Substitution Principle

type IDocument interface {
	open() ([]byte, error)
	save([]byte) error
}

type ReadOnlyDocument struct {
	Document
}

func (r *ReadOnlyDocument) open() ([]byte, error) {
	return r.data, nil
}

func (r *ReadOnlyDocument) save(data []byte) error {
	return errors.New("cannot save read-only document")
}

type WritableDocument struct {
	Document
}

func (w *WritableDocument) open() ([]byte, error) {
	return w.data, nil
}

func (w *WritableDocument) save(data []byte) error {
	w.data = data
	return nil
}

func processDocument(d IDocument) error {
	data, err := d.open()
	if err != nil {
		return fmt.Errorf("opening document error: %w", err)
	}
	fmt.Println("Document opened with data:", string(data))
	err = d.save(data)
	if err != nil {
		return fmt.Errorf("saving document error: %w", err)
	}
	fmt.Println("Document saved successfully")
	return nil
}

// After Liskov Substitution Principle

type ModifyIDocument interface {
	open() ([]byte, error)
}

type ModifyIWritableDocument interface {
	ModifyIDocument
	save([]byte) error
}

type ModifyReadOnlyDocument struct {
	Document
}

func (m *ModifyReadOnlyDocument) open() ([]byte, error) {
	return m.data, nil
}

type ModifyWritableDocument struct {
	Document
}

func (m *ModifyWritableDocument) open() ([]byte, error) {
	return m.data, nil
}

func (m *ModifyWritableDocument) save(data []byte) error {
	m.data = data
	return nil
}

func processModifyDocument(d ModifyIDocument) error {
	data, err := d.open()
	if err != nil {
		return fmt.Errorf("opening document error: %w", err)
	}
	fmt.Println("Document opened with data:", string(data))
	return nil
}

func saveModifyDocument(d ModifyIWritableDocument, data []byte) error {
	err := d.save(data)
	if err != nil {
		return fmt.Errorf("saving document error: %w", err)
	}
	fmt.Println("Document saved successfully")
	return nil
}

func main() {
	fmt.Println("Before Liskov Substitution Principle")
	readOnlyDocument := ReadOnlyDocument{
		Document: Document{
			data:     []byte("Read-only document data"),
			fileName: "read-only.txt",
		},
	}

	writableDocument := WritableDocument{
		Document: Document{
			data:     []byte("Writable document data"),
			fileName: "writable.txt",
		},
	}

	fmt.Println("Process read-only document")
	if err := processDocument(&readOnlyDocument); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Process writable document")
	if err := processDocument(&writableDocument); err != nil {
		fmt.Println(err)
	}

	fmt.Println("After Liskov Substitution Principle")

	modifyReadOnlyDocument := ModifyReadOnlyDocument{
		Document: Document{
			data:     []byte("Modify read-only document data"),
			fileName: "modify-read-only.txt",
		},
	}

	modifyWritableDocument := ModifyWritableDocument{
		Document: Document{
			data:     []byte("Modify writable document data"),
			fileName: "modify-writable.txt",
		},
	}

	fmt.Println("Process modify read-only document")
	if err := processModifyDocument(&modifyReadOnlyDocument); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Process modify writable document")
	if err := processModifyDocument(&modifyWritableDocument); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Save modify writable document")
	if err := saveModifyDocument(&modifyWritableDocument, []byte("Modify writable document data saved")); err != nil {
		fmt.Println(err)
	}
}
