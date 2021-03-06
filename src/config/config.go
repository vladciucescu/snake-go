package config

import (
	"fmt"
	"github.com/magiconair/properties"
)

var props *properties.Properties

func init() {
	const propertiesFile = "snake.properties"
	p, err := properties.LoadFile(propertiesFile, properties.UTF8)
	if err != nil {
		fmt.Printf("Could not read properties file: %v", err)
	}
	props = p
}

func GetBoardDimensions() (rows, columns int) {
	if props == nil {
		return defaultRows, defaultColumns
	}
	rows = props.GetInt("rows", defaultRows)
	columns = props.GetInt("columns", defaultColumns)
	if rows < 2 {
		rows = defaultRows
	}
	if columns < 2 {
		columns = defaultColumns
	}
	return rows, columns
}

func GetStartingPosition() (row, column int) {
	if props == nil {
		return defaultRows, defaultColumns
	}
	row = props.GetInt("startRow", defaultStartingRow)
	column = props.GetInt("startColumn", defaultStartingColumn)
	return row, column
}
