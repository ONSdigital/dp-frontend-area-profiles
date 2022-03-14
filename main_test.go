package main

import (
	"flag"
	"os"
	"testing"

	componenttest "github.com/ONSdigital/dp-component-test"
	feature "github.com/ONSdigital/dp-frontend-area-profiles/features"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/stretchr/testify/assert"
)

var componentFlag = flag.Bool("component", false, "perform component tests")

func InitializeScenario(ctx *godog.ScenarioContext) {
	component, _ := feature.NewAreaProfilesComponent()
	uiFeature := componenttest.NewUIFeature("http://" + component.Config.SiteDomain + component.Config.BindAddr)

	ctx.BeforeScenario(func(*godog.Scenario) {
		uiFeature.Reset()
	})

	ctx.AfterScenario(func(*godog.Scenario, error) {
		uiFeature.Close()
	})

	uiFeature.RegisterSteps(ctx)
	// Start page custom steps
	// p:nth-child(1) > a:nth-child(1)
	ctx.Step(`^the page heading should be "([^"]*)"`, selectedContentShouldExist(uiFeature, "[data-test='h1']"))
	ctx.Step(`^the first paragraph should have a link of "([^"]*)"`, selectedContentShouldExist(uiFeature, "[data-test='p1'] > a:nth-child(1)"))
	ctx.Step(`^the second paragraph should have a link of "([^"]*)"`, selectedContentShouldExist(uiFeature, "[data-test='p1'] > a:nth-child(1)")) // bug in chromedp
	ctx.Step(`^the country section sub heading is "([^"]*)"`, selectedContentShouldExist(uiFeature, "h2"))
	ctx.Step(`^the country section first paragraph contains link with text "([^"]*)"`, selectedContentShouldExist(uiFeature, "div:nth-child(3) > p:nth-child(1) > a:nth-child(1)"))
	ctx.Step(`^the country section second paragraph contains link with text "([^"]*)"`, selectedContentShouldExist(uiFeature, "div:nth-child(3) > p:nth-child(1) > a:nth-child(1)")) // bug in chromedp
	// Area page custom steps
	ctx.Step(`^the relations sub heading should be "([^"]*)"`, selectedContentShouldExist(uiFeature, "[data-test='h2Relations']"))
	ctx.Step(`^the relations sections should have (\d+) external links$`, sectionShouldHaveNthElements(uiFeature, "[data-test='relationLinks'] > div > div > a"))
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
	})
}

func TestComponent(t *testing.T) {
	if *componentFlag {
		status := 0

		var opts = godog.Options{
			Output: colors.Colored(os.Stdout),
			Format: "pretty",
			Paths:  flag.Args(),
		}

		status = godog.TestSuite{
			Name:                 "feature_tests",
			ScenarioInitializer:  InitializeScenario,
			TestSuiteInitializer: InitializeTestSuite,
			Options:              &opts,
		}.Run()

		if status > 0 {
			t.Fail()
		}
	} else {
		t.Skip("component flag required to run component tests")
	}
}

// -----------------------------------------
// Step helper functions
func sectionShouldHaveNthElements(f *componenttest.UIFeature, elementSelector string) func(int) error {
	return func(expectedLength int) error {
		var nodes []*cdp.Node
		err := chromedp.Run(f.Chrome.Ctx,
			f.RunWithTimeOut(f.WaitTimeOut, chromedp.Tasks{
				chromedp.Nodes(elementSelector, &nodes, chromedp.ByQueryAll),
			}))
		assert.Nil(f, err)
		if err != nil {
			return f.StepError()
		}
		assert.Equal(f, expectedLength, len(nodes))
		return f.StepError()
	}
}

func selectedContentShouldExist(f *componenttest.UIFeature, elementSelector string) func(string) error {
	return func(expectedContent string) error {
		var actualContent string
		err := chromedp.Run(f.Chrome.Ctx,
			f.RunWithTimeOut(f.WaitTimeOut, chromedp.Tasks{
				chromedp.Text(elementSelector, &actualContent, chromedp.NodeVisible),
			}),
		)
		if err != nil {
			return err
		}
		assert.Equal(f, expectedContent, actualContent)
		return nil
	}
}
