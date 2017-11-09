# Zeiger und Referenz Typen

Array haben eine feste Länge und einen festen Datentyp.

- Array Beispiel

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/array
	go run main.go

# Iteration - wiederholte Anweisungen (Schleifen)

Iterieren über ein Array mit der for-Schleife

- Iteriere mit der for-Schleife

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/array
	go run main.go


Die for Anweisung ist die einzige Schleife in Go. Sie lässt sich als FOR, WHILE, DO WHILE, DO UNTIL, etc.  benutzen

# Bereich benutzen (Range)

Mit der Range Anweisung kann man über einen Bereich (Range) iterieren. Range ist die eingebaute Iterations-Funktion, die einen Index und einen Wert zurück liefert. Sie kann auf vielen verschiedenen Collection-Typen angewendet werden, inklusive Arrays:

- Iteriere mit Range

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/array
	go run main.go

Range ist sehr mächtig, es wird daher sehr häufig angwendet.

# Übung 

Ändere das Programm so ab, das mittels Range über das Array iteriert wird und alle Werte des Array multipliziert mit 2 ausgegeben werden.

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/exercises/range

Tip: führe die Multiplikation innerhalb der Funktion fmt.Println() aus.

# Auschnitte (Slices)

Ein Array hat eine feste Länge im Gegensatz dazu ein Auschnitt (Slice), hat eine dynamische Länge, eine flexible Sicjt (Ausschnitt) auf die Elemente eines Arrays. In der Praxis werden Ausschnitte häufiger genutzt, als Arrays.

Obwohl man weiß, das die Liste (Array) eine endliche und feste Anzahl von Elementen beinhaltet, benutzt man fast immer einen Ausschnitt zur Handhabung von Daten.

- Ausschnitt (Slice) Beispiele 

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/slice
	go run main.go

# Warum Ausschnitte? (Slices)

- Jede Änderung an einem Array bedeutet die Zuweisung eines neuen Arrays. Ineffizient!
- Man kann einen Ausschnitt ohne neue Zuweisung ändern
- Man kann Operationen auf Unterabschnitten von Arrays und Auschnitten einfach ausführen

Wo liegt nun den Unterschied zwischen einem Ausschnitt und Array?

Ausschnitte haben keine Angabe der Länge in der Deklaration:

	var sl []string // slice
	var ar [5]string // array

# Daten zum Ausschnitt (Slice) hinzufügen

Wir haben gezeigt, wie man mittels append, Daten zum Ausschnitt hinzufügt. Aber man kann auch den gesamten Ausschnitt auf einmal deklarien inklusive Werten:

- Ausschnitt Inlne Deklaration

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/slice
	go run main.go


# Strings (Zeichenfolgen)

Strings sind Slices von Bytes

Ein String ist einfach ein Ausschnitt von Bytes. Go hat von Haus aus Support für UTF-8 und mächtige Werkzeuge zur Arbeit mit nicht ascii Zeichen. ASCII Zeichen nehmen nur ein einzelnes Byte, eines UTF-8 Zeichens (oder Rune) der Länge von 4 Bytes ein. Go erlaubt die einfache Handhabung von Multi-Byte Runen.

- Runen Beispiel 

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/runes
	go run main.go

# Maps

Ein map ist eine _unsortierte_ Menge von Werten, indiziert von einem eindeutigen Schlüssel.

Maps müssen initialisert werden, bevor sie benutzt werden können.

Initialisierung einer map erfolgt mittels der Funktion make().

- Map Beispiel

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/maps
	go run main.go


# Maps

Map Schlüssel werden durch die Operatoren "==" (gleich) und "!=" (ungleich) definiert, deswegen können keine Funktionen, Maps oder Slices als Schlüssel der Maps benutzt werden.

- Maps können ebenfalls Inline deklariert werden:

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/mapsinline
	go run main.go

# Map Nebenläufigkeit (Concurrency)

Maps sind NICHT Thread sicher. Es ist nicht erlaubt von einer Map zu lesen und eine Map zu ändern, in 2 verschiedenen goroutinen. Wir werden später darauf zurück kommen, in den Nebenläufigkeitslektionen. Wir merken uns folgendes dazu: Wenn wir eine Map haben die an mehreren Stellen benutzt wird, dann brauchen einiges an synchronisiertem Zugriff.

Vor Go 1.7, hätte nebenläufiger Zugriff auf eine Map, Unzuverlässlichkeit des Programm zur Folge gehabt, auf Grund einer Race Condition. In Go 1.7 und später, nebenläufiger Zugriff auf eine Map bewirkt ein Stop des Programmes.

 -- TODO -- sync.Map

# Zeiger

In Go können Funktions Paramter mittels Wert (Value) oder Referenz übergeben werden. Wenn der Wert eine kurze Lebenszeit hat und nach dem Funktionsaufruf nicht mehr benötigt wird, dann übergibt man im allgemeinen mittels Wertübergabe (pass oder call-by-value).

- Übergabe mittels Wert Beispiel

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/passbyval
	go run main.go

Beachte: der übergebene Wert außerhalb der Funktion wird nicht modifiziert. es wird innerhalb der Funktion ein neuer Integer Wert erzeugt und das Ergebniss wird zurückgegeben.

Soll der Wert durch Operationen innerhalb der Funktion geändert werden, so muß er mittels Referenz als Zeiger übergeben werden:

- Übergabe mittels Referenz Beispiel

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/passbyref
	go run main.go

# Zeiger Dereferenzierung 

Zeiger Operationen sind ähnlich zu denen der C Programmiersprachen Familie, aber ohne Zeiger Arithmetik.

Benuzte & um die Speicheraddresse einer Variable zu erhalten.

Benutze * zur Dereferenzierung des Inahaltes der Variable.







