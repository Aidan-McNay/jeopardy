//========================================================================
// file.go
//========================================================================
// Functions for interacting with files
//
// Heavily inspired by: https://medium.com/@matryer/golang-advent-calendar-day-eleven-persisting-go-objects-to-disk-7caf1ee3d11d
//
// Author: Aidan McNay
// Date: May 31st, 2024

package file

import (
	"bytes"
	"encoding/json"
	"io"
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

func Save(fileWriter io.WriteCloser, v interface{}) error {
	file_lock.Lock()
	defer file_lock.Unlock()
	defer fileWriter.Close()

	r, err := marshal(v)
	if err != nil {
		return err
	}

	_, err = io.Copy(fileWriter, r)
	return err
}

func Load(fileReader io.ReadCloser, v interface{}) error {
	file_lock.Lock()
	defer file_lock.Unlock()
	defer fileReader.Close()

	return unmarshal(fileReader, v)
}
