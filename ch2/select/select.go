package main

import (
	"syscall"
	"fmt"
)

int main() {
	ret, err := syscall.Select()
	if err != nil {
		fmt.Println("hogehoge")
	}

	if (ret == -1) {

	} else if (!ret) {
		
	}
}