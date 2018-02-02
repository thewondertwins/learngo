# Empezando

## Instalando Go (Mac o Linux)

Descargue e instale Go - siempre utilice los paquetes hospedados en golang.org - nunca use homebrew o apt-get, yum, etcétera. Están rotos o peor -- modificados por alguien más.

Defina GOPATH in .bashrc, .bash_profile, .zshrc etcétera:

	export GOPATH=$HOME/go

Agregue los binarios de Go (compiladores y herramientas) al _path_:

	export PATH=$PATH:/usr/local/go/bin

Cierre la sesión e ingrese de nuevo para que los cambios tomen efecto o

	$ source .bashrc

para recargarlos al instante.

## Instalando Go (Windows)

Descargue e instale Go de golang.org - Seleccione el instalador MSI

Defina GOPATH en Variables de Entorno

	GOPATH=%userdir%/go

Agregue los binarios de Go (compiladores y herramientas) al _path_:

	%userdir%/go/bin

## Verificando la instalación

Desde la línea de comandos:

	go version

Algo como lo siguiente debería aparecer:

	go version 1.8 linux/amd64

## Editando código en Go

Editores más populares de Go:

vim y neovim con el plugin vim-go

emacs con go-mode.el

Visual Studio Code con vscode-go (¡el depurador funciona!) 

Atom con go-plus

IntelliJ IDEA con el plugin para Go

GoLand

## Go Playground

Incluso si no cuenta con un editor configurado localmente usted aun puede utilizar Go con navegador.

[Go Playground](https://play.golang.org)

Go Playground es un servicio web que utiliza los servidores de golang.org. Este servicio recibe un programa escrito en Go, lo compila, lo enlaza y lo ejecuta dentro de un _sandbox_, regresando al final el resultado.

## Limitaciones del Playground

Las siguientes son las limitaciones de los programas ejecutados en el playground:

El playground puede usar casi toda la librería estándar, con algunas excepciones. La única comunicación que un programa en el playground tiene al mundo exterior es escribiendo a la salidas estándard stdout y stderr.

En el playground el tiempo comienza en 2009-11-10 23:00:00 UTC (determinar el significado de esta particular fecha es un ejercicio para el lector). Esto hace más sencillo cachear programas dándoles salidas determinísticas.

También hay límites en tiempo de ejecución además de uso de CPU y memoria.

Por lo tanto: no entrada/salida de archivos, nada útil con tiempo y fechas, tampoco se pueden usar paquetes externos.

## El Go Playground

Inclusive con todas esas limitantes programadores en Go disfrutan el Playground - es un lugar genial para compartir código, inclusive si no puede ejecutarse o compilarse. Usted puede entrar código y compartirlo (dándole click al botón "SHARE") con una URL permanente.

Pruebelo ahora con este enlace:

[Hola Mundo!](https://play.golang.org/p/jDQ5ap1c_ZD)

## El comando Go

Toda la interacción con la linea comandos en Go sera a través del comando `go`.

Algunos comandos comunes:

	go build algun/paquete
	go run main.go
	go test algun/paquete
	go get github.com/algun/paquete
	go install algun/paquete

## Ejercicio

Desde la línea de comandos escriba `go` y presione entrar para mirar las diferentes herramientas el comando `go` implementa. Pruebe algunas como:

	go env
	go list
	go version

## Material de descarga

Desde la línea de comandos utilice el comando `go` para descargar los materiales y ejercicios de este libro:

	go get github.com/thewondertwins/learngo
