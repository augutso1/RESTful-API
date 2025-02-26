
A RESTful API built with Go to manage text posts and analyze sentiment. Powered by Gin, MongoDB, and a sprinkle of backend magic, Sentigo lets users submit posts, retrieve them, and get a basic sentiment score—think of it as a lightweight foundation for AI-driven systems.

## 🌟 Features
- **POST /posts**: Submit a text post (e.g., "I love coding") and get a JSON response with an ID and sentiment ("positive", "neutral").
- **GET /posts**: Retrieve all stored posts with their sentiments.
- **Tech Stack**: Go with Gin for fast HTTP handling, MongoDB for flexible storage.
- **Future-Ready**: Designed with extensibility in mind—plug in real AI models down the road!

## 🚀 Why Sentigo?
I built this to sharpen my backend skills while dipping into sentiment analysis—a stepping stone to AI-powered systems I’m passionate about. It’s a clean, functional API that shows off RESTful design, database integration, and a taste of engineering creativity.

## 🛠️ Getting Started

### Prerequisites
- Go 1.21+ (install via [golang.org](https://golang.org))
- MongoDB (local or [MongoDB Atlas](https://www.mongodb.com/cloud/atlas))
- Git (to clone this repo)

### Installation
1. Clone the repo:
   ```bash
   git clone https://github.com/yourusername/sentigo.git
   cd sentigo
