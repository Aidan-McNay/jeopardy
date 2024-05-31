//========================================================================
// file.go
//========================================================================
// Functions for interacting with files
//
// Heavily inspired by: https://medium.com/@matryer/golang-advent-calendar-day-eleven-persisting-go-objects-to-disk-7caf1ee3d11d
//
// Author: Aidan McNay
// Date: May 31st, 2024

package storage

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"sync"
)

//------------------------------------------------------------------------
// Marshalling
//------------------------------------------------------------------------
// Functions for translating objects into byte streams

func marshal(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func unmarshal(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

//------------------------------------------------------------------------
// Loading and Saving Objects
//------------------------------------------------------------------------

var file_lock sync.Mutex

func Save(path string, v interface{}) error {
	file_lock.Lock()
	defer file_lock.Unlock()

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	r, err := marshal(v)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, r)
	return err
}

func Load(path string, v interface{}) error {
	file_lock.Lock()
	defer file_lock.Unlock()

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return unmarshal(f, v)
}
