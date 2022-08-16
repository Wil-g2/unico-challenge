# UNICO CHALLENGE

Este repositório contém a solução para desafio da empresa Unico

## Instalando e Executando o Projeto

Para instalar as dependências do projeto pode se usar o comando abaixo:

```sh
make deps
```

Se preferir pode utilizar o docker-compose para executar o projeto local para desenvolvimento, antes de executar comando do docker copiar o .env.example para raiz do projeto com nome de .env:

```sh
cp .env.example .env
docker-compose up --build
```

Para importar os dados do arquivos `DEINFO_AB_FEIRASLIVRES_20141.csv`, usar o comando abaixo:

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

![image](https://user-images.githubusercontent.com/26700193/184763042-eb910227-141c-4b2e-aac4-245bd4a5e350.png)

![image](https://user-images.githubusercontent.com/26700193/184763274-7af2e322-2a4d-4ca6-8f60-5b4d9dab7ecd.png)

![image](https://user-images.githubusercontent.com/26700193/184765204-fc9a5e33-9f02-46e9-a1ac-27c32f091022.png)
![image](https://user-images.githubusercontent.com/26700193/184765577-fd3955fd-98fe-408d-bfa8-234e85a3a22d.png)


## Testes de Unidade

Foi feito teste de unidade para camada de Domain, e coverage atual está na imagem abaixo:
![image](https://user-images.githubusercontent.com/26700193/184764069-11d07578-c2bc-4b0c-9c7b-7a286e393271.png)

