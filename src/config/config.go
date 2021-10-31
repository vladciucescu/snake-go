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
	columns = props.GetInt("column", defaultColumns)
	return rows, columns
}

func GetStartingPosition() (row, column int) {
	if props == nil {
		return defaultRows, defaultColumns
	}
	row = props.GetInt("rows", defaultRows)
	column = props.GetInt("column", defaultColumns)
	return row, column
}
