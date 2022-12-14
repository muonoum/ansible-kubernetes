package node_hash

import (
	"encoding/hex"
	"fmt"
	"hash/fnv"
)

type Command struct {
	Value string `arg`
}

func (cmd Command) Run() error {
	hasher := fnv.New32a()
	hasher.Write([]byte(cmd.Value))
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
	return nil
}
