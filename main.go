package main

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// O post é o modelo de dados que será salvo no banco de dados
type Post struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Content   string             `json:"content" binding:"required" bson:"content"`
	Sentiment string             `json:"sentiment" bson:"sentiment"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

// PostRequest representa o corpo da requisição que será enviado para o servidor
type PostRequest struct {
	Content string `json:"content" binding:"required"`
}

var collection *mongo.Collection

func main() {
	// Configura a conexão com o MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
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
		Sentiment: analyzeSentiment(request.Content),
		CreatedAt: time.Now(),
	}

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

func analyzeSentiment(content string) string {
	// Análise de sentimento simples baseada em palavras-chave
	positiveWords := []string{
		"adoro",
		"amo",
		"bom",
		"ótimo",
		"excelente",
		"feliz",
		"gosto",
		"legal",
		"maravilhoso",
		"perfeito",
	}

	negativeWords := []string{
		"odeio",
		"péssimo",
		"ruim",
		"terrível",
		"horrível",
		"triste",
		"detesto",
		"chato",
		"irritante",
		"frustrado",
	}

	contentLower := strings.ToLower(content)

	for _, word := range positiveWords {
		if strings.Contains(contentLower, word) {
			return "positive"
		}
	}

	for _, word := range negativeWords {
		if strings.Contains(contentLower, word) {
			return "negative"
		}
	}

	return "neutral"
}
