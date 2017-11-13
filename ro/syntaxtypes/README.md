# Sintaxă și tipuri de date

Tipurile de date pe care le întâlnim în Go sunt cele pe care majoritatea programatorilor le cunosc deja din alte limbaje de programare.

    uint8       numere întregi pozitive cu dimensiunea de 8 biți (de la 0 la 255)
    uint16      numere întregi pozitive cu dimensiunea de 16 biți (de la 0 la 65535)
    uint32      numere întregi pozitive cu dimensiunea de 32 biți (de la 0 la 4294967295)
    uint64      numere întregi pozitive cu dimensiunea de 64 biți (de la 0 la 18446744073709551615)
    int8        numere întregi cu dimensiunea de 8 biți (de la -128 la 127)
    int16       numere întregi cu dimensiunea de 16 biți (de la -32768 la 32767)
    int32       numere întregi cu dimensiunea de 32 biți (de la -2147483648 la 2147483647)
    int64       numere întregi cu dimensiunea de 64 biți (de la -9223372036854775808 la 9223372036854775807)
    float32     numere raționale (zecimale) definite conform IEEE-754 cu dimensiunea de 32 biti
    float64     numere raționale (zecimale) definite conform IEEE-754 cu dimensiunea de 64 biti
    complex64   numere complexe având partea reală și cea imaginară compuse din numere de tip float32
    complex128  numere complexe având partea reală și cea imaginară compuse din numere de tip float64
    byte        alias pentru uint8
    rune        alias pentru int32

# Tipuri de date specifice arhitecturii utilizate

Tipuri de date specifice arhitecturii utilizate.  (64 biți pe platforme de 64 biți, 32 biți pe platforme de 32 biti)

    uint     32 sau 64 biți în funcție de arhitectura utilizată
    int      aceeași definiție ca și la uint
    uintptr  număr întreg pozitiv având dimensiunea suficient de mare încât să poată stoca valoarea unui pointer sub formă de biți neinterpretați ("raw bytes")

# Tipuri de date non-numerice

string și bool

    string  valori care conțin șiruri de caractere
    bool    valoare de tip boolean (true/false)

# Declararea variabilelor și asignarea de valori

Fără inițializare

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/withoutinit
    go run main.go


Cu inițializare, tip de date explicit

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/explicit
    go run main.go

Cu inițializare, tip de date implicit

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/implicit
    go run main.go

Toate cele trei exemple de mai sus conduc la inițializarea aceluiași număr întreg. Prin intermediul declarării implicite, compilatorul deduce tipul de date al variabilei în etapa de compilare (și nu de rulare).

# Valori de tip zero în Go

Toate tipurile de date elementare au o valoare de tip zero. Astfel, orice variabilă alocată este utilizabilă chiar dacă nu i-a fost asignată nicio valoare în decursul unui program.

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/zero
    go run main.go

# Constante

Constantele sunt variable a căror valoare nu poate fi schimbată în etapa de rulare.

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/constantstring
    go run main.go

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/constantnumber
    go run main.go

Exemplu cu modificarea valorii unei constante - Eroare!

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/modifyconst
    go run main.go

# Iota

Uneori ne dorim să declarăm o serie de constante care urmăresc o anumită secvență numerică:

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/sequence
    go run main.go

Exemplul de mai sus este destul de urât în ceea ce privește frumusețea codului scris. Go furnizează o structura ajutătoare numită iota astfel încât să putem evita repetiția redundantă:

Secvență definită cu iota

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/iota
    go run main.go

Ai observat diferența dintre cele două versiuni? Iota va porni tot timpul de la 0.

Exemplu care demonstrează cum putem sări peste prima valoare a lui iota

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/iotaskip
    go run main.go

# Structuri

O structura reprezintă o colecție de câmpuri.

Structurile sunt tipuri de date cu zero sau mai multe câmpuri definite.

Exemplu de structura:

    cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/structs
    go run main.go

Câmpurile din cadrul unei structuri sunt numite membri. Îi putem accesa prin intermediul unui punct urmat de numele acestuia.


# Organizarea codului

Codul scris în Go este organizat în pachete. Un pachet este reprezentat de toate fișierele situate într-un anumit director pe disk-ul utilizatorului. Un director poate conține doar fișiere aparținând aceluiași pachet.

Ai văzut deja acest aspect în exemplele precedente. Toate aceste exemple au utilizat pachetul "main" ca prima formulare în capul fișierului.

# Organizarea codului

Fisirele care conțin codul sursă din cadrul unui pachet trebuie să declare ca prima formulare numele pachetului din care fiecare ditre acestea face parte:

    // Declararea pachetului
    package main

Programele executabile trebuie în mod obligatoriu să declare un pachet "main" care să conțină o funcție main():

    func main() { ...  }

Codul din cadrul librăriilor trebuie să declare numele pachetului ca fiind același cu numele directorului unde acesta se află. De exemplu, codul din directorul "server" trebuie să declare în capul fiecărui fișier din acest director formularea "package server".

# Vizibilitate

Toate variabilele și tipurile declarate în cadrul unui pachet sunt vizibile oriunde în interiorul acelui pachet.

Acest aspect conduce la ideea că nu există modificatori de acces de tipul "public", "protected" sau "private".

Vizibilitatea unei structuri de date în exterioul pachetului este dată de ideea de majusculă. Astfel, tipurile sau funcțiile al căror nume începe cu litera mare sunt vizibile în afara pachetului curent. Cele al căror nume începe cu litera mică nu sunt vizibile în afara pachetului curent.

Numim acest concept Externalizare. Un simbol care este vizibil în afara pachetului în care a fost declarat se numește a fi "externalizat".

# Rezoluția la nivel de pachet

Cand ai instalat Go mai devreme ai setat variabila de mediu GOPATH în terminal.

GOPATH reprezintă un workspace pentru mai multe proiecte scrise în Go.

GOPATH reprezintă rădăcina workspace-ului și conține trei directoare:

![Go Path](./gopath.png)

# Pachete

Codul tău sursă, precum și codul de care aplicațiile tale depind se situează în directorul "src".

Cand se face build pe o aplicație, executabilul final este plasat în directorul "bin".

Cand se compilează o librărie, versiunea compilată este plasată în directorul "pkg", într-un subdirector corespunzător arhitecturii utilizate de calculator. precum pkg/darwin_amd64.

Toate aceste aspecte sunt importante deoarece variabila de mediu GOPATH determina modul în care compilatorul reușește să rezolve problema dependințelor dintre diversele pachete în codul sursă.

# Rezoluția la nivel de pachet

Spre exemplu, im ipoteza în care codul sursă se poate găsi pe disk la adresa $GOPATH/src/blue/red, numele pachetului va fi "red" și vei importa codul respectiv cu următoarea formulare:

    import "blue/red"

Numim "blue/red" calea de import a pachetului.

Pachetele care se află în surse precum github sau bitbucket trebuie importate împreună cu numele provider-ului în calea de import. Spre exemplu, un proiect aflat pe github numit "captainhook" va avea calea de import:

    "github.com/bketelsen/captainhook"

# Rezoluția la nivel de pachet

Prin urmare, pentru a putea utiliza pachetul respectiv în cod, acesta TREBUIE să fie situat la locația:

    $GOPATH/src/github.com/bketelsen/captainhook

Dacă pachetul captainhook ar fi fost o librărie în loc de executabil, după etapa de compilare, versiunea compilată a pachetului va fi plasată la următoarea locație:

    $GOPATH/pkg/darwin_amd64/github.com/bketelsen/captainhook
    (Presupunând că procesul de compilare s-a executat pe Mac)

# Pachete și GOPATH

Marea majoritate a programatorilor vor utiliza o singură variabilă de mediu GOPATH, va fi setată la început și apoi putem uita de ea.

Cu toate acestea, este posibil să ne dorim să avem așa numitele "camere curate" pentru diverse proiecte sau poate chiar un proiect separat prin crearea unei variabile de mediu GOPATH care să pointeze la locația dorită.

# Exercițiu



Citește prima jumătate a articolului următor apoi rezolvă exercițiile "Your First Program" și "Your First Library"

[Notiuni de baza despre Go](https://golang.org/doc/code.html) 

Acest articol îndrumă utilizatorul să seteze GOPATH astfel încât să pointeze către $HOME/work...ignoră acest sfat. *NU SCHIMBA GOPATH DACĂ A FOST DEJA SETATĂ ÎN PREALABIL*