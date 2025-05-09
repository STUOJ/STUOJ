package vo

import "fmt"

type ValueType uint8

const (
	ValueType_Int   ValueType = 0
	ValueType_Float ValueType = 1
	ValueType_Char  ValueType = 2
)

func (v ValueType) String() string {
	switch v {
	case ValueType_Int:
		return "int"
	case ValueType_Char:
		return "char"
	case ValueType_Float:
		return "float"
	default:
		return "unknown"
	}
}

func GetValueType(s string) ValueType {
	switch s {
	case "int":
		return ValueType_Int
	case "char":
		return ValueType_Char
	case "float":
		return ValueType_Float
	default:
		return ValueType(0)
	}
}

type DataMakerInput struct {
	Rows []DataMakerRow
}

type DataMakerRow struct {
	Values []DataMakerValue
}

type DataMakerValue struct {
	Type  ValueType
	Value float64
}

func (c *DataMakerInput) AppendRow(row DataMakerRow) {
	c.Rows = append(c.Rows, row)
}
func (c *DataMakerInput) Size() uint64 {
	return uint64(len(c.Rows))
}

func (c *DataMakerInput) GetRow(index uint64) DataMakerRow {
	return c.Rows[index]
}

func (c *DataMakerRow) AppendValue(value DataMakerValue) {
	c.Values = append(c.Values, value)
}

func (c *DataMakerRow) Size() uint64 {
	return uint64(len(c.Values))
}

func (c *DataMakerRow) GetValue(index uint64) DataMakerValue {
	return c.Values[index]
}

func (c *DataMakerInput) String() string {
	var s string
	for _, row := range c.Rows {
		for _, value := range row.Values {
			if value.Type == ValueType_Char {
				s += fmt.Sprintf("%c ", byte(value.Value))
			} else if value.Type == ValueType_Int {
				s += fmt.Sprintf("%v ", int(value.Value))
			} else if value.Type == ValueType_Float {
				s += fmt.Sprintf("%v ", value.Value)
			}
		}
		s += "\n"
	}
	return s
}
