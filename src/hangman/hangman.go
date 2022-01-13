package main

import (
	"fmt"
)

var life int
var word string
var length int
var c []byte
var u []byte
var alpha []bool

func check(char byte, c []byte) bool {
	flag := false
	for i := 0; i < int(len(c)); i++ {
		if char == c[i] {
			flag = true
			break
		}
	}
	return flag
}

func print(u []byte) {
	for i := 0; i < len(u); i++ {
		fmt.Printf("%s ", string(u[i]))

	}
}

func main() {
	fmt.Println("Welcome to HangMan:")
	life = 4
	word = "hangman"
	c = []byte(word)
	length = len(word)
	u = make([]byte, length)
	alpha = make([]bool, 26)
	for i := 0; i < length; i++ {
		u[i] = byte('_')
	}
	print(u)
	temp := 0
	for life > 0 && temp < length {
		fmt.Println("Enter a character: ")
		var ip string
		fmt.Scan(&ip)
		char := []rune(ip)
		if alpha[int(char[0])-97] {
			fmt.Println("Already Entered")
			print(u)
			fmt.Println("Life left : ", (life + 1))
			continue
		}
		alpha[int(char[0])-97] = true
		for j := 0; j < length; j++ {
			if byte(char[0]) == c[j] && byte(char[0]) != u[j] {
				temp = temp + 1
			}
			if byte(char[0]) == c[j] {
				u[j] = c[j]
				alpha[int(char[0])-97] = true
			}
		}
		if (!check(byte(char[0]), c)) && alpha[int(char[0])-97] {
			life = life - 1
		} else if !check(byte(char[0]), c) {
			fmt.Println("Already Entered Wrong")
		}
		print(u)
		fmt.Println("Life left : ", (life + 1))
	}
	if life <= 0 {
		fmt.Println("Game Over Life Over word was ", word)
	} else if temp >= length {
		fmt.Println("You won the word was " + word)
	}
}
