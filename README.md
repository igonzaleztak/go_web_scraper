Guidelines project:

A golang web scraper designed for the Hacker News webpage. This scraper will take the first 30 entries from the Hacker News frontpage (API: https://github.com/HackerNews/API) and then it will sort them in the following manner:

    First come entries with more than 5 words in their title
    Long title entries will then be sorted by number of comments
    Short title entries will be sorted by number of points

Then sorted entries will be printed out by console output.
Candidates are required to fill a README in ENGLISH explaining their thought process while completing the assignment, discussing key technical points, blocking points, etc. They should also explain test coverage and how to test the code.

API - HACKER WEB
Add testing

- PLUS: Try scrapping different new sources (Hint: Use goroutines)
- PLUS: Automation (Testing, running, publishing) with scripting, Make, Taskfile,...