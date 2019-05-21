package tree

import (
	"encoding/json"
	"io/ioutil"
)

type Tree struct {
	Lft   *Tree
	Rght  *Tree
	Value string
}

func New() *Tree {
	var t = new(Tree)
	data, err := ioutil.ReadFile("train.json")
	if err != nil {
		t.Value = "leon"
		return t
	}
	json.Unmarshal(data, t)
	return t
}

func (t *Tree) Save() {
	data, _ := json.Marshal(t)
	ioutil.WriteFile("train.json", data, 0777)
}
