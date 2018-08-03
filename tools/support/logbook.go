package support

import (
	"fmt"
	"strings"

	"github.com/sineatos/deag/utility"
)

// Logbook makes evolution record as a chronological list of dictionaries.
//
// Data can be retrieved via the select method given the appropriate names.
type Logbook interface {
	// Record saves the evolution record
	Record(info Dict)
	// Select returns a list of values associated to the names provided in argument in each dictionary of the Statistics object list.
	// One list per name is returned in order.
	Select(names []string) [][]interface{}
	// GetChapter returns the sub Logbook according to name
	GetChapter(name string) Logbook
	// Pop pops the first record
	Pop() Dict
	// Txt returns the string of records start from startindex
	Txt(startindex int) []string
	// Len returns the size of logbook
	Len() int
	// GetHeader returns the names in header
	GetHeader() []string
	// SetHeader sets the header
	SetHeader(names []string)
	// String returns the string of records
	String() string
}

// DefaultLogbook is the default implement of Logbook
type DefaultLogbook struct {
	// buffer records the logs
	buffer []Dict
	// count is the length of logs
	count int
	// chapters is a map of Logbook
	chapters map[string]*DefaultLogbook
	// header
	header []string
	// LogHeader
	LogHeader bool
}

// NewDefaultLogbook returns a default Logbook which buffer size and chapters is cSize
func NewDefaultLogbook(size, cSize int) *DefaultLogbook {
	buffer := make([]Dict, 0, size)
	chapters := make(map[string]*DefaultLogbook, cSize)
	return &DefaultLogbook{buffer: buffer, count: 0, chapters: chapters, LogHeader: true}
}

// Record saves the evolution record
func (logbook *DefaultLogbook) Record(info Dict) {
	noChapter := make(Dict)
	for key, val := range info {
		if v, ok := val.(Dict); ok {
			if chapter, ok1 := logbook.chapters[key]; !ok1 {
				chapter = NewDefaultLogbook(0, 0)
				chapter.Record(v)
				logbook.chapters[key] = chapter
			} else {
				chapter.Record(v)
			}
		} else {
			noChapter[key] = val
		}
	}
	logbook.buffer = append(logbook.buffer, noChapter)
	logbook.count++
}

// Select returns a list of values associated to the names provided in argument in buffer.
// One list per name is returned in order.
func (logbook *DefaultLogbook) Select(names []string) [][]interface{} {
	ansList := make([][]interface{}, logbook.Len())
	cols := len(names)
	for i, data := range logbook.buffer {
		ansList[i] = make([]interface{}, cols)
		for j, name := range names {
			ansList[i][j] = data[name]
		}
	}
	return ansList
}

// GetChapter returns the sub Logbook according to name
func (logbook *DefaultLogbook) GetChapter(name string) Logbook {
	return logbook.chapters[name]
}

// Pop pops the first record
func (logbook *DefaultLogbook) Pop() Dict {
	ans := make(Dict)
	ans = logbook.buffer[0]
	logbook.buffer = logbook.buffer[1:]
	for key, subLogbook := range logbook.chapters {
		subAns := subLogbook.Pop()
		ans[key] = subAns
	}
	return ans
}

// Txt returns the string of records start from startindex
func (logbook *DefaultLogbook) Txt(startindex int) []string {
	columns := logbook.GetHeader()
	chaptersTxt := make(map[string][]string)
	offsets := make(map[string]int)
	columnsLen := make([]int, len(columns))
	for i, h := range columns {
		columnsLen[i] = len(h)
	}
	for name, chapter := range logbook.chapters {
		chaptersTxt[name] = chapter.Txt(startindex)
		if startindex == 0 {
			offsets[name] = len(chaptersTxt[name]) - len(logbook.buffer)
		}
	}

	strMatrix := make([][]string, len(logbook.buffer[startindex:]))
	for i, line := range logbook.buffer[startindex:] {
		strLine := make([]string, len(columns))
		for j, name := range columns {
			var column string
			if _, ok := chaptersTxt[name]; ok {
				column = chaptersTxt[name][i+offsets[name]]
			} else {
				column = fmt.Sprintf("%v", line[name])
			}
			columnsLen[j] = utility.If(columnsLen[j] > len(column), columnsLen[j], len(column)).(int)
			strLine[j] = column
		}
		strMatrix[i] = strLine
	}

	// output header
	if startindex == 0 && logbook.LogHeader {
		nLines := 1
		if len(logbook.chapters) > 0 {
			maxLen := 0
			for _, chapter := range chaptersTxt {
				if lc := len(chapter); lc > maxLen {
					maxLen = lc
				}
			}
			nLines += maxLen - len(logbook.buffer) + 1
		}
		header := make([][]string, nLines)
		for i := range header {
			header[i] = make([]string, 0)
		}
		for j, name := range columns {
			length, spaces := 0, "     " // spaces = '\s' * 8
			if chapter, ok := chaptersTxt[name]; ok {
				for _, line := range chapter {
					if newLine := strings.Replace(line, "\t", spaces, -1); len(newLine) > length {
						length = len(newLine)
					}
				}
				blanks := nLines - 2 - offsets[name]
				for i := 0; i < blanks; i++ {
					header[i] = append(header[i], strings.Repeat(" ", length))
				}
				header[blanks] = append(header[blanks], utility.StringCentered(name, length))
				header[blanks+1] = append(header[blanks+1], strings.Repeat("-", length))
				for i := 0; i < offsets[name]; i++ {
					header[blanks+2+i] = append(header[blanks+2+i], chapter[i])
				}
			} else {
				for _, line := range strMatrix {
					if newLine := strings.Replace(line[j], "\t", spaces, -1); len(newLine) > length {
						length = len(newLine)
					}
				}
				for i, line := range header[:len(header)-1] {
					header[:len(header)-1][i] = append(line, strings.Repeat(" ", length))
				}
				header[len(header)-1] = append(header[len(header)-1], name)
			}
		}
		// fitst header next strMatrix
		newStrMatrix := make([][]string, len(strMatrix)+len(header))
		copy(newStrMatrix, header)
		copy(newStrMatrix[len(header):], strMatrix)
		strMatrix = newStrMatrix
	}

	logTemplatesArray := make([]string, len(columnsLen))
	for i, l := range columnsLen {
		logTemplatesArray[i] = fmt.Sprintf("%%%ds", l)
	}
	text := make([]string, len(strMatrix))
	for i, strLine := range strMatrix {
		tmp := make([]string, len(columnsLen))
		for j, t := range logTemplatesArray {
			tmp[j] = fmt.Sprintf(t, strLine[j])
		}
		text[i] = strings.Join(tmp, "\t")
	}
	return text
}

// Len returns the size of logbook
func (logbook *DefaultLogbook) Len() int {
	return len(logbook.buffer)
}

// GetHeader returns the names in header
func (logbook *DefaultLogbook) GetHeader() []string {
	if logbook.header == nil || len(logbook.header) == 0 {
		elem := logbook.buffer[0]
		header := make([]string, 0, len(elem)+len(logbook.chapters))
		for key := range elem {
			header = append(header, key)
		}
		for chapterName := range logbook.chapters {
			header = append(header, chapterName)
		}
		logbook.header = header
	}
	header := make([]string, len(logbook.header))
	copy(header, logbook.header)
	return header
}

// SetHeader sets the header
func (logbook *DefaultLogbook) SetHeader(names []string) {
	header := make([]string, len(names))
	copy(header, names)
	logbook.header = header
}

// String returns the string of records
func (logbook *DefaultLogbook) String() string {
	return strings.Join(logbook.Txt(0), "\n")
}
