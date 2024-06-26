Criado a API

Rota de login: ` POST /login`

funciona, gera token com validade de uma hora, até o momento usuário e senhas são fake

Rota de consulta: GET

```
GET

/partnerId

/partnerName

/customerId

/customerName

```

Fazem a validação do token e devolvem os dados, a consulta está levando em média 25 segundos para retornar dados, tenho em mente criar paginação para poder otimizar e facilitar leitura.
