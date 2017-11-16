# Syntax und Typen

Go hat keine überraschenden built-in Datentypen.

	uint8       die Menge aller vorzeichenlosen 8-bit (0 to 255)
	uint16      die Menge aller vorzeichenlosen 16-bit integers (0 to 65535)
	uint32      die Menge aller vorzeichenlosen 32-bit integers (0 to 4294967295)
	uint64      die Menge aller vorzeichenlosen 64-bit integers (0 to 18446744073709551615)
	int8        die Menge aller vorzeichenbehafteten 8-bit integers (-128 to 127)
	int16       die Menge aller vorzeichenbehafteten 16-bit integers (-32768 to 32767)
	int32       die Menge aller vorzeichenbehafteten 32-bit integers (-2147483648 to 2147483647)
	int64       die Menge aller vorzeichenbehafteten 64-bit integers (-9223372036854775808 to 9223372036854775807)
	float32     die Menge aller IEEE-754 32-bit Fliesskommazahlen
	float64     die Menge aller IEEE-754 64-bit Fliesskommazahlen
	complex64   die Menge aller Complexen Zahlen mit float32 reellen und imaginären Teilen
	complex128  die Menge aller Complexen Zahlen mit float64 reellen und imaginären Teilen
	byte        alias für uint8
	rune        alias für int32

# Implementierungs-Spezifische Typen

Implementierungs-Spezifische Typen.  (64 bit Breite auf 64 bit Platformen, 32 bit auf 32 bit Platformen)

	uint     entweder 32 oder 64 bits
	int      selber Breite als uint
	uintptr  ein vorzeichenloser integer Wert, gross genug um, die uninterpretierten Bits eines Zeiger Wertes zu speichern

# keine Zahlen (not number) Typen

String und Boolean (Zeichenketten und Wahrheitswerte)

	string  die Menge aller String Werte
	bool    ein Boolean (wahr oder falsch - true/false) Wert

# deklarieren und zuweisen von Werten an Variablen

ohne Initialisierung

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/withoutinit
	go run main.go


mit Initialisierung, explizite Typ

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/explicit
	go run main.go

mit Initialisierung, implizierter Typ

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/implicit
	go run main.go

Alle drei Varianten erzeugen einen integer die nicht voneinander zu unterscheiden sind. Mit implizierter Deklaration, bestimmt der Kompiler den Typ der Variablen während der Kompilierung (nicht zur Laufzeit).

# Null Werte in Go

Alle builtin Typen haben einen Null Wert. Jede zugeweisene Variable ist benutzbar, auch wenn kein Wert zugewiesen wurde.

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/zero
	go run main.go

# Konstanten

Konstanten sind Variablen die nicht verändert werden können zur Laufzeit.

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/constantstring
	go run main.go

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/constantnumber
	go run main.go

Beispiele zur versuchten Änderung von Konstanten - Fail!

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/modifyconst
	go run main.go

# Iota

Manchmal deklariert man Konstanten als Folge einer Reihenfolge:

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/sequence
	go run main.go

Das ist irgendwie hässlich. Go gibt uns zur Kompilierungszeit einen Helfer, iota, der uns hilft die Wiederholung auszulassen:

Folge mit Iota

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/iota
	go run main.go

Welche Unterschiede gibt es in den Deklarationen? Iota startet immer mit 0.

Überspringen des 1. Wertes von iota

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/iotaskip
	go run main.go

# Strukturen

Eine Struktur ist eine Sammlung von Feldern.

Strukturen sind Datentypen mit Null oder mehr Feldern.

Struktur Beispiel:

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/structs
	go run main.go

Felder einer Struktur nennt man members (Mitglieder). Angesprochen werden sie durch einen Punkt und dem Feldnamen. 


# Code Organisation

Go ist organisiert in Paketen. Ein Paket vertritt alle Dateien in einem einzelnem Verzeichnis auf der Festplatte. Ein Verzeichnis enthält nur Dateien vom selben Paket.  

Ihr habt das bereits mehrfach gesehen. Unsere Beispiele haben alle das Paket "main" benutzt, in dem es als erstes im Kopf der Datei deklariert wurde.

# Code Organisation

Quelldateien eines Paketes, müssen den Paketnamen als erste Angabe, in der Quelldatei deklarieren: 

	// Paket Deklaration
	package main

Ausführbare Programme müssen ein "main" Paket haben, das eine main() Funktion deklariert:

	func main() { ...  }

Bibliotheken Code muss einen Paketnamen deklarieren, der mit dem Verzeichnisnamen übereinstimmt, in dem die Datei liegt. Der Code im Verzeichnis "server" hat ein "package server" zu deklarieren.

# Scope (Sichtbarkeit)

Alle Variablen und Datentypen die innerhalb eines Paketes deklariert wurden, sind zu allen anderen in dem selben Paket sichtbar. 

Das bedeutet kein "public" "private" "protected" Modifizierer.

Externe Sichtbarkeit wird kontrolliert durch Grossschreibung. Datentypen und Funktionen die mit eine Grossbuchstaben beginnen sind außerhalb des aktuellen Paketes verfügbar. Datentypen und Funktionen die mit einem kleinem Buchstaben beginnen, sind nicht verfügbar außerhalb des aktuellen Paketes.

Dieses Konzept nennen wir Exportieren. Ein Zeichen das sichtbar ist außerhalb des Paketes, ist "exported" (exportiert).

# Paket Auflösung

Als wir Go installiert haben, wurde auch die GOPATH Umgebungsvariable gesetzt, in unserer entsprechenden Shell.

Ein GOPATH (GOPFAD) ist ein Arbeitsbereich für eines oder mehrere Go Projekte.

GOPATH ist die Wurzel des Arbeitsbereiches und enthält drei Verzeichnisse:

![Go Path](./gopath.png)

# Packete

Unser Quellcode und der Code von dem unsere Anwendungen abhängen, liegt in "src".

Wenn wir die Anwendung builden, wird diese in "bin" gespeichert.

Wenn wir eine Bibliothek (Library) kompilieren, wird diese in "pkg" gespeichert, in einem Unterverzeichniss die unserer Computerarchitektur entspricht. z.B.: pkg/darwin_amd64.

All das ist sehr wichtig, weil unser GOPATH ist, wie der go Kompiler die Referenzen auflöst, im Code von unseren Paketen. 

# Paket Auflösung

Wenn unser Code im Verzeichnis $GOPATH/src/blue/red liegt, dann ist der Paketname "red" und wir können Code importieren mit folgender Anweisung:

	import "blue/red"

Wir nennen "blue/red/ den Import Pfad des Paketes.

Pakete aus Quellcode Repositories wie github und bitbucket benutzen den vollen Pfad des Repositories als Teil ihres Import Pfades. Ein Projekt in meinem github Repository mit dem Namen "captainhook" hat den Import Pfad:

	"github.com/bketelsen/captainhook"

# Paket Auflösung

Deswegen, um das Paket in unserem Code nutzen zu können, muss das Paket in:

	$GOPATH/src/github.com/bketelsen/captainhook

liegen.

Wenn captainhook eine Bibliothek wäre, anstatt einer ausführbaren Datei und diese kompiliert wird, die kompilierte Version des Paketes, würde in:

	$GOPATH/pkg/darwin_amd64/github.com/bketelsen/captainhook
	(Assuming you compiled it on a Mac)

plaziert werden.

# Pakete und GOPATH

Die überwiegende Mehrheit der Entwickler wird eine GOPATH benutzen und diesen in der Umgebungsvariablen setzen und dann den vergessen.

Wie auch immer, ist es mögliche "Reine" Umgebungen zu haben, für unterschiedliche Mengen von Projekten oder gerade für ein individuelles Projekt. Dazu erstellt man einfach einen neuen GOPATH und setzt die Umgebungsvariable auf die Position des Pfades.

# Übungen


Lest die erste Hälfte des Artikels hier und führt die Übungen zu "Dein erstes Programm" und "Deine erste Bibliothek" durch.

[Getting Started with Go](https://golang.org/doc/code.html) 

Der Artikel zeigt wie man den GOPATH setzt auf $HOME/work... bitte das zu ignoieren! *BITTE NICHT DEN GOPATH ÄNDERN, WENN DER BEREITS EXPORTIERT WURDE*



