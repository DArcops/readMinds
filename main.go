package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/darcops/decisionTree/tree"
)

func main() {

	var arbol = tree.New()

	for {
		t := arbol

		for t.Lft != nil {
			if binaryAsk(fmt.Sprintf("tu animal, %v?\n", t.Value)) == "s" {
				t = t.Lft
			} else {
				t = t.Rght
			}
		}

		guess(t)

		if binaryAsk(fmt.Sprintf("quieres seguir jugando?\n")) == "n" {
			arbol.Save()
			fmt.Println("Adios hommie")
			break
		}
	}
}

func guess(t *tree.Tree) {
	if binaryAsk(fmt.Sprintf("A caso piensas en %v %v?\n", prefix(t.Value), t.Value)) == "s" {
		fmt.Println("Eso fue facil chavo")
		return
	}

	current := t.Value

	var ans string
	var info string
	fmt.Println("Que animal era?")
	ans = readLn()

	fmt.Printf("Que diferencia a %v %v de %v %v\n", prefix(t.Value), t.Value, prefix(ans), ans)
	info = readLn()

	t.Value = info

	fmt.Printf("%v %v , %v?\n", prefix(current), current, info)
	if readLn() == "s" {
		t.Lft = &tree.Tree{Value: current}
		t.Rght = &tree.Tree{Value: ans}
	} else {
		t.Rght = &tree.Tree{Value: current}
		t.Lft = &tree.Tree{Value: ans}
	}
}

func prefix(s string) string {
	if string(s[len(s)-1]) == "a" {
		return "una"
	}
	return "un"
}

func binaryAsk(s string) string {
	fmt.Println(s)
	return readLn()
}

func readLn() string {
	var line string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line = scanner.Text()
	return line
}

func printTree(t *tree.Tree) {
	fmt.Println("Root", t.Value)
	if t.Rght != nil && t.Lft != nil {
		fmt.Println("right -> ", t.Rght)
		fmt.Println("left -> ", t.Lft)
		printTree(t.Rght)
		printTree(t.Lft)
	}
}
