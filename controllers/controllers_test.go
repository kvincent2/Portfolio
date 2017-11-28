package controllers

import (
	"testing"

	"github.com/franela/goblin"
)

func TestStatus(t *testing.T) {
	// Create a goblin instance
	g := goblin.Goblin(t)

	var tests = []struct {
		ExpectedCode    int
		TestDescription string
	}{
		{
			ExpectedCode: 200,
			TestDescription: "Should return 200 with status OK",
		},
	}

	for _, test := range tests {
		g.Describe("GetStatus", func() {
			g.It(test.TestDescription, func() {

				code, _, _ := GetStatus()
				// validate code is what is expected
				g.Assert(code).Equal(test.ExpectedCode)
			})
		})
	}
}

func TestHello(t *testing.T) {
	// Create a goblin instance
	g := goblin.Goblin(t)

	hello := make(map[string]string)
	hello["name"] = "Kristina Vincent"

	tests := []struct {
		ExpectedCode    int
		ExpectedBody    map[string]string
		TestDescription string
	}{
		{
			ExpectedCode:    200,
			ExpectedBody:    hello,
			TestDescription: "Should return 200 with Hello World",
		},
	}

	for _, test := range tests {
		g.Describe("GetHello", func() {
			g.It(test.TestDescription, func() {

				code, body, _ := GetHello()
				// validate code is what is expected
				g.Assert(code).Equal(test.ExpectedCode)
				// validate body is what is expected
				g.Assert(body).Equal(test.ExpectedBody)
			})
		})
	}
}

func TestGetProjects(t *testing.T) {
	// Create a goblin instance
	g := goblin.Goblin(t)

	var tests = []struct {
		TestDescription string
	}{
		{
			TestDescription: "Should return repos",
		},
	}

	for _, test := range tests {
		g.Describe("GetProjects", func() {
			g.It(test.TestDescription, func() {

				repos := GetProjects()
				// validate code is what is expected
				g.Assert(len(repos)).Equal(5)
			})
		})
	}
}
