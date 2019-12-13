# OK! AI.Censor Go Client

Podstawowa implementacja biblioteki do Golang która odpytuje OK! AI.Censor

## Przykładowe użycie

```go
package main

import (
	okaeric "github.com/okaeripoland/ai-censor-go-client"
	"log"
)

func main() {
	// dane uwierzytelniające
	token := "TOKEN"

	// tworzymy klienta (w większości przypadków należy go gdzieś zapisać)
	censor, err := okaeric.CreateClient(token, true /* true dla logów podczas testów, false dla produkcji */)
	if err != nil {
		log.Printf("%v", err)
		return
	}

	// zapytanie o przewidywanie
	phrase := "ale z niego !@#%"
	res, err := censor.Predict(phrase)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	log.Printf("%v", res)

	// czy jest to wulgarne?
	swear := res.General.Swear
	if swear {
		log.Printf("Fraza '%s' zostala uznana za wulgarna.", phrase)
	} else {
		log.Printf("Fraza '%s' nie zostala uznana za wulgarna.", phrase)
	}
}

```