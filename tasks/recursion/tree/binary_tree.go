package main

import "fmt"

// напечатает все возможные комбинации 0 и 1
func binaryTree(n int, prefix string) {
	if n == -1 {
		return
	}

	if prefix != "" {
		fmt.Println(prefix)
	}

	n--

	binaryTree(n, prefix+"0")
	binaryTree(n, prefix+"1")
}

func main() {
	binaryTree(3, "")
}
