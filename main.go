package main

import (
	"bytes"
	"code.google.com/p/go.crypto/ssh"
	"github.com/datianshi/sshdownload/config"
	"io"
	"os"
)

func main() {
	sshConfig := config.ParseConfig("config.yml")
	config := &ssh.ClientConfig{
		User: sshConfig.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshConfig.Password),
		},
	}
	client, err := ssh.Dial("tcp", sshConfig.Host+":"+sshConfig.Port, config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	f, err := os.Create(sshConfig.File)
	defer f.Close()
	check(err)
	var b bytes.Buffer
	buf := make([]byte, 1024)
	session.Stdout = &b
	if err := session.Run(sshConfig.Cmd); err != nil {
		panic("Failed to run: " + err.Error())
	}
	for {
		n, err := b.Read(buf)
		if err == io.EOF {
			break
		}
		_, err = f.Write(buf[:n])
		check(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
