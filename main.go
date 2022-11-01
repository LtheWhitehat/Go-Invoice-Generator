package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type Recebedor struct {
	Nome         string
	Cnpj         string
	Endereco     string
	NomeFantasia string
}

type Pagador struct {
	Nome     string
	Endereco string
}

type Fatura struct {
	Numero         int
	Data           string
	Valor          float64
	DataVencimento string
	Descricao      string
}

func main() {
	RosembackTech := Recebedor{
		Nome:         "12926599714 - Lillian da Silva Moura",
		Cnpj:         "44.983.742/0001-50",
		Endereco:     "Praia do Flamengo, 72 apto 223 - Flamengo - Rio de Janeiro - RJ",
		NomeFantasia: "Rosemback Technologies",
	}
	CySource := Pagador{
		Nome:     "CySource Academy",
		Endereco: "Carlebach St 1 , Tel Aviv , Israel",
	}
	Fatura1 := Fatura{
		Numero:         5,
		Data:           time.Now().Format("02/01/2006"),
		Valor:          600.00,
		DataVencimento: time.Now().AddDate(0, 0, 30).Format("02/01/2006"),
		Descricao:      "Servicos de treinamento de informatica",
	}
	pdf := gofpdf.New(gofpdf.OrientationPortrait,
		gofpdf.UnitMillimeter,
		gofpdf.PageSizeA4,
		"")
	w, h := pdf.GetPageSize()
	fmt.Println("largura %v, altura %v", w, h)
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
		RosembackTech.Nome,
		gofpdf.BorderNone, "L",
		false)
	pdf.MultiCell(0, lineHt*1.5,
		"CNPJ: "+RosembackTech.Cnpj,
		gofpdf.BorderNone, "L",
		false)
	pdf.MultiCell(0, lineHt*1.5,
		"Endereco: "+RosembackTech.Endereco,
		gofpdf.BorderNone, "L",
		false)
	pdf.SetFont("Arial", "B", 16)
	_, lineHt = pdf.GetFontSize()
	pdf.MultiCell(0, lineHt*3,
		"Cobrar a: ", gofpdf.BorderBottom, "L", false)
	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, lineHt*1.5,
		CySource.Nome,
		gofpdf.BorderNone, "L", false)
	pdf.MultiCell(0, lineHt*1.5,
		CySource.Endereco, gofpdf.BorderNone, "L", false)

	pdf.SetFont("Arial", "B", 16)
	pdf.MultiCell(0, lineHt*3,
		"Sobre a Fatura",
		gofpdf.BorderBottom,
		"L", false)

	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, lineHt*1.5,
		"Numero: "+strconv.Itoa(Fatura1.Numero),
		gofpdf.BorderNone, "L", false)
	pdf.MultiCell(0, lineHt*1.5,
		"Data: "+Fatura1.Data,
		gofpdf.BorderNone, "L", false)
	pdf.MultiCell(0, lineHt*1.5,
		"Valor: "+strconv.FormatFloat(Fatura1.Valor, 'f', 2, 64),
		gofpdf.BorderNone, "L", false)
	pdf.MultiCell(0, lineHt*1.5,
		"Data de Vencimento: "+Fatura1.DataVencimento,
		gofpdf.BorderNone, "L", false)
	pdf.MultiCell(0, lineHt*1.5,
		"Descricao: "+Fatura1.Descricao,
		gofpdf.BorderNone, "L", false)

	err := pdf.OutputFileAndClose("p1.pdf")
	if err != nil {
		panic(err)
	}
}
