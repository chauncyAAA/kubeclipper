package printer

type ResourcePrinter interface {
	JSONPrint() ([]byte, error)
	YAMLPrint() ([]byte, error)
	TablePrint() ([]string, [][]string)
}
