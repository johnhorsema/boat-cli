package main

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
    "boat"
)

func main() {
	var px boat.Rule

	validateRule := func(input string) error {
		parsed, err := boat.ParseRule(input)
		if err != nil {
			return errors.New("Invalid request")
		}
		px = parsed
		_ = px
		return nil
	}

	validateInput := func(input string) error {
		pass, err := px.Eval(input)
		if pass != true || err != nil {
			return errors.New("Invalid request")
		}
		return nil
	}

	getrule := promptui.Prompt{
		Label:    "rule",
		Validate: validateRule,
	}

	getinput := promptui.Prompt{
		Label:    "boat",
		Validate: validateInput,
	}

	resultRule, err := getrule.Run()

	if err != nil {
		fmt.Printf("Rule failed %v\n", err)
		return
	}

	_ = resultRule

	resultInput, err := getinput.Run()

	if err != nil {
		fmt.Printf("Value failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", resultInput)
}