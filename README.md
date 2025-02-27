# Sentigo

Uma API RESTful construída com Go para gerenciar postagens de texto e analisar sentimentos. Desenvolvida com Gin, MongoDB e integração com LLMs via Ollama, a Sentigo permite que usuários enviem postagens, recuperem-nas e obtenham análise de sentimento e respostas geradas por IA. ("Sentigo" vem de "sentiment" + "Go", unindo sentimento e a linguagem Go!)

## 🌟 Funcionalidades

* **POST /posts**: Envie uma postagem de texto (ex.: "Adoro programar") e receba uma resposta JSON com ID, sentimento analisado por IA e uma resposta gerada automaticamente.
* **GET /posts**: Recupere todas as postagens salvas com seus sentimentos e respostas.
* **Tecnologias**: 
  * Go com Gin para lidar com HTTP de forma rápida
  * MongoDB para armazenamento flexível
  * Ollama para integração com LLM local (modelo Mistral 7B)
* **Análise de IA em Tempo Real**: Cada postagem é analisada por um modelo de linguagem para determinar o sentimento e gerar uma resposta contextual.

## 🧠 Implementação de IA

A Sentigo utiliza o Ollama como backend para executar o modelo Mistral 7B localmente, oferecendo:

* **Análise de Sentimento Avançada**: Em vez de simples correspondência de palavras-chave, a API utiliza um LLM para compreender nuances e contexto.
* **Geração de Respostas**: Para cada postagem, a API gera automaticamente uma resposta contextual e empática.
* **Processamento Local**: Todos os modelos de IA são executados localmente através do Ollama, garantindo privacidade e baixa latência.

## 🚀 Por que Sentigo?

Criei esse projeto para aprimorar minhas habilidades de backend enquanto exploro análise de sentimento e geração de linguagem natural com LLMs. É uma API limpa e funcional que demonstra design RESTful, integração com banco de dados e implementação prática de IA generativa em aplicações de produção.

## 🛠️ Como Começar

### Pré-requisitos

* Go 1.21+ (instale em golang.org)
* MongoDB (local ou MongoDB Atlas)
* Ollama (instale em ollama.ai)
* Git (para clonar este repositório)

### Instalação

1. Clone o repositório:

```
git clone https://github.com/seuusuario/sentigo.git
cd sentigo
```

2. Instale as dependências:

```
go mod tidy
```

3. Inicie o MongoDB (local: mongod, ou use sua URI do Atlas).

4. Instale e configure o Ollama:
```
# Instale o Ollama seguindo as instruções em ollama.ai
# Baixe o modelo Mistral
ollama pull mistral:7b
```

5. Inicie o servidor:

```
go run main.go
```

6. A API estará disponível em: http://localhost:8080.

### Testando a API

Envie uma postagem:
```bash
curl -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -d '{"content":"Estou muito feliz com meu novo projeto!"}'
```

Obtenha todas as postagens:
```bash
curl http://localhost:8080/posts
```

## 📁 Estrutura do Projeto

```
sentigo/
├── main.go         # Lógica principal da API
├── go.mod          # Configuração do módulo Go
└── README.md       # Você está aqui!
```

## 🔧 Planos Futuros

* Adicionar suporte ao Docker para facilitar o deploy com todos os componentes (API, MongoDB, Ollama)
* Implementar diferentes modelos de LLM para comparação de resultados
* Adicionar análise de sentimento em múltiplos idiomas
* Expandir endpoints (ex.: DELETE, UPDATE)
* Implementar cache para respostas comuns

## 💡 Construído Com

* Go - Linguagem backend 
* Gin - Framework HTTP 
* MongoDB - Banco de dados NoSQL
* Ollama - Infraestrutura para execução local de LLMs
* Mistral 7B - Modelo de linguagem para análise de sentimento e geração de respostas

## 🙌 Contribuições

Sinta-se à vontade para fazer um fork, ajustar ou sugerir ideias! Estou aprendendo e crescendo — PRs são bem-vindos.

## 📬 Contato

Me encontre no [LinkedIn](https://www.linkedin.com/in/augutso1/) ou abra uma issue aqui.

Desenvolvido por Augusto, entusiasta backend com um olho no futuro da IA.
