package sheet

func ConvertToColName(colIndex int) (colName string) {
	str := []string{"Z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y"}
	for colIndex > 0 {
		m := colIndex % 26
		colName = str[m] + colName
		colIndex = colIndex / 26
		if m == 0 {
			colIndex--
		}
	}
	return
}
