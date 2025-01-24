package main

import "fmt"

type DataMiner interface {
	ExtractData(file string) string
	ParseData(rawData string) string
	AnalyzeData(parsedData string) string
	GenerateReport(analyzedData string) string
}

type BaseDataMiner struct{}

func (b *BaseDataMiner) Process(file string, miner DataMiner) string {
	rawData := miner.ExtractData(file)
	parsedData := miner.ParseData(rawData)
	analyzedData := miner.AnalyzeData(parsedData)
	report := miner.GenerateReport(analyzedData)
	return report
}

type PDFMiner struct{}

func (p *PDFMiner) ExtractData(file string) string {
	return fmt.Sprintf("PDF data extracted from %s", file)
}

func (p *PDFMiner) ParseData(rawData string) string {
	return fmt.Sprintf("Parsed PDF data: %s", rawData)
}

func (p *PDFMiner) AnalyzeData(parsedData string) string {
	return fmt.Sprintf("Analyzed PDF data: %s", parsedData)
}

func (p *PDFMiner) GenerateReport(analyzedData string) string {
	return fmt.Sprintf("PDF Report: %s", analyzedData)
}

type DOCMiner struct{}

func (d *DOCMiner) ExtractData(file string) string {
	return fmt.Sprintf("DOC data extracted from %s", file)
}

func (d *DOCMiner) ParseData(rawData string) string {
	return fmt.Sprintf("Parsed DOC data: %s", rawData)
}

func (d *DOCMiner) AnalyzeData(parsedData string) string {
	return fmt.Sprintf("Analyzed DOC data: %s", parsedData)
}

func (d *DOCMiner) GenerateReport(analyzedData string) string {
	return fmt.Sprintf("DOC Report: %s", analyzedData)
}

type CSVMiner struct{}

func (c *CSVMiner) ExtractData(file string) string {
	return fmt.Sprintf("CSV data extracted from %s", file)
}

func (c *CSVMiner) ParseData(rawData string) string {
	return fmt.Sprintf("Parsed CSV data: %s", rawData)
}

func (c *CSVMiner) AnalyzeData(parsedData string) string {
	return fmt.Sprintf("Analyzed CSV data: %s", parsedData)
}

func (c *CSVMiner) GenerateReport(analyzedData string) string {
	return fmt.Sprintf("CSV Report: %s", analyzedData)
}

func main() {
	file := "sample_file"
	baseMiner := &BaseDataMiner{}

	pdfMiner := &PDFMiner{}
	fmt.Println(baseMiner.Process(file, pdfMiner))

	docMiner := &DOCMiner{}
	fmt.Println(baseMiner.Process(file, docMiner))

	csvMiner := &CSVMiner{}
	fmt.Println(baseMiner.Process(file, csvMiner))
}
