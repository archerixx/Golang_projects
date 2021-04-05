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

	//introduction
	fmt.Println("Welcome\n Please enter your 'Character' name: ")
	scanner.Scan()
	Character := scanner.Text()
	fmt.Println("Hi", Character, ", I hope you will enjoy this short story game and maybe complete it few times to explore other paths\n\n ")

	//main story plot
	Plot := "After a long night of studying, blankly starting at book, you finally decided to go to bed. He was researching on how to build his first video game, but task deemed to be too much for just one night." +
		"\nAfter few minutes on phone, you fell asleep" +
		"Suddenly he was with his friends, geather at one place next to huge stone, it looked to him like they were waiting for someone"

	StoryStart := storyNode{text: Plot}

	Chapter1_first_C1 := storyNode{text: "\"Healer as always\", said a Armored giant who was sitting in front"}
	Chapter1_first_C2 := storyNode{text: "You looked around and saw group of six people talking about their day. Someone of them wore heavy armored playes, others were in regular light armor made of mostly leather. At the moment, all that seemed normal to you"}

	Chapter1_second_C := storyNode{text: "\"Healer as always\", said a Armored giant who was sitting in front"}

	Chapter1_third_C1 := storyNode{text: ""}
	Chapter1_third_C2 := storyNode{text: "a"}

	StoryStart.addChoice("1", "Hey,  who are we waiting for", &Chapter1_first_C1)
	StoryStart.addChoice("2", "Explore around", &Chapter1_first_C2)

	Chapter1_first_C2.addChoice("1", "After few moments, you decided to ask group in front of you: \"Hey,  who are we waiting for\"", &Chapter1_second_C)

	Chapter1_second_C.addChoice("1", "s", &Chapter1_third_C1)
	Chapter1_second_C.addChoice("2", "a", &Chapter1_third_C2)

	StoryStart.play()
}
