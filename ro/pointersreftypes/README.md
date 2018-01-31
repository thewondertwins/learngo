# Pointeri și tipuri referință

Array-urile au o dimensiune fixă și un tip de date de asemenea fix.

## Exemplu de array

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/array
    go run main.go

# Iteratii

Iterarea asupra unui array utilizând structura repetitivă for

## Iterarea cu for

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/array
    go run main.go

For reprezintă unică structura repetitivă în Go. Utilizează-o când ai nevoie de FOR, WHILE, DO WHILE, DO UNTIL, etc.

# Utilizarea comenzii RANGE

Poți, de asemenea, itera asupra oricărei colecții în Go prin utilizarea comenzii range. Range reprezintă o funcție internă a limbajului care returnează indexul (poziția) și valoarea fiecărui element dintr-o colecție și care poate fi aplicată mai multor tipuri de colecții, printre care și array-uri:

## Iterarea cu Range

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/array
    go run main.go

Range este un concept foarte puternic în Go și îl vei utiliza foarte mult.

# Exercițiu 

Modifică acest exemplu astfel încât să iterezi asupra unui array prin intermediul comenzii range și printează valoarea array-ului multiplicată cu 2.

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/exercises/range

sfat: execută multiplicarea în interiorul funcției fmt.Println().

# Slice-uri


Un array are dimensiune fixă. Un slice, pe de altă parte, are dimensiune variabilă, este flexibil și reprezintă o fereastră către elementele unui array. În practică, slice-urile sunt mult mai des întâlnite decât array-urile.

Aproape întotdeauna vei dori să utilizezi un slice cu excepția cazului în care cunoști dinainte ce dimensiune va avea colecția ta, situație în care vei folosi un array.

## Exemple cu slice-uri 

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/slice
    go run main.go

# De ce să utilizăm slice-urile?

- Orice modificare operată asura unui array alocă un alt array. Ineficient!
- Poți modifică un slice fără a genera o nouă alocare de memorie
- Putem opera pe subsectiuni ale unui array foarte ușor prin intermediul slice-urilor

Cum putem observă diferența dintre un slice și un array?

Slice-urile nu au dimensiunea specificată în declarare:

    var sl []string // slice
    var ar [5]string // array

# Adaugarea de date într-un slice

Am arătat în exemplele precedente utilizarea funcției append în adăugarea de noi valori unui slice existent dar este posibil să declarăm slice-ul cu toate valorile dintr-o bucată:

## Declararea dintr-o bucată a unui slice

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/slice
    go run main.go

# Siruri de caractere (string-uri)

Sirurile de caractere nu sunt altceva decât slice-uri de octeți

Un șir de caractere, sau un string, nu este altceva decât un slice de octeți. Go are suport integrat pentru UTF-8 și instrumente puternice de lucru cu caractere non-ASCII. Deși caracterele ASCII ocupă un singur octet, caractere de tip UTF-8 (sau rune-uri) pot ocupă până la 4 octeți. Go este capabil să manipuleze cu lejeritate rune-uri de tip mulți-octet.

## Exemplu de rune 

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/runes
    go run main.go

# Dictionare

Un dicționar reprezintă o colecție _dezordonată_ de valori indexate cu o cheie unică.

Dictionarele trebuie să fie initializate înainte de a fi utilizate.

Initializarea unui dicționar se face cu ajutorul funcției make().

## Exemplu de dicționar 

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/maps
    go run main.go

# Dictionare

Cheile dintr-un dicționar trebuie să suporte operatorii "==" și "!=". Prin urmare, nu pot fi utilizate funcții, dicționare sau slice-uri ca și tipuri de date pentru chei.

## Si dicționarele pot fi declarate dintr-o bucată:

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/mapsinline
    go run main.go

# Concurenta în cazul dictionarelor

Dictionarele NU sunt threadsafe. Nu este în regulă să citim dintr-un dicționar și să îl modificăm în același timp în două procese concurente diferite. Vom discuta aceste aspecte mai târziu în cadrul lecțiilor despre concurență. Acum nu doresc decât să va faceți un model mental. Când avem un dicționar care este utilizat în mai mult de două locuri în același timp trebuie să folosim un mecanism de sincronizare a accesului.

Inainte de Go 1.7, accesul concurent asupra unui dicționar cauza o serie de comportamente impredictibile ale programului. În Go 1.7 și versiuni ulterioare, accesul concurent asupra unui dicționar va opri programul imediat.

 -- VA URMĂ -- sync.Map

# Pointeri

Go ne permite să pasăm parametri unei funcții, fie prin valoare sau prin referință. În general, parametrii vor fi pasați prin valoare când variabilă nu va fi utilizată după închiderea funcției.

## Exemplu cu pasare prin valoare

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/passbyval
    go run main.go

Dupa cum se poate observă, nu am modificat valoarea pe care am păsat-o, doar am creat o nouă valoare pe care am returnat-o că rezultat.

Daca dorești să operezi asupra unei valori și dorești că această valoarea să fie modificată în cadrul funcției, atunci trebuie pasată prin referință prin intermediul unui pointer:

## Exemplu cu pasare prin referință

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/passbyref
    go run main.go

# Dereferentierea unui pointer

Operatiile cu pointeri sunt similare cu cele întâlnite în cadrul limbajelor din familia C, cu excepția faptului că în Go nu putem efectua operații aritmetice cu pointeri.

Utilizeaza & pentru a putea obține adresa de memorie a unei variabile.

Utilizeaza * pentru a dereferentia o variabilă.