package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choices struct {
	cmd         string
	description string
	nextNode    *storyNode
	nextChoice  *choices
}

type storyNode struct {
	text        string
	storyChoice *choices
}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := choices{cmd, description, nextNode, nil}
	if node.storyChoice == nil {
		node.storyChoice = &choice
	} else {
		currentChoice := node.storyChoice
		for currentChoice.nextChoice != nil {
			currentChoice = currentChoice.nextChoice
		}
		currentChoice.nextChoice = &choice
	}
}

func (node *storyNode) render() {
	fmt.Println(node.text)
	currentChoice := node.storyChoice

	for currentChoice != nil {
		fmt.Println(currentChoice.cmd, ":", currentChoice.description)
		currentChoice = currentChoice.nextChoice
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	currentChoice := node.storyChoice
	for currentChoice != nil {
		if strings.ToLower(currentChoice.cmd) == strings.ToLower(cmd) {
			return currentChoice.nextNode
		}
		currentChoice = currentChoice.nextChoice
	}
	fmt.Println("Sorry, I didn't understand that")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.render()
	if node.storyChoice != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	start := storyNode{text: "text"}

	secRoom := storyNode{text: "second text"}
	secSecRoom := storyNode{text: "second version of second room"}

	thirdRoom := storyNode{text: "third text"}

	start.addChoice("S", "Go S", &secRoom)
	start.addChoice("SS", "Go SS", &secSecRoom)

	secRoom.addChoice("T", "go T", &thirdRoom)

	start.play()
}
