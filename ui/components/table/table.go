package table

import (
	"github.com/karchx/nrs/ui/constants"
)

type Model struct {
	Columns    []Column
	Rows       []Row
	EmptyState *string
  dimensions constants.Dimensions
}

type Column struct {
	Title  string
	Hidden *bool
	Width  *int
	Grow   *bool
}

type Row []string

func NewModel(dimnesions constants.Dimensions, columns []Column, rows []Row, emptyState *string) Model {
  return Model{
    Columns: columns,
    Rows: rows,
    dimensions: dimnesions,
    EmptyState: emptyState,
  }
}
