package client

import (
	"bytes"
	"log"
	"os/exec"
)

func execute(cmd *exec.Cmd) error {
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Println(stderr.String())
		return err
	}

	log.Println(out.String())
	return nil
}
