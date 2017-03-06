// la cartella diventa il nome del package.
package stringutil

//la funzione sarà visbile all'esterno perchè inizia scritta in maiuscolo.
//per questa ragione è "exported", che è come public in altri linguaggi.
// però è visibile anche se minuscola all'interno di altri file interni al packages/cartella

func Reverse(s string) string {
	return reverseTwo(s)
}
