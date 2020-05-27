package main

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
        "boat"
)

func main() {
	validate := func(input string) error {
		px, err := boat.ParseRule("hello")
		pass, err := px.Eval(input)
		if pass != true || err != nil {
			return errors.New("Invalid request")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "boat",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}
