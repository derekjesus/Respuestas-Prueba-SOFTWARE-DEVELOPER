package main

import (
	"fmt"
	"strings"
)

func encriptar(clave, mensaje string) string {
	if mensaje == "" {
		return ""
	}
	if clave == "" {
		clave = "DCJ"
	}

	vocales := "aeiouAEIOU"
	var resultado strings.Builder

	for _, c := range mensaje {
		if strings.ContainsRune(vocales, c) {
			resultado.WriteString(clave)
		}
		resultado.WriteRune(c)
	}

	return resultado.String()
}

func main() {
	clave := "dcj"
	mensaje := "I love prOgrAmming!"
	fmt.Println(encriptar(clave, mensaje)) // dcjI ldcjovdcje prdcjOgrdcjAmmdcjing!
}
