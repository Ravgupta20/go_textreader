package go_textreader

type TextReader struct {
	filepath string
}

func New(filepath string) *TextReader {
	return &TextReader{
		filepath: filepath,
	}
}
