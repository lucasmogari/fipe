package main

import (
	"fmt"
)

func main() {
//	marcas, err := ConsultarMarcas(182, 1)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	for _, marca := range marcas {
//		fmt.Println(marca.Label, ": ", string(marca.Value))
//	}
//	fmt.Println("###############################")
//
//	modelos, err := ConsultarModelos(182, 1, 1)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	for _, modelo := range modelos.Modelos {
//		fmt.Println(modelo.Label, ": ", string(modelo.Value))
//	}
//	for _, ano := range modelos.Anos {
//		fmt.Println(ano.Label, ": ", string(ano.Value))
//	}
//	fmt.Println("###############################")
//	
//	anos, err := ConsultarAnos(182, 1, 26, 6207)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	for _, ano := range anos {
//		fmt.Println(ano.Label, ": ", string(ano.Value))
//	}
//	fmt.Println("###############################")
//	
//	modelo, err := ConsultarModelo(182, 1, 26, 6207, 2014, 1)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Printf("%+v\n", modelo)
//	fmt.Println("###############################")
//	
//	modelo, err = ConsultarModeloPorCodigoFipe(182, 1, 2014, 1, "015089-4")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Printf("%+v\n", modelo)
//	fmt.Println("###############################")
	fmt.Println(ConsultarPeriodosReferencia())
}
