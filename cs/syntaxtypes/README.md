# Syntaxe a Typy

Go používá typy, jejichž výčet vás jistě nijak nepřekvapí.

	uint8       rozsah všech kladných celých čísel do 8 bitů (0 až 255)
	uint16      rozsah všech kladných celých čísel do 16 bitů (0 až 65535)
	uint32      rozsah všech kladných celých čísel do 32 bitů (0 až 4294967295)
	uint64      rozsah všech kladných celých čísel do 64 bitů (0 až 18446744073709551615)
	int8        rozsah všech kladných i záporných celých čísel do 8 bitů (-128 až 127)
	int16       rozsah všech kladných i záporných celých čísel do 16 bitů (-32768 až 32767)
	int32       rozsah všech kladných i záporných celých čísel do 32 bitů (-2147483648 až 2147483647)
	int64       rozsah všech kladných i záporných celých čísel do 64 bitů (-9223372036854775808 až 9223372036854775807)
	float32     rozsah všech desetinných čísel (podle IEEE-754) do 32 bitů
	float64     rozsah všech desetinných čísel (podle IEEE-754) do 64 bitů
	complex64   rozsah všech komplexních čísel s reálnou složkou float32
	complex128  rozsah všech komplexních čísel s reálnou složkou float64
	byte        alias pro uint8
	rune        alias pro int32

# Typy dle implementace

Typy dle implementace.  (64 bitů na 64 bitových platformách, 32 bitů on 32 bitových platformách)

	uint     32 nebo 64 bitů
	int      stejně velká jako uint
	uintptr  kladné celé číslo dostatečně velké pro uložených neinterpretovaných bitů ukazatele

# Nečíselné typy

string a bool

	string  řetězcové hodnoty
	bool    Booleanovská hodnota (true/false)

# Deklarace a přiřazení hodnoty proměnným

Neinicializovaná proměnná

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/withoutinit
	go run main.go


Inicializovaná proměnná s explicitním typem

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/explicit
	go run main.go

Inicializovaná proměnná s implicitním typem

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/implicit
	go run main.go

Výsledkem všech tří je kladné číslo, které se nijak neliší od ostatních. Pokud je použita implicitní deklarace, typ proměnné se rozhoduje již při kompilaci (nikoliv při spuštění).

# Nulové hodnoty v Go

Všechny výchozí typy mají nulové hodnoty. Alokovaná proměnná je použitelná i pokud neobsahuje žádnou hodnotu.

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/zero
	go run main.go

# Konstanty

Konstanty jsou proměnné, které nelze za běhu modifikovat.

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/constantstring
	go run main.go

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/constantnumber
	go run main.go

Příklad modifikace konstanty - skončí chybou!

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/modifyconst
	go run main.go

# Ióta

Někdy chcete deklarovat proměnné, které jsou v pořadí:

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/sequence
	go run main.go

Tato definice není zrovna hezká. Go nabízí kompilační helper zvaný Ióta, díky kterému se můžete vyhnout opakování:

Sekvence s Iótou

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/iota
	go run main.go

Všimněte si rozdílů mezi touto a předchozí verzí. Ióta  vždy začíná od nuly.

Přeskočení první hodnoty Ióty

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/iotaskip
	go run main.go

# Záznam (Struct)

Záznam je kolekce prvků

Záznamy jsou typy s nula nebo více prvky.

Příklad Záznamu:

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/syntaxtypes/demos/structs
	go run main.go

Prvkům v záznamu říkáme říkáme members (členové). Odkazujeme se na ně přes tečku a název prvku.


# Organizace kódu

Go kód je organizován do balíčků. Balíček reprezentuje všechny soubory v jedné složce na disku. Složka smí obsahovat pouze soubory ze stejného balíčku.

Už jste si možná všimli, že všechny naše příklady používají balíček "main" deklarovaný na začátku souboru.

# Organizace kódu

Zdrojové soubory v balíčku musí definovat název balíčku vždy na začátku souboru jako první instrukci:

	// Deklarace balíčku
	package main

Spustitelné programy musí obsahovat balíček "main", který obsahuje funkci main():

	func main() { ...  }

Kód knihovny musí deklarovat název balíčku, který odpovídá názvu složky ve které se nachází. Kód ve složce "server" musí deklarovat "package server".

# Scope (Rámec působnosti)

Všechny proměnné a typy deklarované uvnitř balíčku jsou viditelné všem uvnitř balíčku.

Čili žádné modifikátory jako "public" "private" nebo "protected".

Externí viditelnost se řeší kapitálkami. Typy a funkce začínající velkým písmenem jsou dostupné i mimo svůj balíček. Typy a funkce začínající malým písmenem se neexportují a nejsou k dispozici mimo svůj balíček.

Tomuto konceptu říkáme exportování. Symbol, který je vidět mimo svůj balíček je "exported" (exportovaný).

# Importování balíčků

Při instalaci Go jste nastavovali proměnnou prostředí GOPATH.

GOPATH je pracovní prostor pro jeden nebo více Go projektů.

GOPATH je kořenová složka pracovního prostoru a obsahuje tyto tři složky:

![Go Path](./gopath.png)

# Balíčky

Váš zdrojový kód, a kód závislostí je ve složce "src"

Když aplikaci zkompilujete, je umístěna do složky "bin".

Když zkompilujete jakoukoliv knihovnu, je umístěna do složky "pkg". V ní pak do subsložky podle architektury vašeho počítače. Například pkg/darwin_amd64.

Důležité je to zejména proto, že vaše GOPATH určuje jak go kompilátor řeší reference na balíčky ve vašem kódu.

# Importování balíčků

Pokud se váš kod nachází v $GOPATH/src/modra/cervena, název balíčku je "cervena" a jeho kód byste importovali za použití následujícího kódu:

	import "modra/cervena"

Identifikátor "modra/cervena" označujeme jako importní cestu balíčku.

Balíčky, které se nachází v repozitářích jako jsou github či bitbucket mají v importní cestě balíčku plnou cestu. Projekt v mém github repozitáři nazvaném "captainhook" má importní cestu

	"github.com/bketelsen/captainhook"

# Importování balíčků

To znamená, že aby bylo možné jej použít ve vašem kódu, balíček se MUSÍ nacházet v:

	$GOPATH/src/github.com/bketelsen/captainhook

Pokud by captainhook byla knihovnou, po kompilaci by se její zkompilovaná verze nacházela v:

	$GOPATH/pkg/darwin_amd64/github.com/bketelsen/captainhook
	(Za předpokladu, že byla zkompilována na Macu)

# Balíčky a GOPATH

Většina vývojářů používá pouze jednu GOPATH. Při instalaci ji nastaví jako proměnnou prostředí a už se jí nezabývají.

Je však možné mít "čístou louku" pro sady projektů nebo dokonce jednotlivé projekty, jednoduše vytvořením nové GOPATH a nastavení proměnné prostředí do této složky.

# Cvičení



Přečtěte si první půlku následujícího článku a poté si vyzkoušejte cvičení "Your First Program" (Můj první program) a "Your First Library" (Moje první knihovna)

[Začínáme s Go](https://golang.org/doc/code.html)

Článek navádí k nastavení GOPATH do $HOME/work... ignorujte to prosím. *NEMĚŇTE JIŽ NASTAVENOU GOPATH*



