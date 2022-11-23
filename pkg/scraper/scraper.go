package scraper

import (
	"context"

	"github.com/chromedp/chromedp"
)

type Scraper struct {
	context context.Context
	cancel  context.CancelFunc
}

func New() *Scraper {
	ctx, cancel := chromedp.NewContext(context.Background())

	return &Scraper{
		context: ctx,
		cancel:  cancel,
	}
}

func (s *Scraper) Close() {
	s.cancel()
}
