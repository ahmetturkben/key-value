package domain

type FileReaderRepository interface {
	Read() []byte
}

type FileWriterRepository interface {
	Create()
	Delete()
	Write(data []byte)
}

type FileRepository interface {
	FileReaderRepository
	FileWriterRepository
}
