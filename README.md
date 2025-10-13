# 🐾 go-petshop

Sistema de gerenciamento para pet shops desenvolvido em Go, utilizando o framework **Chi** para roteamento HTTP e **PostgreSQL** como banco de dados. O projeto oferece funcionalidades completas para cadastro de cães, gerenciamento de avaliações e dicas personalizadas por raça.

## 📋 Sobre o Projeto

O **go-petshop** é uma API RESTful que permite a gestão eficiente de informações sobre cães em um pet shop. O sistema possibilita o cadastro completo de animais, registro de avaliações (consultas, comportamento, saúde) e fornece dicas específicas baseadas na raça do animal, auxiliando tutores e funcionários do pet shop a oferecerem o melhor cuidado possível.

## 🚀 Tecnologias

- **Go 1.21+** - Linguagem de programação
- **Chi Router** - Framework web minimalista e performático
- **PostgreSQL 15+** - Banco de dados relacional
- **pgx** - Driver PostgreSQL para Go
- **golang-migrate** - Gerenciamento de migrations
- **godotenv** - Gerenciamento de variáveis de ambiente
- **UUID** - Geração de identificadores únicos

## ✨ Funcionalidades

### 📝 Gestão de Cães
- Cadastro completo de cães (nome, raça, idade, peso, tutor)
- Listagem com filtros (por raça, tutor, idade)
- Atualização de informações
- Exclusão de registros
- Busca por ID

### ⭐ Sistema de Avaliações
- Registro de avaliações (consultas veterinárias, comportamento, grooming)
- Histórico completo de avaliações por cão
- Notas e observações detalhadas
- Data e tipo de avaliação
- Filtros por período e tipo

### 💡 Dicas por Raça
- Biblioteca de dicas específicas por raça
- Categorias: alimentação, exercícios, cuidados especiais, temperamento
- Recomendações personalizadas
- CRUD completo de dicas
- Associação automática de dicas ao cadastrar cão

### 🔐 Autenticação e Autorização
- Sistema de usuários (funcionários do pet shop)
- Autenticação via JWT
- Níveis de permissão (admin, funcionário, veterinário)
- Proteção de rotas sensíveis


## 📦 Instalação

### Pré-requisitos
- Go 1.21 ou superior
- PostgreSQL 15 ou superior
- Make (opcional, para comandos facilitados)

### Passos

1. **Clone o repositório**
```bash
git clone https://github.com/seu-usuario/go-petshop.git
cd go-petshop
```

2. **Instale as dependências**
```bash
go mod download
```

3. **Configure as variáveis de ambiente**
```bash
cp .env.example .env
# Edite o arquivo .env com suas configurações
```

Exemplo de `.env`:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=petshop_db
DB_SSLMODE=disable

SERVER_PORT=8080
JWT_SECRET=seu_secret_super_seguro
```

4. **Execute as migrations**
```bash
make migrate-up
# ou
go run cmd/migrate/main.go up
```

5. **Inicie o servidor**
```bash
make run
# ou
go run cmd/api/main.go
```

A API estará disponível em `http://localhost:8080`

## 🔌 Endpoints da API

### Cães

| Método | Endpoint | Descrição | Auth |
|--------|----------|-----------|------|
| GET | `/api/v1/dogs` | Lista todos os cães | ✓ |
| GET | `/api/v1/dogs/:id` | Busca cão por ID | ✓ |
| POST | `/api/v1/dogs` | Cadastra novo cão | ✓ |
| PUT | `/api/v1/dogs/:id` | Atualiza cão | ✓ |
| DELETE | `/api/v1/dogs/:id` | Remove cão | ✓ |
| GET | `/api/v1/dogs/breed/:breed` | Lista cães por raça | ✓ |

### Avaliações

| Método | Endpoint | Descrição | Auth |
|--------|----------|-----------|------|
| GET | `/api/v1/reviews` | Lista todas avaliações | ✓ |
| GET | `/api/v1/reviews/:id` | Busca avaliação por ID | ✓ |
| GET | `/api/v1/dogs/:dogId/reviews` | Lista avaliações de um cão | ✓ |
| POST | `/api/v1/reviews` | Cria nova avaliação | ✓ |
| PUT | `/api/v1/reviews/:id` | Atualiza avaliação | ✓ |
| DELETE | `/api/v1/reviews/:id` | Remove avaliação | ✓ |

### Dicas por Raça

| Método | Endpoint | Descrição | Auth |
|--------|----------|-----------|------|
| GET | `/api/v1/tips` | Lista todas as dicas | - |
| GET | `/api/v1/tips/:id` | Busca dica por ID | - |
| GET | `/api/v1/tips/breed/:breed` | Lista dicas por raça | - |
| POST | `/api/v1/tips` | Cria nova dica | ✓ |
| PUT | `/api/v1/tips/:id` | Atualiza dica | ✓ |
| DELETE | `/api/v1/tips/:id` | Remove dica | ✓ |

### Autenticação

| Método | Endpoint | Descrição | Auth |
|--------|----------|-----------|------|
| POST | `/api/v1/auth/register` | Registra novo usuário | - |
| POST | `/api/v1/auth/login` | Realiza login | - |


## 👨‍💻 Autor

Desenvolvido com ❤️ para aprendizado e prática de Go

---

⭐ Se este projeto foi útil para você, considere dar uma estrela no GitHub!
