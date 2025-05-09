package vo

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

var DatamakeLimit = uint64(10000)

var currentLimit uint64

type CommonTestcaseInput struct {
	Rows []CommonTestcaseRow `json:"rows,omitempty"`
}

type CommonTestcaseRow struct {
	RowSizeId uint64                `json:"row_size_id,omitempty"`
	Values    []CommonTestcaseValue `json:"values,omitempty"`
}

type CommonTestcaseValue struct {
	ValueSizeId uint64  `json:"value_size_id,omitempty"`
	Type        string  `json:"type,omitempty"`
	Max         float64 `json:"max,omitempty"`
	Min         float64 `json:"min,omitempty"`
	MaxId       uint64  `json:"max_id,omitempty"`
	MinId       uint64  `json:"min_id,omitempty"`
}

func (c *CommonTestcaseInput) Unfold() (DataMakerInput, error) {
	currentLimit = 0
	rand.Seed(uint64(time.Now().UnixNano()))
	var input DataMakerInput
	var hsh []float64
	hsh = append(hsh, 0)
	for _, row := range c.Rows {
		if row.RowSizeId > 0 && row.RowSizeId < uint64(len(hsh)) {
			for i := 0; i < int(hsh[row.RowSizeId]); i++ {
				newRow, err := row.Unfold(&hsh)
				if err != nil {
					return input, err
				}
				input.AppendRow(newRow)
			}
		} else {
			newRow, err := row.Unfold(&hsh)
			if err != nil {
				return input, err
			}
			input.AppendRow(newRow)
		}
	}
	return input, nil
}

func (c *CommonTestcaseRow) Unfold(hsh *[]float64) (DataMakerRow, error) {
	var row DataMakerRow
	for _, v := range c.Values {
		if v.ValueSizeId > 0 && v.ValueSizeId < uint64(len(*hsh)) {
			for i := 0; i < int((*hsh)[v.ValueSizeId]); i++ {
				value, err := v.Unfold(hsh)
				if err != nil {
					return row, err
				}
				row.AppendValue(value)
			}
		} else {
			value, err := v.Unfold(hsh)
			if err != nil {
				return row, err
			}
			row.AppendValue(value)
		}
	}
	return row, nil
}

func (c *CommonTestcaseValue) Unfold(hsh *[]float64) (DataMakerValue, error) {
	if currentLimit > DatamakeLimit {
		return DataMakerValue{}, errors.New("生成次数超过" + fmt.Sprint(DatamakeLimit))
	}
	currentLimit++
	if c.MaxId > 0 && c.MaxId < uint64(len(*hsh)) {
		c.Max = (*hsh)[c.MaxId]
	}
	if c.MinId > 0 && c.MinId < uint64(len(*hsh)) {
		c.Min = (*hsh)[c.MinId]
	}
	t := GetValueType(c.Type)
	v := rand.Float64()*(c.Max-c.Min) + c.Min
	*hsh = append(*hsh, v)
	return DataMakerValue{
		Type:  t,
		Value: v,
	}, nil
}
