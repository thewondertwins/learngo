# Como começar

## Instalando Go (Mac ou Linux)

Baixe e instale Go - sempre utilize os pacotes disponíveis em golang.org e nunca use `homebrew`, `apt-get`, `yum`, etc. Esses podem estar quebrados ou, pior, terem sido modificados por outra pessoa.

Crie uma variável GOPATH no seu .bashrc (ou .bash_profile, .zshrc, etc):

    export GOPATH=$HOME/go

Adicione os binários de Go (compiladores e ferramentas) ao seu PATH:

  export PATH=$PATH:/usr/local/go/bin

Feche o terminal para salvar as modificações ou utilize
  
    $ source .bashrc

para atualizar o ambiente.


## Instalando Go (Windows)

Baixe e instale Go (disponível em golang.org) - utilize o instalador MSI. 

Configure o GOPATH nas Variáveis de Ambiente:

    GOPATH=%userdir%/go

Adicione os binários de Go (compiladores e ferramentas) ao seu PATH:

    %userdir%/go/bin  

## Verifique a instalação

No terminal, digite:
  
    go version

Você verá algo como:

    go version 1.8 linux/amd64


## Escrevendo código Go

Editores populares para Go:

Vim e Neovim com o plugin vim-go 

Emacs com go-mode.el

Visual Studio Code com vscode-go (funciona com debugging!) 

Atom com go-plus

IntelliJ IDEA com o plugin Go


## Go Playground

Se você não tem um editor configurado na sua máquina, pode experimentar Go no browser.

[The Go Playground](https://play.golang.org)

O Go Playground é um serviço web que roda nos servidores do site golang.org. Esse serviço recebe um programa em Go, que é compilado, linkado e executado dentro de uma sandbox para então exibir o resultado. 

## Limitações do Playground

Os programas que podem ser rodados no Playground sofrem algumas limitações.

O Playground pode utilizar a maior parte da standard library, com algumas exceções. A única comunicação com o resto do mundo permitida no Playground é escrever para as saídas padrão (stdout) e de error (stderr).

No Playground, o tempo começa a ser contado a partir de 2009-11-10 23:00:00 UTC (deixamos a tarefa de descobrir o significado desta data para a leitora e o leitor). Desta maneira, é mais fácil manter um cache dos programas dando a eles outputs determinísticos.

Existem também limites de tempo de execução e de uso de CPU e memória. 

Resumindo: nenhum arquivo IO; nada útil com tempo ou datas; e não se pode utilizar pacotes externos.

Apesar dessas limitações, os desenvolvedores Go adoram o Go Playground. É um ótimo lugar para compartilhar código, mesmo que esse não possa ser executado ou compilado. Basta escrever o código e clicar no botão de compartilhar ("Share") para obter uma URL permanente para aquele conteúdo.

Teste agora com este link: 

[Hello World!](https://play.golang.org/p/992fMmkkxr) 

## Comandos Go

Toda a sua interação com Go via terminal será por meio do comando `go`.

Alguns comandos básicos:
  
    go build some/package
    go run main.go
    go test some/package
    go get github.com/someone/package
    go install some/package


## Exercício

Digite `go` no seu terminal e tecle Enter para ver as ferramentas implementadas por este comando. Teste algumas delas, como:

    go env
    go list
    go version

## Material para download

Utilize o comando `go`, no seu terminal, para baixar os materiais e exercícios deste livro:

    go get github.com/thewondertwins/learngo

