package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type Recebedor struct {
	Nome         string `json:"nome"`
	Cnpj         string `json:"cnpj"`
	Endereco     string `json:"endereco"`
	NomeFantasia string `json:"nomeFantasia"`
}

type Pagador struct {
	Nome     string `json:"nome"`
	Endereco string `json:"endereco"`
}

type Fatura struct {
	Descricao      string  `json:"descricao"`
	DiasVencimento int     `json:"diasVencimento"`
	ID             int     `json:"id"`
	Valor          float64 `json:"valor"`
}

func byteValue(caminho string) []byte {
	arquivo, err := os.Open(caminho)

	byteValue, err := ioutil.ReadAll(arquivo)
	if err != nil {
		panic(err)
	}
	return byteValue
}

func createFatura(recebedor string, pagador string, fatura string) (Recebedor, Pagador, Fatura) {
	byteValueRecebedor := byteValue(recebedor)
	byteValuePagador := byteValue(pagador)
	byteValueFatura := byteValue(fatura)

	var Rec Recebedor
	var Pag Pagador
	var Fat Fatura

	json.Unmarshal(byteValueRecebedor, &Rec)
	json.Unmarshal(byteValuePagador, &Pag)
	json.Unmarshal(byteValueFatura, &Fat)
	return Rec, Pag, Fat
}

func main() {
	Recebedor, Pagador, Fatura := createFatura(
		"./files/Recebedor/RosembackTech.json",
		"./files/Clientes/CySource.json",
		"./files/Faturas/Fatura1.json")

	emissao := time.Now().Format("02/01/2006")
	fmt.Println(Fatura)
	vencimento := time.Now().AddDate(0, 0, Fatura.DiasVencimento).Format("02/01/2006")

	pdf := gofpdf.New(gofpdf.OrientationPortrait,
		gofpdf.UnitMillimeter,
		gofpdf.PageSizeA4,
		"")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 38)
	_, lineHt := pdf.GetFontSize()
	pdf.MultiCell(0, lineHt, "Rosemback Technologies",
		gofpdf.BorderBottom, "C", false)
	pdf.MoveTo(0, lineHt*2.0)
	pdf.SetFont("Arial", "B", 24)
	_, lineHt = pdf.GetFontSize()
	pdf.MultiCell(0, lineHt*2,
		"Invoice",
		gofpdf.BorderBottom, "C", false)

	pdf.SetFont("Arial", "B", 16)
	_, lineHt = pdf.GetFontSize()
	pdf.MultiCell(0, lineHt*3, "Informacoes do Recebedor",
		gofpdf.BorderBottom, gofpdf.AlignLeft, false)
	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, lineHt*1.5,
		Recebedor.Nome,
		gofpdf.BorderNone, "L",
		false)
	pdf.MultiCell(0, lineHt*1.5,
		"CNPJ: "+Recebedor.Cnpj,
		gofpdf.BorderNone, "L",
		false)
	pdf.MultiCell(0, lineHt*1.5,
		"Endereco: "+Recebedor.Endereco,
		gofpdf.BorderNone, "L",
		false)

	pdf.SetFont("Arial", "B", 16)
	_, lineHt = pdf.GetFontSize()
	pdf.MultiCell(0, lineHt*3,
		"Cobrar a: ", gofpdf.BorderBottom, "L", false)
	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, lineHt*1.5,
		Pagador.Nome,
		gofpdf.BorderNone, "L", false)
	pdf.MultiCell(0, lineHt*1.5,
		Pagador.Endereco, gofpdf.BorderNone, "L", false)

	pdf.SetFont("Arial", "B", 16)
	pdf.MultiCell(0, lineHt*3,
		"Sobre a Fatura",
		gofpdf.BorderBottom,
		"L", false)

	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, lineHt*1.5,
		"Numero: "+strconv.Itoa(Fatura.ID),
		gofpdf.BorderNone, "L", false)
	pdf.MultiCell(0, lineHt*1.5,
		"Data: "+emissao,
		gofpdf.BorderNone, "L", false)
	pdf.MultiCell(0, lineHt*1.5,
		"Valor: "+strconv.FormatFloat(Fatura.Valor, 'f', 2, 64),
		gofpdf.BorderNone, "L", false)
	pdf.MultiCell(0, lineHt*1.5,
		"Data de Vencimento: "+vencimento,
		gofpdf.BorderNone, "L", false)
	pdf.MultiCell(0, lineHt*1.5,
		"Descricao: "+Fatura.Descricao,
		gofpdf.BorderNone, "L", false)

	err2 := pdf.OutputFileAndClose("p1.pdf")
	if err2 != nil {
		panic(err2)
	}
}
