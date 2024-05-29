package scraper

// Scraper interface defines the methods that a scraper must implement.
type Scraper interface {
	Scrap() error
	Print()
}
