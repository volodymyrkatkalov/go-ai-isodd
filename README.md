# Go AI IsOdd 🚀

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/volodymyrkatkalov/go-ai-isodd)](https://goreportcard.com/report/github.com/volodymyrkatkalov/go-ai-isodd)

Welcome to **Go AI IsOdd**, a whimsical Go package that leverages the power of artificial intelligence (ChatGPT) to answer a timeless question: *Is this number even or odd?* Why do the math yourself when you can ask an AI? 🤖

---

## ✨ Features

- **AI-Powered**: Queries ChatGPT via OpenAI's API to determine if a number is even or odd.
- **Simple API**: Easy-to-use client with a single method: `isodd(number int)`.
- **Robust**: Includes error handling and testing with a mock server.
- **Fancy**: Because who doesn’t love over-engineering a simple problem?

---

## 🛠️ Installation

Get the package with:
```bash
go get github.com/volodymyrkatkalov/go-ai-isodd
```

You'll need an OpenAI API key. Set it as an environment variable:
```bash
export OPENAI_API_KEY='your-api-key-here'
```

## 🚀 Usage
Here's how to use it in your Go code:
```go

package main

import (
    "fmt"
    "os"
    "github.com/yourusername/go-ai-isodd"
)

func main() {
    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        fmt.Println("Please set OPENAI_API_KEY environment variable")
        return
    }

    client := isodd.NewChatGPTClient(apiKey)
    result, err := client.IsOdd(42)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Is 42 odd? %v\n", result)
}
```

**Output:**

```bash
ChatGPT says: 42 is yes
Is 42 odd? true
```

Run it:
```bash
go run main.go
```

## 🧪 Testing
The package includes unit tests with a mock server to simulate ChatGPT responses:
```bash
go test
```

Test cases cover positive, negative, and zero inputs to ensure reliability.
## 📋 Requirements
* Go: 1.21 or higher
* OpenAI API Key: Obtainable from platform.openai.com
* Dependencies: 
    * github.com/sashabaranov/go-openai (automatically fetched with go get)

## 🎨 Example Output
Asking ChatGPT about various numbers:
* `client.IsOdd(2)` → `ChatGPT says: 2 is no` → `false`
* `client.IsOdd(3)` → `ChatGPT says: 3 is yes` → `true`
* `client.IsOdd(0)` → `ChatGPT says: 0 is no` → `false`

## 🤝 Contributing
Contributions are welcome! Here’s how to get involved:

* Fork the repository
* Create a feature branch (git checkout -b feature/awesome-thing)
* Commit your changes (git commit -m "Add awesome thing")
* Push to the branch (git push origin feature/awesome-thing)
* Open a Pull Request

Please include tests and follow Go conventions.

## 📜 License
This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgements
* OpenAI for the ChatGPT API
* sashabaranov/go-openai for the excellent Go client
* You, for checking out this awesome project!

## ⭐ Star this repo if you find it amusing or useful! ⭐
Created with ❤️ by volodymyrkatkalov
