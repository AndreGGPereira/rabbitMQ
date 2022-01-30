# api Test
==================

# Migratrion

Para fazer um UP ou Down em uma migração basta executar o comando abaixo
com a utiliação do mage https://magefile.org/.
O comando buscara os arquivos sql na pasta "migrations"
Caso for necessario alterar a url de acesso ao banco no arquivo magefile.go

exemplo para up:
## mage Migration up
exemplo para down:
## mage Migration down

# .ENV

Criar variavel de ambiente para realizar a conexão com RabbitMQ 
exemplo: 
#  AMQP_SERVER_URL= amqp://guest:guest@message-broker:5672/