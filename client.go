package main

import (
	"encoding/json"
	"net"
	"os"
)

func RunClient(sock string) error {
	c, err := net.Dial("unix", sock)
	defer c.Close()
	if err != nil {
		return err
	}

	m := message{}
	m.DrvPath = os.Getenv("DRV_PATH")
	m.OutPaths = os.Getenv("OUT_PATHS")

	payload, err := json.Marshal(m)
	if err != nil {
		return err
	}

	_, err = c.Write(payload)
	if err != nil {
		return err
	}

	return nil
}
