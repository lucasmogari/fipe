package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

//const (
//	Carro    = 1
//	Moto     = 2
//	Caminhao = 3
//)
//
//const (
//	Gasolina = 1
//	Alcool   = 2
//	Diesel   = 3
//)
//

var client = &http.Client{}

type Entry struct {
	Label string
	// Value pode ser int ou string
	Value json.RawMessage
}

type Marca Entry

type Ano Entry

type ModelosMarca struct {
	Modelos []Entry
	Anos    []Entry
}

type Modelo struct {
	Valor            string
	Marca            string
	Modelo           string
	AnoModelo        int
	Combustivel      string
	CodigoFipe       string
	MesReferencia    string
	Autenticacao     string
	TipoVeiculo      int
	SiglaCombustivel string
	DataConsulta     string
}

func ConsultarMarcas(codigoReferencia, tipoVeiculo int) (marcas []Marca, err error) {
	data := &url.Values{}
	data.Add("codigoTabelaReferencia", strconv.Itoa(codigoReferencia))
	data.Add("codigoTipoVeiculo", strconv.Itoa(tipoVeiculo))

	return marcas, unmarshal("http://www.fipe.org.br/IndicesConsulta-ConsultarMarcas", data, &marcas)
}

func ConsultarModelos(codigoReferencia, tipoVeiculo, codigoMarca int) (modelos ModelosMarca, err error) {
	data := &url.Values{}
	data.Add("codigoTabelaReferencia", strconv.Itoa(codigoReferencia))
	data.Add("codigoTipoVeiculo", strconv.Itoa(tipoVeiculo))
	data.Add("codigoMarca", strconv.Itoa(codigoMarca))

	return modelos, unmarshal("http://www.fipe.org.br/IndicesConsulta-ConsultarModelos", data, &modelos)
}

func ConsultarModelosDoAno(codigoReferencia, tipoVeiculo, codigoMarca, anoModelo int) (modelos ModelosMarca, err error) {
	data := &url.Values{}
	data.Add("codigoTabelaReferencia", strconv.Itoa(codigoReferencia))
	data.Add("codigoTipoVeiculo", strconv.Itoa(tipoVeiculo))
	data.Add("codigoMarca", strconv.Itoa(codigoMarca))
	data.Add("anoModelo", strconv.Itoa(anoModelo))

	return modelos, unmarshal("http://www.fipe.org.br/IndicesConsulta-ConsultarModelosAtravesDoAno", data, &modelos)
}

func ConsultarAnos(codigoReferencia, tipoVeiculo, codigoMarca, codigoModelo int) (anos []Ano, err error) {
	data := &url.Values{}
	data.Add("codigoTabelaReferencia", strconv.Itoa(codigoReferencia))
	data.Add("codigoTipoVeiculo", strconv.Itoa(tipoVeiculo))
	data.Add("codigoMarca", strconv.Itoa(codigoMarca))
	data.Add("codigoModelo", strconv.Itoa(codigoModelo))

	return anos, unmarshal("http://www.fipe.org.br/IndicesConsulta-ConsultarAnoModelo", data, &anos)
}

func ConsultarAnosPorCodigoFipe(codigoReferencia, tipoVeiculo int, codigoFipe string) (anos []Ano, err error) {
	data := &url.Values{}
	data.Add("codigoTabelaReferencia", strconv.Itoa(codigoReferencia))
	data.Add("codigoTipoVeiculo", strconv.Itoa(tipoVeiculo))
	data.Add("modeloCodigoExterno", codigoFipe)

	return anos, unmarshal("http://www.fipe.org.br/IndicesConsulta-ConsultarAnoModeloPeloCodigoFipe", data, &anos)
}

func ConsultarModelo(codigoReferencia, tipoVeiculo, codigoMarca, codigoModelo, anoModelo, codigoTipoCombustivel int) (modelo Modelo, err error) {
	data := &url.Values{}
	data.Add("codigoTabelaReferencia", strconv.Itoa(codigoReferencia))
	data.Add("codigoTipoVeiculo", strconv.Itoa(tipoVeiculo))
	data.Add("codigoMarca", strconv.Itoa(codigoMarca))
	data.Add("codigoModelo", strconv.Itoa(codigoModelo))
	data.Add("anoModelo", strconv.Itoa(anoModelo))
	data.Add("codigoTipoCombustivel", strconv.Itoa(codigoTipoCombustivel))

	return modelo, unmarshal("http://www.fipe.org.br/IndicesConsulta-ConsultarValorComTodosParametros", data, &modelo)
}

func ConsultarModeloPorCodigoFipe(codigoReferencia, tipoVeiculo, anoModelo, codigoTipoCombustivel int, codigoFipe string) (modelo Modelo, err error) {
	data := &url.Values{}
	data.Add("codigoTabelaReferencia", strconv.Itoa(codigoReferencia))
	data.Add("codigoTipoVeiculo", strconv.Itoa(tipoVeiculo))
	data.Add("anoModelo", strconv.Itoa(anoModelo))
	data.Add("codigoTipoCombustivel", strconv.Itoa(codigoTipoCombustivel))
	data.Add("modeloCodigoExterno", codigoFipe)
	data.Add("tipoConsulta", "codigo")

	return modelo, unmarshal("http://www.fipe.org.br/IndicesConsulta-ConsultarValorComTodosParametros", data, &modelo)
}

func unmarshal(url string, data *url.Values, result interface{}) error {
	body, err := post(url, data)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &result)
}

func post(url string, data *url.Values) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data.Encode()))
	req.Header.Set("Referer", "http://www.fipe.org.br/pt-br/indices/veiculos/")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
