package main

import (
	"bufio"
)

func ReadLine(reader *bufio.Reader) (line []byte, err error) {
	var buffer []byte
	isPrefix := true
	line = make([]byte, 0, 0)
	for isPrefix {
		if buffer, isPrefix, err = reader.ReadLine(); err != nil {
			return nil, err
		}
		line = append(line, buffer...)
	}
	return line, nil
}
