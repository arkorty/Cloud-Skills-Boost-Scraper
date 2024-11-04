# Scraper

A Go application that scrapes profile information from a list of URLs, processes user data, and exports it to JSON. This tool is designed for efficient data collection and structured output for further analysis or reporting.

---

### Features

- **Automated Web Scraping**: Collects profile data from multiple URLs with a single command.
- **CSV Input, JSON Output**: Accepts a CSV file with user information and outputs structured JSON data.
- **Data Validation**: Ensures only valid records are processed and logs any issues encountered.
- **Detailed User Profiles**: Extracts essential details such as completed assignments, badges, and counts.
- **Configurable Assignments List**: Allows customization of assignment tracking based on user needs.
- **Error Handling & Logging**: Provides clear logging of errors and successes for each profile processed.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/arkorty/Cloud-Skills-Boost-Scraper.git
   ```
2. Navigate to the project directory:
   ```bash
   cd Cloud-Skills-Boost-Scraper
   ```
3. Build the project:
   ```bash
   go build ./cmd/scraper
   ```

### Usage

To run the scraper, use the following command:

```bash
./scraper <input.csv> <output.json>
```

- `<input.csv>`: Path to the CSV file containing user names, emails, and profile URLs.
- `<output.json>`: Path where the JSON output will be saved.

Example:

```bash
./scraper input.csv output.json
```

### Contributing

We welcome contributions! Please fork this repository and submit a pull request if you have any improvements or fixes to suggest.

### License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
