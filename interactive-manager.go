package main

import (
	"fmt"
	"gopkg.in/AlecAivazis/survey.v1"
)

// the questions to ask
var qs = []*survey.Question{
	{
		Name:      "name",
		Prompt:    &survey.Input{Message: "enter service name:"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "color",
		Prompt: &survey.Select{
			Message: "Choose a configuration:",
			Options: []string{"static", "proxy", "webproxy"},
			Default: "static",
		},
	},
	{
		Name:   "final",
		Prompt: &survey.Input{Message: "Start the service now?"},
	},
}

func main() {
	// the answers will be written to this struct
	answers := struct {
		Name   string // survey will match the question and field names
		Config string `service:"conf"` // or you can tag fields to match a specific name
		Final  int    // if the types don't match exactly, survey will try to convert for you
	}{}

	// questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s chose %s.", answers.Name, answers.Config)
}
