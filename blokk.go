package khain

type Blokk struct {
	Timestamp     int64
	PrevBlockHash []byte
	Hash          []byte
}
