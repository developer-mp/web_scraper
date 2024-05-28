### Web Scraper

Web Scraper is an application engineered to navigate through the web and extract text from a website based on user-defined keywords. This tool streamlines the process of gathering information from the internet, enabling users to compile data efficiently. In addition to its core functionality, Web Scraper is enhanced with powerful text manipulation capabilities using advanced AI technologies provided by Google's Gemini API. These functionalities include summarization, sentiment analysis, and translation, all of which can be performed on the extracted text. Additionally, users have the option to save the results in a file, making data management and utilization more convenient.

### Architecture & Design

The architecture and design of the Web Scraper application incorporate a robust and efficient tech stack that ensures optimal performance and scalability:

- **Backend Technology**: Built using Go (Golang), known for its high performance and efficiency in handling concurrent tasks.
- **HTML Parsing**: Utilizes goquery, a powerful library that simplifies DOM traversal and manipulation in Go, for parsing and manipulating HTML documents.
- **Frontend Framework**: Employs Vue.js for the frontend, providing a reactive and composable user interface architecture suitable for dynamic web applications.
- **Data Storage**: Relies on Amazon DynamoDB, a highly scalable NoSQL database service, ensuring quick access to scraped data results and reliable performance under varying loads.
- **Caching Strategy**: Integrates Redis for caching purposes, employing a Read-Through and Write-Through Caching strategy to enhance responsiveness and efficiency.
- **Asynchronous Programming**: Utilizes asynchronous programming techniques to perform scraping tasks concurrently, maximizing resource utilization and preventing blocking.

Additionally, the application includes several key mechanisms for enhanced functionality and reliability:

- **Data Deduplication**: Implements mechanisms to detect and skip redundant information, avoiding scraping duplicate data.
- **Rate Limiting and Throttling**: Integrates mechanisms to prevent overwhelming the servers of the target website, mitigating the risk of being banned.
- **Proxy Rotation**: Implements strategic distribution of requests across multiple IP addresses to avoid IP bans or detection.

### Features

- _Web Page Scraping:_ Allows users to scrape web pages for information based on predefined keywords, efficiently extracting relevant data;
- _Text Summarization_: Automatically condenses the extracted text into a concise summary, making it easier to quickly grasp the main points of the content;
- _Sentiment Analysis_: Analyzes the emotional tone of the extracted text, helping users understand the general sentiment expressed within the content;
- _Text Translation_: Translates the scraped text into various languages, facilitating broader accessibility and understanding;
- _File Output_: Saves both the original and manipulated scraped data in a file, enabling easy storage and further processing for various uses;

### Installation

To run the application locally, execute the following command:

- go run .
