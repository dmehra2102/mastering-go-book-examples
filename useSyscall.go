package main

import (
	"fmt"
	"syscall"
)

func SyscallUsage() {
	message := []byte{'H', 'e', 'l', 'l', 'o', '!', '\n'}
	syscall.Write(1, message)

	pid, _, _ := syscall.SyscallN(39, 0, 0, 0)
	fmt.Println("My pid is", pid)
	uid, _, _ := syscall.SyscallN(24, 0, 0, 0)
	fmt.Println("User ID:", uid)

}
