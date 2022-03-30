package utils

import "github.com/ONSdigital/dp-frontend-area-profiles/mapper"

const (
	ErrorTitleFor404       = "404 - The webpage you are requesting does not exist on the site"
	ErrorDescriptionFor404 = `<p> The page may have been moved, updated or deleted or you may have typed the web address incorrectly, please check the url and spelling. Alternatively, please try the <a href="#nav-search">search</a>, or return to the <a href="/" title="Our homepage" target="_self">homepage</a>.</p>`
	ErrorTitleFor500       = "Sorry, there is a problem with the service"
	ErrorDescriptionFor500 = `<p>Please try again later.</p>
	<p>If the problem persists, check our <a href="https://twitter.com/ONS">twitter</a> feed for updates.</p>
	<p>Our scheduled publications can also be found on our <ahref="https://backup.ons.gov.uk/">backup site</a>.</p>`
)

// SetErrorDetails returns errors details for 404 status code
// or else returns error message details as for a 500 status code.
func SetErrorDetails(status int, errorDetails *mapper.ErrorDetails) {
	if status == 404 {
		errorDetails.Title = ErrorTitleFor404
		errorDetails.Description = ErrorDescriptionFor404
	} else {
		errorDetails.Title = ErrorTitleFor500
		errorDetails.Description = ErrorDescriptionFor500
	}
}
