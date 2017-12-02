# Para começar

## Instale Go (Mac ou Linux)

Baixe e instale Go - sempre utilize os pacotes de golang.org e nunca use `homebrew`, `apt-get`, `yum`, etc. Esses arquivos podem estar quebrados ou terem sido modificados por seus criadores.

Crie uma variável GOPATH no seu .bashrc (ou .bash_profile, .zshrc, etc):

  export GOPATH=$HOME/go

Adicione os binários Go (compiladores e ferramentas) ao seu PATH:

  export PATH=$PATH:/usr/local/go/bin

Feche o terminal para salvar as modificações ou utilize
  $ source .bashrc

para fazer um hot-reload.


## Instale Go (Windows)

Baixe e instale Go (disponível em golang.org) utilizando o instalador MSI. 

Crie uma GOPATH entre as variáveis de ambitente utilizadas:

  GOPATH=%userdir%/go

Adicione os binários Go (compiladores e ferramentas) ao seu PATH:

  %userdir%/go/bin  

## Verifique a instalação

No terminal, digite:
  
  go version

Você verá algo como:

  go version 1.8 linux/amd64


## Editando código em Go

Editores populares:

vim e neovim com o plugin vim-go 

emacs com go-mode.el

Visual Studio Code com vscode-go (funciona com debugging!) 

Atom com go-plus

IntelliJ IDEA com o plugin Go


##  Go Playground

Se você não tem um editor configurado na sua máquina, pode codar em Go pelo browser.

[The Go Playground](https://play.golang.org)

O Go Playground roda nos servidores de golang.org. Esse web service recebe um programa em Go, que é compilado, linkado e rodado dentro de uma sandbox para então retornar um output. 

## Limitações do Playground

Os programas que podem ser rodados no Playground têm algumas limitacões.

Pode-se utilizar a maior parte da standard library, com algumas exceções. Um programa que roda no Playground comunica-se com o exterior unicamente escrevendo outputs e erros padrões.

No Playground, o tempo começa em 2009-11-10 23:00:00 UTC (deixamos a tarefa de descobrir o significado desta data para a leitora e o leitor). Desta maneira, é mais fácil manter um cache dos programas dando a eles outputs determinísticos.

Há ainda limites de tempo de execução e de uso de CPU e memória. 

Resumindo: nenhum arquivo IO; nada útil com tempo ou datas; e não se pode utilizar pacotes externos.

Apesar das limitações, os desenvovledores Go adoram o Go Playground. É um ótimo meio para compartilhar código, mesmo que esse não possa ser rodado ou compilado. Basta escrever o código e clicar no botão "compartilhar" para obter uma URL permanente para aquele conteúdo.

Teste agora com este link: 

[Hello World!](https://play.golang.org/p/992fMmkkxr) 

## Comandos Go

Toda a sua interação com Go via terminal será por meio do comando `go`.

Os comandos mais comuns são:
  
  go build some/package
  go run main.go
  go test some/package
  go get github.com/someone/package
  go install some/package


## Exercício

Digite `go` no seu terminal e tecle Enter para ver as ferramantas implementadas por este comando. Teste algumas delas, como:

  go env
  go list
  go version

## Material para download

Utilize o comando `go`, no seu terminal, para baixar os materiais e exercícios deste livro:

  go get github.com/thewondertwins/learngo

