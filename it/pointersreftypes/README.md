# Puntatori e tipi di riferimento

Gli array hanno una capacità prefissata, e un tipo predefinito.

## Esempio di Array

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/array
	go run main.go

# Iterare

Iterare su un array usando il ciclo for

## Iterare con il for

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/array
	go run main.go

For è l'unico costrutto che rappresenta un ciclo in Go. Può essere usato per implementare FOR, WHILE, DO WHILE, DO UNTIL, etc.

# Utilizzo del RANGE

Puoi anche iterare su una collezione in Go utilizzando l'istruzione range. Range è una funzione iterativa inclusa nel linguaggio, che restituisce l'indice e il valore di diversi tipi di collezione, inclusi gli array:

## Iterate with Range

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/array
	go run main.go

Range è molto utile, e lo vedrete spesso.

# Esercizio

Modifica questo esempio per iterare sull'array utilizzando l'istruzione range, e stampa il valore dell'array moltiplicato per 2.

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/exercises/range

suggerimento: esegui la moltiplicazione all'interno della funzione fmt.Println().

# Slice

Un array ha una dimensione prefissata. Una slice, al contrario, è una vista con capacità dinamica sugli elementi di un array. In effetti, le slice sono più comuni degli array.

A meno che non sappiate a priori che la vostra lista conterrà un numero finito e fisso di elementi, probabilmente andrete ad utilizzare quasi sempre le slice per gestire i dati.

## Esempio di Slice

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/slice
	go run main.go

# Perché slice?

- Ogni modifica ad un array va ad allocare un nuovo array. Inefficiente!
- È possibile modificare una slice senza allocazione
- È possibile e semplice operare su sottosezioni di Array e Slice

Come si distungue una slice da un array?

Le slice non hanno la definizione della capacità nella dichiarazione:

	var sl []string // slice
	var ar [5]string // array

# Aggiungere dati ad una slice

Abbiamo mostrato come usare la append per aggiungere valori ad una slice, ma è anche possibile dichiarare una intera slice ed i suoi valori contemporaneamente:

## Dichiarazione di una slice inline

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/slice
	go run main.go

# Stringhe

Le Strighe sono Slice di Byte

Una stringa è semplicemente una slice di byte. Go include nativamente il supporto per UTF-8, e gli strumenti per lavorare con caratteri non ASCII. Al contrario dei caratteri ASCII che occupano solo un singolo byte, i caratteri UTF-8 (detti anche Rune) possono occupare fino a 4 byte. Go vi consente di gestire facilmente rune composte da più byte.

## Esempio di Rune

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/runes
	go run main.go

# Mappe

Una mappa è un set di valori _non ordinato_ indicizzati da una chiave univoca.

Le mappe devono essere inizializzate prima di poter essere usate.

Una mappa si inizializza per mezzo della funzione make().

## Esempio di Mappa

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/maps
	go run main.go

Le chiavi di una mappa devono necessariamente avere la definizione degli operatori di uguaglianza e disuguaglianza "==" e "!=". Per questo motivo non è possibile usare funzioni, mappe o slice come chiavi di una mappa.

## Anche le mappe possono essere dichiarate inline:

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/mapsinline
	go run main.go

# Mappe Concorrenza

Le mappe NON sono threadsafe. È profondamente sbagliato leggere da una mappa e modificare la mappa in due Go routine separate. Vedremo più avanti questi concetti nelle sezioni sulla concorrenza: per ora prendetene atto. Quando una mappa viene usata in più posti, sarà necessario sincronizzarne l'accesso.

Prima di Go 1.7, l'accesso concorrente ad una mappa avrebbe avuto come esito un programma non affidabile a causa delle race condition (contesa di risorse). A partire da Go 1.7, l'accesso concorrente ad una mappa bloccherà il programma.

 -- TODO -- sync.Map

# Puntatori

Go consente il passaggio di parametri a funzione per valore o per riferimento. Generalmente passerete parametri per valore quando il tipo ha una vita breve e non dovrà essere usato dopo la chiamata a funzione.

## Esempio di passaggio parametro per valore

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/passbyval
	go run main.go

Notare che non abbiamo modificato il valore che è stato passato in ingresso, abbiamo invece creato un nuovo intero e restituito questo come risultato.

Se fosse necessario operare su un valore e far si che venga modificato durante l'operazione, questo deve essere passato per riferimento usando un puntatore:

## Esempio di passaggio parametro per riferimento

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/passbyref
	go run main.go

# Dereferenziare un puntatore

Le operazioni sui puntatori sono simili a quelle presenti nella famiglia di linguaggi come C, ma in Go non è possibile usare l'aritmetica dei puntatori.

Utilizza il simbolo `&` per estrarre l'indirizzo di una variabile.

Utilizza il simbolo `*` per dereferenziare una variabile







