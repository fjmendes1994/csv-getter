package csvGetter

import (
	"strings"
)


type csvGetter struct {
	header          string
	Delimiter       string
	numberOfColumns int
	indexes         map[string]int
}

func NewCsvGetter(header string) *csvGetter {
	csvGetter := &csvGetter{
		Delimiter: ",",
		header:    header,
	}

	columns := strings.Split(csvGetter.header, csvGetter.Delimiter)

	csvGetter.numberOfColumns = len(columns)

	for i, column := range columns {
		csvGetter.indexes[column] = i
	}
	return csvGetter
}

func (c *csvGetter) Get(row string, columnName string) (string, bool) {
	values := strings.Split(row, c.Delimiter)

	if columIndex, ok := c.indexes[columnName]; ok {
		return values[columIndex], true
	} else {
		return "", false
	}

}
