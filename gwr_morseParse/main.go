package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gen2brain/beeep"
)

type MorseCharacter struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Value string `json:"value"`
}

type MorseCodeTemplate struct {
	Version    string           `json:"version"`
	Characters []MorseCharacter `json:"characters"`
}

func getCode(char string) MorseCharacter {

	var Err MorseCharacter

	content, err := os.ReadFile("intMorse.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var morseSet MorseCodeTemplate
	err = json.Unmarshal(content, &morseSet)
	if err != nil {
		log.Fatal("Error usring Unmarshal(): ", err)
	}

	for _, ch := range morseSet.Characters {
		if ch.Value == strings.ToLower(char) {
			return ch
		}
	}

	Err.Name = "Error"
	Err.Code = ""
	Err.Value = "No Character " + char + "found"

	return Err
}

func dash() {
	err := beeep.Beep(440.0, 600)
	if err != nil {
		log.Panic(err)
	}
}
func dot() {
	err := beeep.Beep(440.0, 200)
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	phrase := "0505"
	var mphrase []string
	for _, char := range phrase {
		c := strings.ToLower(string(char))

		mc := getCode(c)
		if mc.Name == "Error" {
			//mc.Value = "#"
			mc.Code = "E"
			//log.Fatal(mc.Value)
		}

		mphrase = append(mphrase, mc.Code)
	}

	fmt.Println(mphrase)

	for _, char := range mphrase {
		for i, _ := range char {
			switch string(char[i]) {
			case ".":
				dot()
				fmt.Println(".")
			case "-":
				dash()
				fmt.Println("-")

			}
		}

	}
	beeep.Notify("This is my title", "This is the message", "assets/gopher.jpg")

}
