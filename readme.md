# Projeto de API para Teste Técnico

## Visão Geral

Este projeto consiste na criação de uma API em Golang para um teste técnico. A API permite autenticação e consulta de dados de parceiros, clientes e informações de faturamento.

## Tecnologias e Ferramentas

- **BackEnd:** Golang 1.22.4

## Como Executar

1. Certifique-se de ter o Golang 1.22.4 instalado em seu ambiente.
2. Clone o repositório do projeto.
3. Navegue até o diretório do projeto e execute o comando para baixar as dependências:

   ```bash
   go mod tidy
   ```

4. Inicie o servidor com o comando:

   ```bash
   go run main.go
   ```

5. Utilize ferramentas como Postman ou Insomnia para fazer as requisições. As rotas disponíveis estão listadas abaixo.

## Rotas

### Autenticação

- **Login:** `POST /login`
  - **Content-Type:** application/json
  - **Descrição:** Gera um token do tipo _Bearer_ com validade de uma hora.

### Rotas de Consulta

As seguintes rotas fazem a validação do token e retornam os dados solicitados:

- `GET /partnerId`
- `GET /partnerName`
- `GET /customerId`
- `GET /customerName`
- `GET /invoiceNumber`
- `GET /productId`
- `GET /skuId`
- `GET /skuName`
- `GET /subscriptionId`
- `GET /chargeStartDate`
- `GET /chargeEndDate`
- `GET /billingPreTaxTotal`
- `GET /pricingPreTaxTotal`
- `GET /usageDate`
- `GET /consumedService`
- `GET /resourceGroup`
- `GET /resourceURI`

> **Nota:** As consultas estão levando em média 25 segundos para retornar dados. Há planos para implementar paginação para otimizar e facilitar a leitura dos dados.
> A paginação foi inserida porém a resposta ainda leva tempo.
