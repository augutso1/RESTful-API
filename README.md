# Sentigo

Uma API RESTful construída com Go para gerenciar postagens de texto e analisar sentimentos. Desenvolvida com Gin, MongoDB e um toque de mágica backend, a Sentigo permite que usuários enviem postagens, recuperem-nas e obtenham uma pontuação básica de sentimento — pense nisso como uma base leve para sistemas baseados em IA. ("Sentigo" vem de "sentiment" + "Go", unindo sentimento e a linguagem Go!)

## 🌟 Funcionalidades
- **POST /posts**: Envie uma postagem de texto (ex.: "Adoro programar") e receba uma resposta JSON com ID e sentimento ("positivo", "neutro").
- **GET /posts**: Recupere todas as postagens salvas com seus sentimentos.
- **Tecnologias**: Go com Gin para lidar com HTTP de forma rápida, MongoDB para armazenamento flexível.
- **Preparada para o Futuro**: Projetada com extensibilidade em mente — conecte modelos de IA reais no futuro!

## 🚀 Por que a Sentigo?
Criei esse projeto para aprimorar minhas habilidades de backend enquanto exploro análise de sentimento — um passo inicial rumo a sistemas baseados em IA, que me fascinam. É uma API limpa e funcional que demonstra design RESTful, integração com banco de dados e um toque de criatividade em engenharia.

## 🛠️ Como Começar

### Pré-requisitos
- Go 1.21+ (instale em [golang.org](https://golang.org))
- MongoDB (local ou [MongoDB Atlas](https://www.mongodb.com/cloud/atlas))
- Git (para clonar este repositório)

### Instalação
1. Clone o repositório:
   ```bash
   git clone https://github.com/seuusuario/sentigo.git
   cd sentigo
