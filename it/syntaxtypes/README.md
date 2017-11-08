# Sintassi e Tipi

Go include dei tipi piuttosto scontati;

	uint8       l'insieme di tutti gli interi di 8-bit senza segno (da 0 a 255)
	uint16      l'insieme di tutti gli interi di 16-bit senza segno (da 0 a 65535)
	uint32      l'insieme di tutti gli interi di 32-bit senza segno (da 0 a 4294967295)
	uint64      l'insieme di tutti gli interi di 64-bit senza segno (da 0 a 18446744073709551615)
	int8        l'insieme di tutti gli interi di 8-bit con segno (da -128 a 127)
	int16       l'insieme di tutti gli interi di 16-bit con segno (da -32768 a 32767)
	int32       l'insieme di tutti gli interi di 32-bit con segno (da -2147483648 a 2147483647)
	int64       l'insieme di tutti gli interi di 64-bit con segno (da -9223372036854775808 a 9223372036854775807)
	float32     l'insieme di tutti i numeri IEEE-754 a 32-bit a virgola mobile	float64     l'insieme di tutti i numeri IEEE-754 a 64-bit a virgola mobile
	complex64   l'insieme di tutti i numeri complessi con parti reale ed immaginaria float32
	complex128  l'insieme di tutti i numeri complessi con parti reale ed immaginaria float64
	byte        alias di uint8
	rune        alias di int32

# Tipi dipendenti dall'implementazione

Tipi dipendenti dall'implementazione (64 bit su piattaforma a 64 bit, 32 bit su piattaforma a 32 bit)

	uint     32 o 64 bits
	int      stessa dimensione di uint
	uintptr  intero senza segno sufficientemente grande da memorizzare i bit non interpretati di un puntatore

# Tipi non numerici

string e bool

	string  il set dei valori stringa
	bool    valore booleano (vero/falso)

# Dichiarare ed assegnare valori a variabili

Senza inizializzazione

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/withoutinit
	go run main.go


Con inizializzazione, tipo esplicito

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/explicit
	go run main.go

Con inizializzazione, tipo implicito

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/implicit
	go run main.go

Tutti e tre producono un intero che è indistinguibile dagli altri. Con la dichiarazione implicita, il compilatore determina il tipo della variabile in fase di compilazione (non in fase di esecuzione).

# Valori Zero in Go

Ogni tipo fornito dal linguaggio possiede un valore zero (di inizializzazione). Qualsiasi variabile allocata è quindi utilizzabile, anche quando non le sia mai stato assegnato un valore.

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/zero
	go run main.go

# Costanti

Le costanti sono variabili che non possono essere modificate durante l'esecuzione.

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/constantstring
	go run main.go

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/constantnumber
	go run main.go

Esempio di modifica di una costante - Fallisce!

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/modifyconst
	go run main.go

# Iota

A volte potreste voler dichiarare delle costanti che seguano una determinata sequenza:

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/sequence
	go run main.go

Questo è piuttosto sgradevole. Go ci fornisce uno strumento, in fase di compilazione, chiamato iota, che ci consente di evitare la ripetizione manuale:

Sequenza con Iota

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/iota
	go run main.go

Notate la differenza tra le due versioni? Iota parte sempre da 0.

Saltare il primo valore di Iota

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/iotaskip
	go run main.go

# Strutture

Una struttura (`struct`) è una collezione di campi.

Le strutture sono tipi con zero o più campi.

Esempio di struttura:

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/structs
	go run main.go

I campi di una struttura sono chiamati membri. Potete fare riferimento ad essi utilizzando il punto e il nome del campo.

# Organizzazione del codice

Il codice Go è organizzato in `packages`. Un `package` rappresenta tutti i file all'interno di una singola directory su disco. One directory può contenere solo file appartenenti allo stesso package.

Avrete già notato più volte che nei nostri esempi fino ad ora abbiamo sempre usato il package "main", dichiarato in cima al file.

I file sorgente all'interno di un package devono dichiarare il nome del package in cima al file come prima istruzione:

	// Dichiarazione del package
	package main

I programmi eseguibili devono avere un package "main" che dichiari a sua volta una funzione main():

	func main() { ...  }

Il codice delle librerie deve dichiarare un nome di package che corrisponda al nome della cartella nella quale si trova. Se il codice si trova nella cartella "server", deve quindi dichiarare il "package server".

# Visibilità

Tutte le variabili e i tipi dichiarati all'interno di un package sono visibili ovunque all'interno del medesimo package.

Questo significa che non esistono modificatori "public" "private" "protected".

La visibilità all'esterno del package viene controllata tramite capitalizzazione (maiuscolo). I tipi e le funzioni che iniziano con una lettera maiuscola, sono disponibili anche all'esterno del package in cui risiedono. I tipi e le funzioni che iniziano, al contrario, con la lettera minuscola, sono "non esportati", pertanto visibili soltanto all'interno del loro stesso package.

In Go chiamiamo questo concetto Esportazione. Un simbolo che è visibile al di fuori del suo package si dice "esportato".

# Risoluzione dei Package

Quando avete installato Go in precedenza, avete configurato una variabile di ambiente `GOPATH` all'interno della shell.

GOPATH è una specie di workspace per uno o più progetti Go.

GOPATH è la radice del workspace e contiene tipicamente le tre directory:

![Go Path](./gopath.png)

# Packages

Il vostro codice sorgente, e il codice da cui le vostre applicazioni dipendono, si trova all'interno di "src".

Quando compilate una applicazione, il binario viene salvato in "bin".

Quando compilate una libreria, questa viene salvata in "pkg", in una sottodirectory dipendente dal nome dell'architettura del costro computer, come ad esempio pkg/darwin_amd64 (se Mac).

Tutto questo è importante in quanto il vostro GOPATH è quanto determina il modo in cui il compilatore Go risolve i riferimenti ai package all'interno del codice.

# Percorso di import

Se il vostro codice risiede all'interno di `$GOPATH/src/blue/red`, il suo nome di package è "red", ed è possibile importare tale codice mediante la seguente istruzione nel codice:

	import "blue/red"

Chiamiamo "blue/red" il percorso di import del package.

I package che risiedono nei repositories di codice sorgente come Github e Bitbucket avranno un percorso di import completamente definito, includendo il dominio. Ad esempio un progetto nel mio repository Github chiamato "captainhook" ha un percorso di import:

	"github.com/bketelsen/captainhook"

Di conseguenza, per utilizzare tale package all'interno del vostro codice, questo package deve necessariamente risiedere in:

	$GOPATH/src/github.com/bketelsen/captainhook

Se captainhook fosse una libreria al posto di un eseguibile, in fase di compilazione, la versione compilata verrebbe generata in:

	$GOPATH/pkg/darwin_amd64/github.com/bketelsen/captainhook
	(Assimendo che abbiate compilato su un Mac)

# Package e GOPATH

La grande maggioranza degli sviluppatori utilizza un unico GOPATH, configurato come variabile d'ambiente, senza doverlo più cambiare o considerare.

Tuttavia, è possible andare a creare degli spazi vuoti per diversi insiemi di progetti, o addirittura per un singolo progetto, semplicemente creando un nuovo GOPATH e configurando la variabila d'ambiente che punti a questa nuova posizione.

# Esercizio

Leggi la prima metà dell'articolo qui sotto, e poi fai l'esercizio "Your First Program" e "Your First Library"

[Getting Started with Go](https://golang.org/doc/code.html) 

Questo articolo ti dice di configurare il GOPATH come `$HOME/work`... ignoralo. *NON CAMBIARE IL GOPATH CHE HAI GIÀ ESPORTATO*