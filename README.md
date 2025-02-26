# Sentigo

Uma API RESTful construída com Go para gerenciar postagens de texto e analisar sentimentos. Desenvolvida com Gin, MongoDB e um toque de mágica backend, a Sentigo permite que usuários enviem postagens, recuperem-nas e obtenham uma pontuação básica de sentimento — pense nisso como uma base leve para sistemas baseados em IA. ("Sentigo" vem de "sentiment" + "Go", unindo sentimento e a linguagem Go!)

## 🌟 Funcionalidades
- **POST /posts**: Envie uma postagem de texto (ex.: "Adoro programar") e receba uma resposta JSON com ID e sentimento ("positivo", "neutro").
- **GET /posts**: Recupere todas as postagens salvas com seus sentimentos.
- **Tecnologias**: Go com Gin para lidar com HTTP de forma rápida, MongoDB para armazenamento flexível.
- **Preparada para o Futuro**: Projetada com extensibilidade em mente — conecte modelos de IA reais no futuro!

## 🚀 Por que Sentigo?
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
2. Instale as dependências:
   ```bash
   go mod tidy
3. Inicie o MongoDB (local: mongod, ou use sua URI do Atlas).

4. Inicie o servidor:
   ```bash
   go run main.go

5. A API estará disponível em: http://localhost:8080.

## 📁 Estrutura do Projeto
```bash
sentigo/
├── main.go         # Lógica principal da API
├── go.mod          # Configuração do módulo Go
└── README.md       # Você está aqui!
```
## 🔧 Planos Futuros
Adicionar suporte ao Docker para facilitar o deploy.
Integrar um LLM real (ex.: via Hugging Face) para análise avançada de sentimento.
Expandir endpoints (ex.: DELETE, UPDATE).

## 💡 Construído Com
Go - Linguagem backend
Gin - Framework HTTP
MongoDB - Banco de dados NoSQL

## 🙌 Contribuições
Sinta-se à vontade para fazer um fork, ajustar ou sugerir ideias! Estou aprendendo e crescendo — PRs são bem-vindos.

## 📬 Contato
Me encontre no [LinkedIn](https://www.linkedin.com/in/augutso1/) ou abra uma issue aqui.

Desenvolvido por Augusto, entusiasta backend com um olho no futuro da IA.
