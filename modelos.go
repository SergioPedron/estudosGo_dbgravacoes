package dbgravacoes

// Define a struct Album com os nomes dos campos da tabela, bem como o nome das campos para geração do JSON
type Album struct {
	ID      int64   `json:"id"` // ao serializar JSON, definie o nome do campo.   "id"
	Titulo  string  `json:"titulo"`
	Artista string  `json:"artista"`
	Preco   float32 `json:"preco"`
}
