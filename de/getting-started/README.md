# Start

## Go installieren (Mac oder Linux)

Go herunterladen und installieren -  bitte benutzt immer die Installer von golang.org - niemals Pakete mittels homebrew oder apt-get, yum etc. installieren. Diese können beschädigt sein oder durch den upstream modifiziert.

GOPATH setzen - in der .bashrc, .bash_profile, .zshrc etc:

	export GOPATH=$HOME/go

Hinzufügen der go Binärdateien (Compiler und Werkzeuge) zum Benutzer Pfad:

	export PATH=$PATH:/usr/local/go/bin

Um die geänderten Einstellungen zu aktivieren einmal abmelden und wieder anmelden oder von der Kommandozeile

	$ source .bashrc

ausführen um die Änderungen sofort zu aktivieren.

## Go installieren (Windows)

Go herunterladen und installieren von golang.org - den MSI Installer benutzen

GOPATH setzen in den Benutzer Umgebungsvariablen

	GOPATH=%userdir%/go

Hinzufügen der go Binärdateien (Compiler und Werkzeuge) zum Benutzer Pfad:

	%userdir%/go/bin

## Überprüfen der Installation

Von der Kommandozeile:

	go version

Folgende ähnliche Ausgabe sollte erscheinen:

	go version 1.8 linux/amd64

## Go Code schreiben

beliebte Go Editoren:

vim und neovim mit dem vim-go plugin

emacs mit go-mode.el

Visual Studio Code mit vscode-go (unterstützt debugging!)

Atom mit go-plus

IntelliJ IDEA mit Go plugin

GoLand

## Der Go Spielplatz

Gerade wenn kein Editor lokal konfiguriert wurde, kann man mit Go testen vom Browser aus.

[The Go Playground](https://play.golang.org)

Der Go Spielplatz ist ein Web Dienst, der auf den golang.org Webservern zur Verfügung steht. Der Dienst empfängt ein Go Programm, kompiliert und linkt es und führt das Programm innerhalb einer Sandbox aus. Die Ausgabe des Programms wird dann zurückgegeben zum Web Dienst.

## Einschränkungen für den Go Spielplatz

Es gibt einige Einschränkungen für den Spielplatz:

Im Spielplatz kann fast alles, mit einigen Ausnahmen, aus der Standard Blibliothek benutzt werden. Die einzige Kommunikationsmöglichkeit für ein Go Programm (im Spielplatz) zur Außenwelt ist, über den Standard Output und Standard Error.

Im Spielplatz beginnt die Zeit mit 2009-11-10 23:00:00 UTC (die Bedeutung dieses Datums vorher zubestimmen, ist eine Übung für den Leser). Dies erleichtert das Zwischenspeichern von Programmen, indem ihnen eine vorher bestimmbare Ausgabe gegeben wird.

Es gibt Einschränkungen für die Ausführungszeit der CPU und Speichernutzung.

Deswegen: Keine Datei Ein- und Ausgabe, nichts brauchbares für Zeit und Datum, keine Benutzung externer Pakete.

## Der Go Spielplatz

Gerade mit all diesen Einschränkungen, Go Entwickler lieben den Spielplatz - es ist ein großartiger Bereich um Code zu teilen, gerade auch wenn der Code nicht kompiliert oder läuft. Man kann Code einfügen und dann den "Teilen" (Share) Knopf drücken und erhält eine permanente URL zu diesem Code.

Versuche es einmal mit disem Link:

[Hello World!](https://play.golang.org/p/992fMmkkxr)

## Das Kommando Go

Alle Kommandozeilen Interaktionen mit Go, werden mit dem `go` Kommando ausgeführt.

Verschiedene gebräuchliche Kommandos:

	go build some/package
	go run main.go
	go test some/package
	go get github.com/someone/package
	go install some/package


## Übung

Von der Kommandozeile aus gib `go` ein und drücke Enter um zu sehen welche verschiedenen Werkzeuge das `go`Kommando beinhaltet. Versuche etwas wie dieses:

	go env
	go list
	go version

## Lade zusätzliches Material herunter

von der Kommandozeile, benutze das Kommando `go` um zusätzliches Material und Übungen zu diesem Buch herunterzuladen:

go get github.com/thewondertwins/learngo	

