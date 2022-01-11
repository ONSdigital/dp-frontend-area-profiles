package main

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/ONSdigital/dp-frontend-area-profiles/features/steps"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var componentFlag = flag.Bool("component", false, "perform component tests")

type ComponentTest struct {
}

func (f *ComponentTest) InitializeScenario(ctx *godog.ScenarioContext) {
	component, err := steps.NewAreaProfilesComponent()
	if err != nil {
		log.Fatal(context.Background(), "fatal error initialising a test scenario", err)
		os.Exit(1)
	}

	ctx.BeforeScenario(func(*godog.Scenario) {
		component.Reset()
	})

	ctx.AfterScenario(func(*godog.Scenario, error) {
		if err = component.Close(); err != nil {
			log.Warn(context.Background(), "error closing identity component", log.FormatErrors([]error{err}))
		}
	})

	component.RegisterSteps(ctx)
}

func (f *ComponentTest) InitializeTestSuite(ctx *godog.TestSuiteContext) {}

func TestComponent(t *testing.T) {
	if *componentFlag {
		status := 0

		var opts = godog.Options{
			Output: colors.Colored(os.Stdout),
			Format: "pretty",
			Paths:  flag.Args(),
		}

		f := &ComponentTest{}

		status = godog.TestSuite{
			Name:                 "feature_tests",
			ScenarioInitializer:  f.InitializeScenario,
			TestSuiteInitializer: f.InitializeTestSuite,
			Options:              &opts,
		}.Run()

		if status > 0 {
			t.Fail()
		}
	} else {
		t.Skip("component flag required to run component tests")
	}
}
