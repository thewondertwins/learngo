# Come Iniziare

## Installare Go 

Mac o Linux

Scarica ed installa Go - usa sempre i packages scaricati da golang.org - mai usare homebrew o apt-get, yum, etc. Non funzionano o, peggio, possono essere stati modificati da altri.

Configura la variabile GOPATH in .bashrc, .bash_profile, .zshrc etc:

	export GOPATH=$HOME/go

Aggiungi gli eseguibili go (compilatore ed altri strumenti) alla tua variabile PATH:

	export PATH=$PATH:/usr/local/go/bin

Chiudi e riapri la sessione per abilitare il nuovo PATH, oppure esegui

	$ source .bashrc

per ricaricare la configurazione.


## Installare Go

Windows

Scarica ed installa Go - usa l'installer MSI

Configura la variabile GOPATH nelle variabili di sistema dell'utente:

	GOPATH=%userdir%/go

Aggiungi gli eseguibili go (compilatore ed altri strumenti) al tuo path:

	%userdir%/go/bin	

## Verificare l'installazione

Da un terminale/command prompt:
	
	go version

Dovresti vedere qualcosa come:

	go version 1.8 linux/amd64


## Modificare codice Go

Gli editor Go più diffusi:

vim e neovim con il plugin vim-go

emacs con go-mode.el

Visual Studio Code con vscode-go (funziona anche il debugger!) 

Atom con go-plus

IntelliJ IDEA con il Go plugin

GoLand

## Il Playground Go

Anche se non hai ancora configurato un editor sulla tua macchina, puoi comunque provare ad usare Go dal tuo browser.

[Il Playground Go](https://play.golang.org)

Il Playground Go è un servizio web che gira sui server di golang.org. Il servizio riceve in input un listato Go, lo compila, linka e lo esegue all'interno di una sandbox, quindi restituisce l'output.

## Limitazioni del Playground

Esistono delle limitazioni ai programmi che puoi eseguire nel playground:

Il playground può usare la maggior parte della libreria standard, con alcune accezioni. L'unico canale di comunicazione che il playground ha con il mondo esterno è attraverso lo standard output e lo standard error.

Nel playground l'alba dei tempi è il 2009-11-10 23:00:00 UTC (dedurre il significato di questa data è un esercizio che lasciamo al lettore). Questo rende più facile mantenere una cache dei programmi eseguiti fornendo un output deterministico.

Esistono anche delle limitazioni rispetto al tempo di esecuzione, e all'utilizzo della CPU e della memoria.

Ricapitolando: Nessun IO su file, niente di utile usando tempi o date, e non è possibile utilizzare package esterni.

A dispetto di queste limitazioni, gli sviluppatori Go amano il Go Playground - è molto pratico per condividere il codice, anche se non gira o non compila. Puoi scrivere del codice e cliccare "SHARE", il quale restituirà un URL permanente a quel listato.

Provalo adesso, con questo link:

[Hello World!](https://play.golang.org/p/992fMmkkxr) 

## Il Comando Go

Tutte le tue interazioni via linea di comando con Go avverranno attraverso il comando `go`.

Alcuni esempi di comandi frequenti:
	
	go build some/package
	go run main.go
	go test some/package
	go get github.com/someone/package
	go install some/package


## Esercizio

Dalla tua linea di comando scrivi `go` e pressa invio per vedere i diversi strumenti che il comando `go` mette a disponsizione. Provate alcuni, come:

	go env
	go list
	go version

## Scarica il materiale

Dalla tua linea di comando, usa il comando `go` per scaricare il materiale e gli esercizi di questo libro:

	go get github.com/thewondertwins/learngo

