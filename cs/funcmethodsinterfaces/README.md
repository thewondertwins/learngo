# Funkce

Pro definici funkce používáme klíčové slovo "func".

Funkce mají jméno, volitelně vstupní parametry, a volitelně návratové hodnoty.

## Příklad funkce

	cd $GOPATH/src/github.com/gophertrain/material/funcmethodsinterfaces/demos/funcs
	go run main.go

Funkce v Go jsou First Class objekty. Lze je tedy přiřadit do proměnné či je předávat jako parametry.  

## Argumenty funkcí

	cd $GOPATH/src/github.com/gophertrain/material/funcmethodsinterfaces/demos/funcvalues
	go run main.go

# Metody

Metody jsou syntaktickou berličkou pro práci s typem.

CO?

	func ChangeEmail(u *User, newEmail string) { ... } // Ošklivé řešení

	func (u *User) ChangeEmail(newEmail string) { ... } // Čisté řešení

Obě výše uvedené definice jsou použitelné, ale jedna z nich je daleko čitelnější.

Metody lze přiřadit kterémukoliv pojmenovanému typu. Pokud chcete upravovat již existující typ, předávajte ukazatel. Pokud jej nechcete modifikovat, předávejte hodnotu. 

## Příklad metody

	cd $GOPATH/src/github.com/gophertrain/material/funcmethodsinterfaces/demos/method
	go run main.go


Metody v Go jsou first-class Citizens.

To znamená, že je můžete přiřadit do proměnné a nad ní provádět operace.


## First Class Metody 

	cd $GOPATH/src/github.com/gophertrain/material/funcmethodsinterfaces/demos/firstmethod
	go run main.go

# Interfaces (Rozhraní)

Interface vám umožní specifikovat CHOVÁNÍ.

Pokud něco má určité vlastnosti, můžeme je použít.

Interface jsou Typy. Definujeme je jako Typy.

Interface zpravidla neobsahuje velké množství metod, 1 či 2 jsou nejběžnější.

	Čím větší je interface, tím slabší je abstrakce. -- Rob Pike

Pojmenování Interface by mělo popisovat jejich činnost.

# Příklad Interface

Stringer - Typ, který má metodu vracející řetězec

io.Writer - Typ, který má metodu zapisující do bufferu

io.ReadCloser - Typ, který má metodu, která čte data ze streamu a po skončení jej zavře

# Vytváření kvalitních Interface

Kvalitní Interface definuje pouze malý počet specifických akcí:

- Zápis bytů do bufferu (io.Writer)
- Vrácení řetězce, který daný typ reprezentuje (fmt.Stringer)

# Interface ze standardní knihovny

Příklady Interface ze standardní knihovny Go:

[db/sql: Interface připojení](https://golang.org/pkg/database/sql/driver/#Conn)

[encoding: Interfaces překladače](https://golang.org/pkg/encoding/)

[net/http: Interface zpracování HTTP](https://golang.org/pkg/net/http/#Handler)

Z těchto příkladů lze vidět, že Interface by měly obsahovat pouze malé sady definice chování.


# Interface v praxi

## Příklad Interface

	cd $GOPATH/src/github.com/gophertrain/material/funcmethodsinterfaces/demos/interfaces
	go run main.go

## Prázdný Interface

Každý Interface, který jsme doposud viděli deklaroval jednu nebo více funkcí. To ovšem není nutnost.

Pokud definujete Inteface bez funkcí, lze jej pak použít pro jakýkoliv Typ.

V Go používáme prázdný interface pro reprezantaci "čehokoliv".

## Příklad prázdného Interface

	cd $GOPATH/src/github.com/gophertrain/material/funcmethodsinterfaces/demos/empty
	go run main.go

## Předpokládaná hodnota

Tohle může být zrádné. Jakýkoliv Typ můžete předávat bez obav o ztrátu typu. Ale pokud dostanete "interface{}" jak máte vědět co s hodnotou dělat? 

## Příklad předpokládání hodnoty

	cd $GOPATH/src/github.com/gophertrain/material/funcmethodsinterfaces/demos/assert
	go run main.go

# Cvičení

Doplňte poslední příklad, aby uměl rozpoznat a vypsat float hodnotu.

	cd $GOPATH/src/github.com/gophertrain/material/funcmethodsinterfaces/exercises/assert
	go run main.go
