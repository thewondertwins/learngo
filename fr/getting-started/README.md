# Installation 

## Installer Go (Mac ou Linux)

Téléchargez et installer Go - Toujours utiliser les paquets de golang.org et ne jamais utiliser `homebrew`, `apt-get`, `yum`, etc. Ces archives ne sont pas à jour, cassées ou encore modifié par le créateur du paquet.

Ajouter une variable GOPATH dans votre .bashrc (ou .bash_profile, .zshrc, etc):

  export GOPATH=$HOME/go

Ajouter les binaires go (compilateur et outils) à votre PATH:

	export PATH=$PATH:/usr/local/go/bin

Relancer votre terminal pour effectuer les modifications ou utilisez

	$ source .bashrc

pour recharger à chaud.


## Installer Go (Windows)

Téléchargez et installez Go depuis golang.org - Utilizez l'installeur MSI

Ajouter GOPATH dans les Variables d'Environnement utilisteur:

	GOPATH=%userdir%/go

Ajouter les binaires go (compilateur et outils) à votre PATH:

	%userdir%/go/bin

## Verifer l'installation

Dans un shell:

	go version

Vous devriez voir quelque chose comme:

	go version 1.8 linux/amd64


## Ecrire du code Go

Editeur Populaire pour Go:

vim et neovim avec le plugin vim-go

emacs avec go-mode.el

Visual Studio Code avec vscode-go (le débugueur fonctionne!)

Atom avec go-plus

IntelliJ IDEA avec le plugin Go

GoLand

## Le Go Playground

Même si vous n'avez pas d'éditeur configuré en local, vous pouvez quand même jouer avec Go dans votre navigateur.

[Le Go Playground](https://play.golang.org)

Le Go Playground est un service web qui tourne sur les serveurs de golang.org. Le service reçoit un programme Go, le compile, le lance dans une sandbox et retourne le résultat.

## Limites du Playground

Il y a cependant des limites aux programmes qui peuvent être lancé dans le playground:

Le playground peut utiliser la plus grande de la librairie standard, avec quelques exceptions. Les seules communications avec le reste du monde autorisées dans le playground sont la sortie standard et la sortie d'erreur.

Dans le playground, le temps commence à 2009-11-10 23:00:00 UTC (Déterminer la signification de cette date est laissée comme exercice au lecteur). Cette propriété permet de mettre les programmes en cache grâce à leurs résultats déterministe.

## The Go Playground

Même avec ses limites, les développeurs Go adorent le Go Playground - c'est un excellent outil pour partager du code, même s'il ne peut pas se lancer ou compiler. Vous pouvez écrire du code et ensuite cliquer sur le bouton "SHARE", ce qui va créer une URL permanente vers ce code.

Essayez maintenant avec ce lien:

[Hello World!](https://play.golang.org/p/992fMmkkxr)

## Les commandes Go

Toutes les interactions avec Go en ligne de commande se font à travers de la commande `go`.

Voici quelques commandes basique:

	go build some/package
	go run main.go
	go test some/package
	go get github.com/someone/package
	go install some/package



## Exercice

Dans votre shell, tapez `go` puis Entrée pour voir la liste des différents outils proposés par la commande `go`. Essayez-en quelques-uns commme:

	go env
	go list
	go version

## Télecharger le support du livre

Dans votre shell, utilisez la commande `go` pour télécharger le support et les exercices du livre:

	go get github.com/thewondertwins/learngo

