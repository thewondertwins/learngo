# Noțiuni de baza

## Instalare Go 

Mac sau Linux

Descarcă și instalează Go - întotdeauna utilizează pachetele oficiale pe care le poți găsi la adresa golang.org - niciodată nu utiliza homebrew sau apt-get, yum, etc. Sunt defecte sau mai rău de atât -- modificate de către cineva pe serverul de download.

Seteaza variabila de mediu GOPATH în .bashrc, .bash_profile, .zshrc etc:

    export GOPATH=$HOME/go

Adaugă executabilele de Go (compilatoare sau instrumente) în variabila de mediu PATH:

    export PATH=$PATH:/usr/local/go/bin

Restartează sesiunea de terminal pentru a putea observa schimbările efectuate sau execută următoarea comanda

    $ source .bashrc

pentru hot-reload.


## Instalare Go 

Windows

Descarcă și instalează Go - Utilizează installer-ul MSI

Setează variabila de mediu GOPATH

    GOPATH=%userdir%/go

Adaugă executabilele de Go (compilatoare sau instrumente) în variabila de mediu PATH:

    %userdir%/go/bin    

## Verifică Instalarea

Din linia de comandă a unui terminal execută:
    
    go version

Si ar trebui să vezi ceva asemănător:

    go version 1.8 linux/amd64


## Editarea codului Go

Editoare de Go populare:

vim sau neovim cu plugin-ul vim-go 

emacs cu go-mode.el

Visual Studio Code cu vscode-go (funcționează și pentru debugging!) 

Atom cu go-plus

IntelliJ IDEA cu plugin-ul de Go


## Playground-ul de Go

Chiar dacă nu ai un editor configurat local pe calculatorul tău, poți executa și analiza cod Go din browser.

[Playground-ul de Go](https://play.golang.org)

Playground-ul de Go este un serviciu web care rulează pe serverele de la adresa golang.org. Serviciul acceptă ca input un program scris în Go pe care îl compilează și îl rulează într-un mediu izolat tip sandbox după care returnează un output.

## Limitarile Playground-ului

Exista o serie de limitări priving programele care pot rula în playground:

Playground-ul poate utiliza o mare parte din librăria standard, cu anumite excepții. Singurele mijloace de comunicare ale unui program scris în playground cu autorul său sunt standard output și standard error.

In cadrul playground-ului, timpul de start este considerat a fi 2009-11-10 23:00:00 UTC (lăsăm ca un fel de exercițiu pentru cititor să descopere însemnătatea acestei date). Prin intermediul strategiei de caching, acest aspect conduce la comportamentul determinist al programelor scrise în playground.

Există, de asemenea, limite privind gradul de utilizare a resurselor (CPU și memorie).

In concluzie: Nu sunt permise operațiuni de tip IO cu fișiere, operațiunile cu date și timp sunt destul de limitate și de asemenea nu se pot utiliza pachete externe.

##  Playground-ul de Go

Desi există aceste nenumărate limitări, programatorii de Go adoră acest serviciu - reprezintă un loc ideal de a imparatsi cod cu alți programatori, chiar dacă acel cod nu poate fi compilat sau rulat. Poți insera cod, apoi, prin apăsarea butonului "SHARE", va fi generat un link permanent către codul respectiv. 

Incearcă-l acum cu acest link: 

[Hello World!](https://play.golang.org/p/992fMmkkxr) 

## Comanda de terminal Go

Toate interacțiunile cu limbajul Go din cadrul terminalului se vor face prin intermediul comenzii `go`

Mai jos avem câteva comenzi foarte des întâlnite:
    
    go build some/package
    go run main.go
    go test some/package
    go get github.com/someone/package
    go install some/package



## Exercițiu

Din linia de comandă a unui terminal scrie `go` și apasă tasta "Enter" pentru a putea observa diversele instrumente asociate comenzii `go`. Încearcă ceva de genul:

    go env
    go list
    go version

## Descărcarea Materialului

Din linia de comandă a terminalului utilizează comandă `go` pentru a putea descărca materialele didactice și exercițiile din această carte:

    go get github.com/thewondertwins/learngo

