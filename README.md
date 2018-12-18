
# go-crash-course
[![License: mit](https://img.shields.io/badge/license-mit-green.svg?style=flat-square)](https://opensource.org/licenses/mit)

**Lições básicas sobre a Linguagem Go**
---
---

**Link's úteis**

[Go Code](https://golang.org/doc/code.html)

[Effective Go](https://golang.org/doc/effective_go.html)

[Go by exemples](https://gobyexample.com/)

[Go Design Patterns](https://www.godesignpatterns.com/)

[Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)


**Instalação**

Após instalar o Go ( [baixe aqui](https://golang.org/dl/) ) verifique o PATH utilizando o comando: 

```bash
$go env GOPATH
/users/renato/go
```

Os projetos devem ser feitos dentro da pasta, tradicionalmente os projetos devem ficar dentro de uma pasta chamada `src`, então, uma vez que você esteja dentro do GOPATH, crie a pasta **src**:
 
```bash
$go env GOPATH
/users/renato/go
$cd renato/go
$mkdir src
$cd src
```

dentro da pasta source você pode criar uma nova pasta para cada projeto. 

**Olá Mundo**
---
---

Vamos começar criando um projeto chamado OlaMundo e dentro dele um arquivo chamado ola.go:
```bash
$cd renato/go/src
$mkdir Olamundo
$cd Olamund
$touch ola.go
```
Abra o arquivo ola.go utilizando seu editor de texto favorito, vamos começar seu primeiro programa em Go:


```go
//define que este fonte será o pacote principal do projeto
package main

//importa o pacote fmt para utilizar função de printar dados no console
import (
	"fmt"
)

//Todo programa desenvolvido em Go deve possuir uma function chamada Main
//Main não recebe parametros e não possui retorno
func main() {
	fmt.Println("olá mundo!")
}
```

Para executar o arquivo utilize o comando go run nome.go, exemplo:
```bash
$cd renato/go/src/Olamundo
$go run ola.go
#o resultado será
olá mundo!
```

Você pode encotrar o arquivo [ola.go](https://code.engpro.totvs.com.br/renato.cunha/go-crash-course/src/branch/master/src/Olamundo/ola.go) dentro da pasta scr deste projeto.


**Variáveis**
---
---

Em Go podemos definir variáveis de duas formas, explicitando diretamente sua tipagem, ou atribuindo um valor diretamente: 

*tipando*
```go
package main

import (
	"fmt"
)

func main() {
	var a int = 1 //Variável a é um inteiro com conteúdo igual a 1
    var b int = 2 //Variável b é um inteiro com conteúdo igual a 2
    var c int = a + b //Variável c é um inteiro com o resultado da soma de A e B

    fmt.Println(c)
}
```

Podemos declarar as mesmas variáveis de uma forma mais simples, atribuindo diretamente os valores: 
```go
package main

import (
	"fmt"
)

func main() {
	a := 1
    b := 2
    c := a + b

    fmt.Println(c)
}
```

Ambas formas possuem o mesmo retorno. 

Arrays também devem ser tipados e também devem ter seu tamanho pré definido:
```go
package main

import (
	"fmt"
)

func main() {
	//caso os elementos não sejam iniciados, seus valores serão os básicos para seus tipos: 0 para inteiros, '' para string e false para bool.
    var arra1[10] int
    

}
```

