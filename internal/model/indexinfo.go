package model

type IndexInfo struct {
	IndexName string
	StoreSize string // Assuming StoreSize is a string
	SizeKB    float64
	Status    string
	Health    string
}
