package main

import (
	"bytes"
	"encoding/json"
	"github.com/lucasmogari/fipe/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"github.com/lucasmogari/fipe/Godeps/_workspace/src/golang.org/x/net/html"
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

type PeriodoReferencia struct {
	Label string
	Codigo string
}

func ConsultarPeriodosReferencia() ([]PeriodoReferencia, error) {
	doc, err := goquery.NewDocument("http://www.fipe.org.br/pt-br/indices/veiculos")
	if err != nil {
		return nil, err
	}

	options := doc.Find("select#selectTabelaReferenciacarro option")
	periodos := make([]PeriodoReferencia, 0, options.Length())

	options.Each(func(i int, s *goquery.Selection) {
		value, exists := s.Attr("value")
		if exists {
			periodos = append(periodos, PeriodoReferencia{s.Text(), value})
		}
	})
	return periodos, nil
}

func ConsultarMarcas(codigoPeriodoReferencia, tipoVeiculo int) (marcas []Marca, err error) {
	data := &url.Values{}
	data.Add("codigoTabelaReferencia", strconv.Itoa(codigoPeriodoReferencia))
	data.Add("codigoTipoVeiculo", strconv.Itoa(tipoVeiculo))

	return marcas, unmarshal("http://www.fipe.org.br/IndicesConsulta-ConsultarMarcas", data, &marcas)
}

func ConsultarModelos(codigoPeriodoReferencia, tipoVeiculo, codigoMarca int) (modelos ModelosMarca, err error) {
	data := &url.Values{}
	data.Add("codigoTabelaReferencia", strconv.Itoa(codigoPeriodoReferencia))
	data.Add("codigoTipoVeiculo", strconv.Itoa(tipoVeiculo))
	data.Add("codigoMarca", strconv.Itoa(codigoMarca))

	return modelos, unmarshal("http://www.fipe.org.br/IndicesConsulta-ConsultarModelos", data, &modelos)
}

func ConsultarModelosDoAno(codigoPeriodoReferencia, tipoVeiculo, codigoMarca, anoModelo int) (modelos ModelosMarca, err error) {
	data := &url.Values{}
	data.Add("codigoTabelaReferencia", strconv.Itoa(codigoPeriodoReferencia))
	data.Add("codigoTipoVeiculo", strconv.Itoa(tipoVeiculo))
	data.Add("codigoMarca", strconv.Itoa(codigoMarca))
	data.Add("anoModelo", strconv.Itoa(anoModelo))

	return modelos, unmarshal("http://www.fipe.org.br/IndicesConsulta-ConsultarModelosAtravesDoAno", data, &modelos)
}

func ConsultarAnos(codigoPeriodoReferencia, tipoVeiculo, codigoMarca, codigoModelo int) (anos []Ano, err error) {
	data := &url.Values{}
	data.Add("codigoTabelaReferencia", strconv.Itoa(codigoPeriodoReferencia))
	data.Add("codigoTipoVeiculo", strconv.Itoa(tipoVeiculo))
	data.Add("codigoMarca", strconv.Itoa(codigoMarca))
	data.Add("codigoModelo", strconv.Itoa(codigoModelo))

	return anos, unmarshal("http://www.fipe.org.br/IndicesConsulta-ConsultarAnoModelo", data, &anos)
}

func ConsultarAnosPorCodigoFipe(codigoPeriodoReferencia, tipoVeiculo int, codigoFipe string) (anos []Ano, err error) {
	data := &url.Values{}
	data.Add("codigoTabelaReferencia", strconv.Itoa(codigoPeriodoReferencia))
	data.Add("codigoTipoVeiculo", strconv.Itoa(tipoVeiculo))
	data.Add("modeloCodigoExterno", codigoFipe)

	return anos, unmarshal("http://www.fipe.org.br/IndicesConsulta-ConsultarAnoModeloPeloCodigoFipe", data, &anos)
}

func ConsultarModelo(codigoPeriodoReferencia, tipoVeiculo, codigoMarca, codigoModelo, anoModelo, codigoTipoCombustivel int) (modelo Modelo, err error) {
	data := &url.Values{}
	data.Add("codigoTabelaReferencia", strconv.Itoa(codigoPeriodoReferencia))
	data.Add("codigoTipoVeiculo", strconv.Itoa(tipoVeiculo))
	data.Add("codigoMarca", strconv.Itoa(codigoMarca))
	data.Add("codigoModelo", strconv.Itoa(codigoModelo))
	data.Add("anoModelo", strconv.Itoa(anoModelo))
	data.Add("codigoTipoCombustivel", strconv.Itoa(codigoTipoCombustivel))

	return modelo, unmarshal("http://www.fipe.org.br/IndicesConsulta-ConsultarValorComTodosParametros", data, &modelo)
}

func ConsultarModeloPorCodigoFipe(codigoPeriodoReferencia, tipoVeiculo, anoModelo, codigoTipoCombustivel int, codigoFipe string) (modelo Modelo, err error) {
	data := &url.Values{}
	data.Add("codigoTabelaReferencia", strconv.Itoa(codigoPeriodoReferencia))
	data.Add("codigoTipoVeiculo", strconv.Itoa(tipoVeiculo))
	data.Add("anoModelo", strconv.Itoa(anoModelo))
	data.Add("codigoTipoCombustivel", strconv.Itoa(codigoTipoCombustivel))
	data.Add("modeloCodigoExterno", codigoFipe)
	data.Add("tipoConsulta", "codigo")

	return modelo, unmarshal("http://www.fipe.org.br/IndicesConsulta-ConsultarValorComTodosParametros", data, &modelo)
}

func getAttr(t html.Token, attrName string) string {
	for _, attr := range t.Attr {
		if attr.Key == attrName {
			return attr.Val
		}
	}
	return ""
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
