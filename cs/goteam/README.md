# Programujte jako autoři Go

V této kapitole knihy se naučíte jak psát kód, který vypadá jako by patřil do standardní knihovny Go, a proč je to tak důležité.  

Témata:

- jak organizovat kód do balíčků a co by takový balíček měl obsahovat  
- styl zápisu a konvence, které ve standardní knihovně převažují  
- jak psát čistý a dobře čitelný kód
- nepsané Go zvyklosti, které jdou dál než použití “go fmt” a díky kterým budete vypadat jako ostřílený přispěvatel do Go


# Obsah první části

V této části se zaměříme na tři hlavní oblasti, díky kterým bude váš kód vypadat profesionálněji, bude rychlejší a půjde jej snáze číst a chápat.

- Balíčky
- Jmenné konvence
- Konvence zdrojového kódu

# Balíčky

## Organizace kódu

Organizace kódu je oblast, kde často tápou i ostřílení Go vývojáři.

Dvě konkrétní oblasti organizace kódu v go, které mají obrovský dopad na použitelnost, testovatelnost a funkčnost kódu jsou:

- Pojmenování balíčků
- Organizace balíčků

Nelze o nich mluvit odděleně, protože se navzájem ovlivňují.

## Organizace balíčků - Knihovny

Začněme s Go samotným jako naší výchozí inspirací. Seznam nejvyšší úrovně balíčků ve standardní knihovně Go 1.7:

	archive   cmd       crypto    errors    go        index     math      path      sort      syscall   unicode
	bufio     compress  database  expvar    hash      internal  mime      reflect   strconv   testing   unsafe
	builtin   container debug     flag      html      io        net       regexp    strings   text      vendor
	bytes     context   encoding  fmt       image     log       os        runtime   sync      time

Okamžitě je nám jasné

	Kód balíčku má jednoznačný účel 

Pokud skupina balíčků poskytuje sadu funkcionalit s různou implementací, řadí se pod svého předka. Podívejme se na obsah balíčku *encoding*

	ascii85     base32      binary      encoding.go hex         pem
	asn1        base64      csv         gob         json        xml

Balíček *encoding* však obsahuje pouze jeden soubor *encoding.go*, který definuje interface pro implementace svých potomků.


Díky použití této struktury pojmenování je jasné, že *encoding/json* je balíček, který se stará o práci s formátem *json*. 

Společné prvky:

- Jméno balíčku popisuje jeho účel
- Podle názvu lze jednoduše pochopit co balíček dělá
- Názvy jsou spíše krátké
- Pokud je to potřeba, použijeme popisný nadřazený balíček a několik potomků implementujících skutečnou funkcionalitu -- jako u balíčku *encoding*


## Organizace balíčků - Aplikace

Všechny balíčky, které jsme viděli jsou knihovny. Předpokládá se, že budou importovány a použity ve spustitelné programu jako je služba nebo nástroj příkazové řádky.

Jak by měl být organizován kód vaší spustitelné aplikace?

Pokud pracujete s aplikací, organizace balíčku je lehce odlišná. Rozdíl je v tom, že aplikace všechny balíčky spojí dohromady.

Organizace balíčku aplikace má velký dopad na testovatelnost a funkčnost vašeho systému.

	Když píšete aplikaci, měli byste psát kód, který je lehce pochopitelný, jednoduše refaktorovatelný a snadný na údržbu pro někoho dalšího.

Většina knihoven se soustředí na to aby poskytla jednoznačně zaměřenou funcionalitu; logování, překlad, přístup k síti.

Vaše aplikace tyto knihovny spojí dohromady, čímž vznikne nástroj nebo služba. Výsledek však bude pokrývat daleko větší rozsah funkcionality.  

Když vytváříte aplikaci, měli byste svůj kód organizovat do balíčků, ty by se měly zaměřovat na dvě oblasti:

- Doménové Objekty
- Služby

*Doménové objekty* jsou typy, které definují vaši business logiku.

*Služby* jsou balíčky, které pracují s nebo na doménových objektech.

[Rozvržení standardního balíčku od Bena Johnsona](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)

Doménový objekt je základ vaší aplikace. Pokud vytváříte skladový systém, vaše doménové objekty budou *Produkt* a *Dodavatel*. Pokud vytváříte systém pro HR, vaše doménové objekty budou *Zaměstnanec*, *Oddělení* a *Pobočka*.

Balíček obsahující doménové objekty by měl také deklarovat interface (rozhraní) mezi doménovými objekty a zbytkem světa. Ty definují věci, které s doménovými objekty chcete dělat.

- *ProductService* (služba k produktu)
- *SupplierService* (služba k dodavateli)
- *AuthenticationService* (authentizační služba)
- *EmployeeStorage* (úložiště zaměstnanců)
- *RoleStorage* (úložiště rolí)

Balíček s doménovými objekty by měl být základ vašeho aplikačního repozitáře. Tím zajístíte, že každý kdo bude procházet vaše kódy okamžitě ví, s jakými objekty se pracuje a jaké se nad nimi provádí operace.

Balíček s doménovými objekty neboli *základní* balíček by něměl mít žádné další externí závislosti. Jeho jediný účel je popis vašich objektů a jejich chování.

	Implementace vaší doménové logiky by měla být v odděleném balíčku se závislostí na *základním* balíčku.

Mezi závilosti mimo jiné patří:

- Externí datové zdroje
- Logika přenosu dat (http, RPC)

Pro každou závislost by měl existovat vlastní balíček.

Proč samostatný balíček pro každou závislost?

- Jednodušší testování
- Snadná náhrada/výměna
- Vyhneme se tak kruhovým závislostem

Důsledkům takové organizace balíčků se budeme dále věnovat v kapitolách *Interfaces (rozhraní)* a *Testování*.

# Konvence

## Jmenné konvence

    Existují dva nejnáročnější problémy v programování
    a to správná invalidace cache, pojmenovávání věcí a chyby o jedničku.
*Zdroj*: Každý vývojář na Twitteru


Pojmenování věcí *je* složité, ale čím více času věnujete správnému pojmenování vašich objektů, funkcí a balíčku, tím čitelnější váš kód bude.


## Jmenné konvence - Balíčky

Název balíčku by měl mít tyto vlastnosti:

- je krátký

	Raději "transport" než "transportmechanisms" (transportní mechanismus)

- jasný

    Jméno objasňující funkci balíčku: "bytes"
    Jméno popisující implementaci externí závislosti: "postgres"

Balíčky by měly mít jednoznačnou funkcionalitu. Vyhněte se *všeřešícím* balíčkům:

	util (nástroje)
	helpers (pomocné funkce)
	atd.

Často je to pouze znamením, že vám někde chybí interface.

    util.ConvertOtherToThing() (změň tohle na věc) by měl být refaktorován na interface pro *Věc*

*všeřešící* balíčky budou pravděpodobně prvním místem, kde začnete narážet na problémy při testování či kruhových závislostech.

## Jmenné konvence - Proměnné

Mezi běžné konvence pro pojmenování proměnných patří:

- používejte camelCase ne snake_case
- pro indexy používejte jednopísmenné proměnné
	- for i:=0; i < 10; i++ {}
- používejte krátké ale popisné názvy pro pojmenování všeho ostatního
	- var count int (var pocet int)
	- var cust Customer (var zak Zakaznik)

Používáním zbytečně krátkých jmen pro proměnné nic nezískáte. K pojmenování můžete použít jednoduchý princip:

    Čím dále od deklarace používáte proměnnou tím delší (popisnější) by její název měl být.

Zde pár příkladů pojmenování proměnné na základě kontextu:

.link https://github.com/golang/go/blob/master/src/fmt/format.go#L12
.link https://github.com/golang/go/blob/master/src/fmt/format.go#L326
.link https://github.com/golang/go/blob/master/src/flag/flag.go#L293

- pro pojmenování pole prvků použijte opakující se písmeno
	- var tt []*Thing (var vv []*Věc)

- uvnitř cyklu používejte jednopísmenné proměnné
	- for i, t := range tt {}

Tyto konvence se běžně používají i v samotném zdrojovém kódu Go.

## Jmenné konvence - funkce a metody

Funkcním v balíčku nepřidávejte do názvu jméno balíčku:

	SPRÁVNĚ:  log.Info()
	ŠPATNĚ:   log.LogInfo()

Samotný název balíčku definuje jeho účel, není třeba jej opakovat.

Go kód neobsahuje settery ani gettery.

	SPRÁVNĚ:  custSvc.Customer()
	ŠPATNĚ:   custSvc.GetCustomer()

## Jmenné konvence - Interfaces (rozhraní)

Pokud váš interface obsahuje pouze jednu funkci, připojte za její jméno "-er":

	type Stringer interface{ //Textovač
		String() string //Text
	}

Pokud má váš interface více funkcí, pojmenujte ho tak aby popisoval svou funkcionalitu:

	type CustomerStorage interface { //zákaznické úložiště
		Customer(id int) (*Customer, error) //zákazník
		Save(c *Customer)  error //uložit
		Delete(id int) error //smazat
	}

Puritáni by mohli namítnout, že všechny interface by měli mít koncovku `-er`. Dle mého názoru je popisnost a čitelnost přednější.

## Jmenné konvence - Zdrojový kód

Uvnitř balíčku oddělujte kód do logických celků.

Pokud balíček pracuje s více objekty, ponechte logiku pro každý objekt ve vlastním souboru:

	src/goteam/cviceni/sklad/postgres

	objednavky.go
	dodavatele.go
	produkty.go

V balíčku, který definuje vaše doménové objekty, definujte typy a jejich interface (rozhraní) pro každý objekt v odděleném souboru:

	src/goteam/cviceni/sklad

	objednavky.go -- Obsahuje typ Order (objednavka) a interface (rozhraní) OrderStorage (úložiště objednávek)


# Dodatečné tipy

- Komentáře piště jako vždy jako celé věty.

	```
	// Objednávka reprezentuje objednávku od zákazníka
	type Order struct {}
	```
- Používejte `goimports` na správu vašich importů aby byli ve správném pořadí. Nejprve standardní knihovny, poté externí.

- Vyhněte se používání konstrukce `else`. Zvláště pak při práci s chybami.
	```
	if err != nil {
		// práce s chybou
		return // nebo continue, atd.
	} 
	// normální kód
	```

# Pište kód jako autoři Go

Budete-li dodržovat tyto konvence, váš kód bude lépe čitelný, snáze udržitelný a kdokoliv další jej rychleji pochopí.


# Cvičení - Samostatně

V *$GOPATH/src/github.com/thewondertwins/learngo/material/goteam/exercises/inventory/* je kostra aplikace včetně všech interface (rozhraní).

*Úkol:*

Vytvořte odpovídající strukturu balíčku pod složkou _inventory_ předpokládající:

- Databáze pro uložení dat bude postgres
- Dodavatel poskytující SupplierService (službu dodavatele) bude "Acme")
- Dvě služby pro přenos dat: REST skrze HTTP a RPC přes TCP

Stačí pouze vytvořit složky reprezentující vaše balíčky. Žádný kód není třeba psát.

*Očekáváný výstup:*

Porozumění důsledkům závislostí při vytváření balíčků
