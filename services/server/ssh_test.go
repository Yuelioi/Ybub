package server_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/crypto/ssh"
)

func TestCommand(t *testing.T) {
	server := struct {
		Host         string
		Port         int
		Username     string
		Password     string
		IdentityFile string
	}{
		Host:         "106.54.205.173",
		Port:         22,
		Username:     "root",
		Password:     "",              // 如果用密码登录就填这里
		IdentityFile: "~/.ssh/id_rsa", // 如果用私钥登录
	}

	client, err := getSSHClient(server)
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	stdout, _ := session.StdoutPipe()
	stderr, _ := session.StderrPipe()
	stdin, _ := session.StdinPipe()

	if err := session.Shell(); err != nil {
		log.Fatal(err)
	}

	// 运行命令示例
	fmt.Fprint(stdin, "cd ~\n")
	fmt.Fprint(stdin, "ls -l\n")
	fmt.Fprint(stdin, "exit\n")

	// 逐行打印 stdout
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println("OUT:", scanner.Text())
		}
	}()

	// 逐行打印 stderr
	scannerErr := bufio.NewScanner(stderr)
	for scannerErr.Scan() {
		fmt.Println("ERR:", scannerErr.Text())
	}

	if err := session.Wait(); err != nil {
		log.Fatal(err)
	}
}

func getSSHClient(server struct {
	Host         string
	Port         int
	Username     string
	Password     string
	IdentityFile string
}) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User:            server.Username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if server.IdentityFile != "" {
		identityPath := server.IdentityFile
		if len(identityPath) >= 2 && identityPath[:2] == "~/" {
			home, err := os.UserHomeDir()
			if err != nil {
				return nil, err
			}
			identityPath = filepath.Join(home, identityPath[2:])
		}

		key, err := os.ReadFile(identityPath)
		if err != nil {
			return nil, err
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return nil, err
		}
		config.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	} else if server.Password != "" {
		config.Auth = []ssh.AuthMethod{ssh.Password(server.Password)}
	} else {
		return nil, fmt.Errorf("必须提供密码或私钥")
	}

	return ssh.Dial("tcp", fmt.Sprintf("%s:%d", server.Host, server.Port), config)
}
