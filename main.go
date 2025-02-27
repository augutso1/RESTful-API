package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Post é o modelo de dados que será salvo no banco de dados
type Post struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Content   string             `json:"content" binding:"required" bson:"content"`
	Sentiment string             `json:"sentiment" bson:"sentiment"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	Response  string             `json:"response" bson:"response"`
}

// PostRequest representa o corpo da requisição que será enviado para o servidor
type PostRequest struct {
	Content string `json:"content" binding:"required"`
}

type ollamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type ollamaResponse struct {
	Response string `json:"response"`
}

func callOllama(payload ollamaRequest) string {
	jsonData, _ := json.Marshal(payload)
	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Ollama error: %v", err)
		return "Error"
	}
	defer resp.Body.Close()
	var ollamaResp ollamaResponse
	json.NewDecoder(resp.Body).Decode(&ollamaResp)
	return ollamaResp.Response
}

func analyzeWithOllama(content string) (sentiment, response string) {
	sentimentPayload := ollamaRequest{
		Model:  "mistral:7b",
		Prompt: "Analyze the sentiment of this text: " + content,
		Stream: false,
	}

	sentimentResp := callOllama(sentimentPayload)

	// Response generation
	responsePayload := ollamaRequest{
		Model:  "mistral:7b",
		Prompt: "Respond briefly to this: " + content,
		Stream: false,
	}
	responseResp := callOllama(responsePayload)

	return sentimentResp, responseResp
}

var client *mongo.Client
var collection *mongo.Collection

func main() {
	// Configura a conexão com o MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// Verifica a conexão com o MongoDB
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Configura o banco de dados e a coleção
	database := client.Database("sentiment_db")
	collection = database.Collection("posts")

	// Cria o roteador do Gin
	router := gin.Default()
	router.Use(cors.Default())

	// Define as rotas
	router.POST("/posts", createPost)
	router.GET("/posts", getPosts)

	// Inicia o servidor
	log.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}

func createPost(c *gin.Context) {

	var request PostRequest
	// Associa o JSON da requisição ao struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cria um novo post
	post := Post{
		ID:        primitive.NewObjectID(),
		Content:   request.Content,
		CreatedAt: time.Now(),
	}

	// Analisa o sentimento e gera resposta usando Ollama
	post.Sentiment, post.Response = analyzeWithOllama(request.Content)

	// Insere o post no MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save post"})
		return
	}

	// Retorna o post salvo
	c.JSON(http.StatusCreated, post)
}

func getPosts(c *gin.Context) {
	var posts []Post

	// Obtém todos os posts do MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts"})
		return
	}
	defer cursor.Close(ctx)

	// Decodifica os documentos
	if err := cursor.All(ctx, &posts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode posts"})
		return
	}

	// Retorna os posts
	c.JSON(http.StatusOK, posts)
}
