package dbgravacoes

import (
	"database/sql"
	"fmt"
)

//-----------------------------------------------------------------------------------------------------------------------------------

// Recebe o manipulador do banco e retorna um slice de struct Album com todos os álbuns da tabela.  Inicia com letra maiúsula pois a função é utilizada em outros módulos
func AlbunsTodos(db *sql.DB) ([]Album, error) {
	var albuns []Album

	// Para consultar várias linhas deve-se utilizar 'Query' que retorna rows de um select que são percorridas logo abaixo.
	rows, err := db.Query("Select * From album")
	if err != nil {
		x := fmt.Errorf("erro em AlbunsTodos: %w", err)
		return nil, x
	}
	defer rows.Close() // Garante o fechamento antes do retorno da função.  Seria um equivalente ao Try/Finally

	// Percorre as rows retornadas do select
	for rows.Next() {
		var alb Album
		// scan atribui os campos do select aos ponteiros que apontam para os campos da struct Album
		if err := rows.Scan(&alb.ID, &alb.Titulo, &alb.Artista, &alb.Preco); err != nil {
			return nil, fmt.Errorf("erro em AlbunsTodos: %w", err)
		}
		albuns = append(albuns, alb) // adiciona a struct do select a slice dos albuns
	}
	// se a consulta falhar, verificar um erro aqui é a única maneira de descobrir que os resultados estão incompletos.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro na consulta AlbunsTodos: %v", err)
	}
	return albuns, nil
}

//-----------------------------------------------------------------------------------------------------------------------------------//

// Recebe o manipulador do banco e um Id de álbum. Retorna a struct do álbum encontrado, um indicador com True se o álbum foi encontrado, e um erro (se houver)
func AlbumPorID(db *sql.DB, id int64) (Album, bool, error) {
	var alb Album
	// QueryRow retorna uma única row.  Não retorna erro pois eles são postergados para o scan
	row := db.QueryRow("Select * From album Where Id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Titulo, &alb.Artista, &alb.Preco); err != nil {
		if err == sql.ErrNoRows {
			return alb, false, nil
		}
		return alb, true, fmt.Errorf("AlbumsPorId %d: %v", id, err)
	}
	return alb, true, nil
}

//-----------------------------------------------------------------------------------------------------------------------------------//

// Recebe o manipulador do banco e uma struct do Album a ser incluído no banco. Retorna o ID do álbum inserido e um erro (se houver)
func AdicionaAlbum(db *sql.DB, alb Album) (int64, error) {
	result, err := db.Exec("Insert into album (Titulo, Artista, Preco) Values (?, ?, ?)", alb.Titulo, alb.Artista, alb.Preco)
	if err != nil {
		return 0, fmt.Errorf("AdicionaAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AdicionaAlbum: %v", err)
	}
	return id, nil
}

//-----------------------------------------------------------------------------------------------------------------------------------//
