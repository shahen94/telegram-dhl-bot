package scraper

import (
	"fmt"

	"github.com/chromedp/chromedp"
)

const locationClass = "c-tracking-result--status-copy-message"
const containerClass = "c-tracking-result--status"

// GetLastLocation returns the channel which will be closed when the last location is found
func (s *Scraper) GetLastLocation(trackingNumber string) string {
	var location string

	defer s.Close()

	println("Getting location for", trackingNumber)

	err := chromedp.Run(
		s.context,
		chromedp.Navigate(fmt.Sprintf("https://www.dhl.com/us-en/home/tracking/tracking-express.html?submit=1&tracking-id=%s", trackingNumber)),
		chromedp.Text(fmt.Sprintf("h2.%s", locationClass), &location, chromedp.NodeVisible),
	)

	println("Location found:", location)

	if err != nil {
		return ""
	}

	return location
}
