package flag

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type flagHandler struct{}

func NewFlagHandler() *flagHandler {
	return &flagHandler{}
}

type Body struct {
	Flag string `json:"flag"`
}

func (fh *flagHandler) PutFlag(w http.ResponseWriter, r *http.Request) {
	var b Body
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		w.Write([]byte("error occured: " + err.Error()))
		return
	}

	cwd, err := os.Getwd()
	if err != nil {
		w.Write([]byte("error occured: " + err.Error()))
		return
	}

	path := filepath.Join(cwd, "flag.txt")
	newFilepath := filepath.FromSlash(path)
	file, err := os.Create(newFilepath)
	if err != nil {
		w.Write([]byte("error occured: " + err.Error()))
		return
	}
	defer file.Close()

	fw := bufio.NewWriter(file)
	fw.WriteString(b.Flag)
	fw.Flush()
}

func (fh *flagHandler) GetFlag(w http.ResponseWriter, r *http.Request) {
	cwd, err := os.Getwd()
	if err != nil {
		w.Write([]byte("error occured: " + err.Error()))
		return
	}

	path := filepath.Join(cwd, "flag.txt")
	newFilepath := filepath.FromSlash(path)

	bytes, err := ioutil.ReadFile(newFilepath)

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
