# Ukazatele a typy reference

Array (pole) má pevnou délku a pevný datový typ.

- Příklad array

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/array
	go run main.go
	```


# Iterace

Polem iterujeme použitím cyklu for

- Iterování s for

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/array
	go run main.go
	```


For je jediný druh cyklu v Go. Používejte jej místo FOR, WHILE, DO WHILE, DO UNTIL, atd.

# Použití rozsahu (Range)

Iterovat lze nad jakoukoliv kolekcí s použitím příkazu range. Range je zabudovaná iterační funkce, která vrací index a hodnotu mnoha různých typů kolekcí, včetně array:

- Iterování s Range

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/array
	go run main.go
	```

Range je velmi užitečný příkaz, který budete využívat často.

# Cvičení

Upravte tento příklad aby iteroval nad array s použitím příkazu range a zobrazte hodnotu vynásobenou 2.

	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/exercises/range

nápověda: násobení proveďte uvnitř funkce fmt.Println().

# Slices


Zatímco array má vždy fixní délku, Slice má dynamickou délku, a je tak flexibilnější než array. V praxi spíše používáme slice než array.

Pokud nevíte že máte konečný a pevný seznam prvků, skoro vždy budete při práci s daty používat slice.

- Slice příklady

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/slice
	go run main.go
	```


# Proč slice?

- Každá změna array alokuje nové array. Neefektivní!
- Slice lze změnit bez alokování
- S částí array i slice lze jednoduše pracovat

Jak poznat rozdíl mezi slice a array?

Slices není deklarován s délkou:

	var sl []string // slice
	var ar [5]string // array

# Přidání prvku do slice

Přidání prvků jsme si již ukázali, slice lze rovnou definovat i včetně všech jeho hodnot:

- Inline Deklarace slice

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/slice
	go run main.go
	```


# Řetězce

Řetězce jsou slice bajtů

Řetězec je vždy jen slice bajtů. Go má zabudovanou podporu pro UTF-8 a výkonné nástroje pro práci s ne-ascii znaky. Zatímco ASCII znaky zabírají pouze jeden bajt, UTF-8 znaky (neboli Runy) můžou zabírat až 4 bajty. Go umí s vícebajtovými runami jednoduše pracovat.

- Příklad Run

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/runes
	go run main.go
	```

# Mapy

Mapa je _neseřazená_ sada hodnot indexovaných unikátním klíčem.

Mapy musí být před použitím inicializovány.

Mapu inicializujeme přes make() funkci.

- Příklad Mapy

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/maps
	go run main.go
	```


# Mapy

Klíče map musí být porovnatelné přes "==" a "!=". Jako klíč tedy nelze použít funkce, jiné mapy nebo slice.

- Mapy lze také deklarovat včetně jejich hodnot:

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/mapsinline
	go run main.go
	```

# Konkurenčnost map

Mapy NEJSOU threadsafe (vláknově bezpečné). Není v pořádku číst nebo modifikovat mapu v dvou různých goroutinách. Na tohle se ještě podíváme v lekcí o konkurenčnosti. Zatím si jen zapamatujte, že pokud máte mapu, která je používána na více místech, přístup k ní by měl být synchronizovaný.

Do verze GO 1.7 konkurenční přístup k mapě způsobil nestabilitu programu. Od Go verze 1.7 a výše způsobí konkurenční přístup k mapě pád programu.

 -- TODO -- sync.Map

# Ukazatele

V Go můžete jako parametry funkce předávat hodnoty nebo ukazatele. Obecně je dobré předávat jako hodnoty, pokud jsou jen dočasné a nebudou používány po zavolání funkce.

- Příklad předání hodnoty

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/passbyval
	go run main.go
	```

Všimněte si, že jsme předanou hodnotu nezměnili. Vytvořili jsme nové číslo a to vrátili.

Pokud chcete hodnotu měnit a tyto změny zachovat, předejte je za použití ukazatele:

- Příklad předání ukazatele

    ```
	cd $GOPATH/src/github.com/thewondertwins/learngo/material/pointersreftypes/demos/passbyref
	go run main.go
	```

# Dereferencování ukazatele

Operace s ukazateli jsou podobné jako v rodině jazyků C, nelze však nad nimi dělat aritmetické operace.

Použitím & získáte adresu proměnné.

Použitím * dereferencujete proměnnou.







