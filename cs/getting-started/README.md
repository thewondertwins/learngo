# Začínáme

## Instalace Go (Mac nebo Linux)

Stáhnutí a instalace Go - vždy používejte balíčky z golang.org - nikdy ne homebrew, apt-get, yum, etc. Bývají rozbité nebo hůř - někým pozměněné.

Nastavte GOPATH v .bashrc, .bash_profile, .zshrc atd:

	export GOPATH=$HOME/go

Přidejte go binárky (kompilátory a nástroje) do $PATH:

	export PATH=$PATH:/usr/local/go/bin

Odhlašte a přihlašte se nebo spusťte

	$ source .bashrc

pro načtení změněné konfigurace.


## Instalace Go (Windows)

Stáhněte a nainstalujte Go z golang.org - Použijte MSI instalátor

Přidejte GOPATH mezi uživatelské Proměnné Prostředí

	GOPATH=%userdir%/go

Přidejte go binárky (kompilátory a nástroje) do %PATH%:

	%userdir%/go/bin

## Ověření správné instalace

V příkazovém řádku spusťte:
	
	go version

Výsledek by měl vypadat nějak takhle:

	go version 1.8 linux/amd64


## Editace Go kódu

Populární Go Editory:

vim a neovim s pluginem vim-go

emacs s go-mode.el

Visual Studio Code s vscode-go (umí i debug!)

Atom s go-plus

IntelliJ IDEA s pluginem Go


## Go Playground

I když nemáte na svém počítači editor pro Go, můžete si s Go hrát přímo v prohlížeči.

[Go Playground](https://play.golang.org)

Go Playground je webová služba, která běží na serverech golang.org. Služba umí přijmout, zkompilovat, nalinkovat a spustit program uvnitř sandboxu a vrátit jeho výstup.

## Omezení Go Playground

Na programy, které lze spustit v Go Playground se vztahují následující omezení:

Go Playground smí využívat většinu functionalit ze standardní knihovny. Jediným způsobem jakým může program komunikovat s okolním světěm je poslat výstup do stdout nebo stderr.

Čas v sandboxu vždy začíná v 2009-11-10 23:00:00 UTC (zjištění významu tohoto data budiž pro čtenáře prvním úkolem). Díky tomu lze programy snadněji kešovat, protože mají deterministický výstup.

Platí také omezení na čas běhu programu a využití paměti a procesoru.

To znamená: Žádná práce se soubory, nic použitelného co se času týče a nemožnost použít externí balíčky.

##  Go Playground

I přes všechna tato omezení se Go Playground těší oblibě mezi Go vývojáři - je to perfektní místo pro sdílení kódu i pokud jej nelze zkompilovat. Můžete zadat kód, kliknout na tlačítko "SHARE" čímž získáte permanentní odkaz k vašemu kódu.

Pojďme to zkusit s tímto kódem:

[Hello World!](https://play.golang.org/p/992fMmkkxr) 

## Go příkazy

Veškerá interakce s Go v příkazové řádce bude probíhat přes příkaz `go`.

Několik častých příkladů:
	
	go build some/package
	go run main.go
	go test some/package
	go get github.com/someone/package
	go install some/package



## Cvičení

Když v příkazové řádce spustíte příkaz `go` zobrazí se vám přehled různých nástrojů, které příkaz `go` nabízí. Zkuste například:

	go env
	go list
	go version

## Stáhněte si materiály

V příkazové řádce použijte příkaz `go` ke stáhnutí materiálů a cvičení z této knihy:

	go get github.com/thewondertwins/learngo

