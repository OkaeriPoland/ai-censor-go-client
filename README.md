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

## Opis zwracanych własności

#### PredictResponse

| Własność  | Opis |
| ------------- | ------------- |
| General General | Sekcja ogólna odpowiedzi |
| Details Details | Sekcja szczegółów odpowiedzi |
| Elapsed Elapsed | Sekcja informacji dotyczących czasu przetwarzania |


#### General

| Własność  | Opis |
| ------------- | ------------- |
| Swear bool | Informacja o tym, czy wiadomość została uznana za wulgarną |
| Breakdown string | Przetworzona wiadomość ułatwiająca ewentualne debugowanie błędnych wykryć, przydatna do wyświetlania dla administracji w logach |


#### Details

| Własność  | Opis |
| ------------- | ------------- |
| BasicContainsHit bool | Informacja o tym, czy wiadomość zawierała zakazane frazy |
| ExactMatchHit bool | Informacja o tym, czy wiadomość była zablokowaną frazą (np. wyrażenie jd) |
| AILabel string | Ocena ai (`ok` lub `swear`) |
| AIProbability float64 | Wartość od `0` do `1` określająca prawdopodobieństwo dotyczące prawdziwości `aiLabel` |


#### Elapsed 

| Własność  | Opis |
| ------------- | ------------- |
| All float64 | Całkowity czas w milisekundach przez który zapytanie było obsługiwane wewnętrznie |
| Processing float64 | Czas przez jaki zostały wykonane oceny wulgarności |
