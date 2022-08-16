# UNICO-CHALLENGE

Este repositório contém a solução para desafio empresa Unico

## Instalando e Executando o Projeto

Para instalar as dependências do projeto pode se usar o comando abaixo

```sh
make deps
```

Se preferir pode utilizar o docker-compose para executar o projeto local para desenvolvimento, antes de executar comando do docker copiar o .env.example para raiz do projeto com nome de .env

```sh
cp .env.example .env
docker-compose up --build
```

Para importar os dados do arquivos `DEINFO_AB_FEIRASLIVRES_20141.csv`, usar o comando abaixo

```sh
make run-cmd
```

Se preferir executar a aplicação sem o docker-compose, pode se utilizar o comando abaixo, lembrando que executar o comando para instalar as dependências do projeto antes, e configurar as credenciais do MySQL e executar o mesmo antes de executar o comando.

```sh
make run-api
```

Gerando relatóri de coverage do projeto pode-se utilizar o comando abaixo:

```sh
make test-cov
```

## Documentação com Swagger

Para acessar o swagger acesse o endereço da api `/swagger/index.html` após executar a API.
