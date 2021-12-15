package steps

import "github.com/cucumber/godog"

func (c *AreaProfileComponent) RegisterSteps(ctx *godog.ScenarioContext) {
	c.apiFeature.RegisterSteps(ctx)
}
