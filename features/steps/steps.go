package steps

import (
	"github.com/cucumber/godog"
)

func (c *AreaProfileComponent) RegisterSteps(ctx *godog.ScenarioContext) {
	c.uiFeature.RegisterSteps(ctx)

	//ctx.Step(`^I navigate to "([^"]*)"`, c.iNavigateTo)
	//ctx.Step(`^element "([^"]*)" should be visible$`, c.elementShouldBeVisible)
	//ctx.Step(`^the beta phase banner should be visible$`, c.theBetaBannerShouldBeVisible)
	//ctx.Step(`^the improve this page banner should be visible$`, c.theImproveThisPageBannerShouldBeVisible)
	//ctx.Step(`^the page should have the following content$`, c.thePageShouldHaveTheFollowingContent)
	//ctx.Step(`^should match snapshot "([^"]*)"$`, c.shouldMatchSnapshot)
}

//func (c *AreaProfileComponent) iNavigateTo(route string) error {
//	err := chromedp.Run(c.chrome.ctx,
//		chromedp.Navigate(c.BaseURL() + route),
//	)
//	if err != nil {
//		return c.apiFeature.StepError()
//	}
//
//	return nil
//}
//
//func (c *AreaProfileComponent) elementShouldBeVisible(elementSelector string) error {
//	err := chromedp.Run(c.chrome.ctx,
//		c.RunWithTimeOut(&c.chrome.ctx, c.waitTimeOut, chromedp.Tasks{
//			chromedp.WaitVisible(elementSelector),
//		}),
//	)
//	assert.Nil(c.apiFeature, err)
//
//	return c.apiFeature.StepError()
//}
//
//func (c *AreaProfileComponent) theBetaBannerShouldBeVisible() error {
//	return c.elementShouldBeVisible(".ons-phase-banner")
//}
//
//func (c *AreaProfileComponent) theImproveThisPageBannerShouldBeVisible() error {
//	return c.elementShouldBeVisible(".improve-this-page")
//}
//
//func (c *AreaProfileComponent) thePageShouldHaveTheFollowingContent(expectedAPIResponse *godog.DocString) error {
//	var contentElements map[string]string
//
//	err := json.Unmarshal([]byte(expectedAPIResponse.Content), &contentElements)
//	if err != nil {
//		return err
//	}
//
//	for selector, expectedContent := range contentElements {
//		var actualContent string
//		err = chromedp.Run(c.chrome.ctx,
//			c.RunWithTimeOut(&c.chrome.ctx, c.waitTimeOut, chromedp.Tasks{
//				chromedp.Text(selector, &actualContent, chromedp.NodeVisible),
//			}),
//		)
//
//		if err != nil {
//			return err
//		}
//
//		assert.Equal(c.apiFeature, expectedContent, actualContent)
//	}
//
//	return c.apiFeature.StepError()
//}
//
//func (c *AreaProfileComponent) shouldMatchSnapshot(snapshotName string) error {
//	var buf []byte
//	err := chromedp.Run(c.chrome.ctx,
//		c.RunWithTimeOut(&c.chrome.ctx, c.waitTimeOut, chromedp.Tasks{
//			chromedp.FullScreenshot(&buf, 90),
//		}),
//	)
//	if err != nil {
//		return err
//	}
//
//	err = cupaloy.Snapshot(c.apiFeature, buf)
//	if err != nil {
//		return err
//	}
//
//	return c.apiFeature.StepError()
//}
