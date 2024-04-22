### Web Scraper

Web Scraper is an application engineered to navigate through the web and extract text from a website based on user-defined keywords. This tool streamlines the process of gathering information from the internet, enabling users to compile data efficiently. In addition to its core functionality, Web Scraper is enhanced with powerful text manipulation capabilities using advanced AI technologies provided by Google's Gemini API. These functionalities include summarization, sentiment analysis, and translation, all of which can be performed on the extracted text. Additionally, users have the option to save the results in a file, making data management and utilization more convenient.

### Architecture & Design

The architecture and design of the Web Scraper application incorporate a robust and efficient tech stack that ensures optimal performance and scalability. The backend is built using Go (Golang), known for its high performance and efficiency in handling concurrent tasks. For parsing and manipulating HTML documents, the backend utilizes goquery, a powerful library that simplifies DOM traversal and manipulation in Go.

On the frontend, Vue.js is employed, which provides a reactive and composable user interface architecture. This makes it particularly suitable for building dynamic web applications that require real-time updates, such as a web scraping tool.

For data storage, the application uses Amazon DynamoDB, a highly scalable NoSQL database service. This choice ensures quick access to scraped data results and reliable performance under varying loads. Additionally, Redis is integrated into the system for caching purposes. It employs a Read-Through and Write-Through Caching strategy. This approach ensures that data read requests first check the cache; if the requested data is not found, it is retrieved from the primary store, cached for future access, and then returned to the user. Write operations, on the other hand, simultaneously update the data in the cache and the primary store. This significantly enhances the application's responsiveness and efficiency by reducing the load on the database and speeding up data retrieval processes.

### Features

- _Web Page Scraping:_ Allows users to scrape web pages for information based on predefined keywords, efficiently extracting relevant data;
- _Text Summarization_: Automatically condenses the extracted text into a concise summary, making it easier to quickly grasp the main points of the content;
- _Sentiment Analysis_: Analyzes the emotional tone of the extracted text, helping users understand the general sentiment expressed within the content;
- _Text Translation_: Translates the scraped text into various languages, facilitating broader accessibility and understanding;
- _File Output_: Saves both the original and manipulated scraped data in a file, enabling easy storage and further processing for various uses;

### Installation

To run the application locally, execute the following command:

- go run .
