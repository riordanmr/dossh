// Program to connect via SSH to a hosting provider and
// create a SQL dump of the Sedona Library holds database.
// Mark Riordan   2023-07-31
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func main() {
	// From https://linuxhint.com/golang-ssh-examples/
	// Create SSH configuration.
	hostKeyCallback, err := knownhosts.New("/Users/mrr/.ssh/known_hosts")
	if err != nil {
		log.Fatal(err)
	}
	sshpw := os.Getenv("SSHPW")
	if len(sshpw) == 0 {
		log.Fatal("You must set the environment variable SSHPW")
	}
	config := &ssh.ClientConfig{
		User: "u35710771",
		Auth: []ssh.AuthMethod{
			ssh.Password(sshpw),
		},
		HostKeyCallback: hostKeyCallback,
	}

	// Connect to the SSH host.
	conn, err := ssh.Dial("tcp", "home92543841.1and1-data.host:22", config)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Create an SSH session.
	session, err := conn.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Compute the command to be executed by the SSH host.
	sql_filenm := "holdsdb-" + time.Now().Format("2006-01-02") + ".sql"
	holdsdbpw := os.Getenv("HOLDSDBPW")
	if len(holdsdbpw) == 0 {
		log.Fatal("You must set the environment variable HOLDSDBPW")
	}
	cmd := "mysqldump --host=db5013161349.hosting-data.io --user=dbu4913091 --password=" +
		holdsdbpw + " dbs11045614 >volunteer/" + sql_filenm +
		" && ls -la volunteer"
	fmt.Println(cmd)

	// Send the command to the SSH host for execution.
	var buff bytes.Buffer
	session.Stdout = &buff
	if err := session.Run(cmd); err != nil {
		log.Fatal(err)
	}
	// Print the results.
	fmt.Println(buff.String())
}
