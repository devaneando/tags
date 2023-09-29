package model

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"the_dudes_age,omitempty"`
}

func (p *Person) Load(fn string) error {
	f, err := os.Open(fn)
	if err != nil {
		return err
	}

	defer f.Close()
	return p.ReadFrom(f)
}

func (p *Person) ReadFrom(r io.Reader) error {
	dec := json.NewDecoder(r)
	err := dec.Decode(p)
	return err
}

func (p Person) Save(fn string) error {
	f, err := os.Create(fn)
	if err != nil {
		return err
	}

	defer f.Close()
	return p.WriteTo(f)
}

func (p Person) WriteTo(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(p)
}

func (p Person) Check() error {
	if p.Name == "" {
		return errors.New("Bad name")
	}
	return nil
}
