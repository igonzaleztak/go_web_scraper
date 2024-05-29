package scraper

// Scraper interface defines the methods that a scraper must implement.
type Scraper interface {
	Scrap() error // Scrap scraps the information from the API
	Print()       // Print formats the scraped data and prints it to console
}
