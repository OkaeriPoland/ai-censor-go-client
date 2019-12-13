# OK! AI.Censor Java Client

Podstawowa implementacja biblioteki do Golang która odpytuje OK! AI.Censor

## Przykładowe użycie

```go
package main

import (
	okaeric "github.com/okaeripoland/ai-censor-go-client"
    "log"
)

func main(){
    wiadomosc := "ale z niego !@#%"
    // true dla logów podczas testów, false dla produkcji
    okC, err := okaeric.CreateClient("TOKEN",true)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	res, err := okC.Predict(wiadomosc)
	if err != nil {
		log.Printf("%v", err)
		return
	}
    log.Printf("%v",res)
}
```