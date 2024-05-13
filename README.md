# estudosGo_dbgravacoes

Estudos Golang - package com funções para conexão e acesso ao banco gravacoes a ser utilizado nos módulos para estudos na linguagem Go

Package utilizado no módulo github.com/SergioPedron/estudosGo_webServiceGin

* conecta.go
  * Conectadbgravacoes() *sql.DB
     * Conecta ao banco mySQl gravacoes.

* albuns.go
  * AlbunsTodos(db *sql.DB) ([]Album, error)
      * Recebe o manipulador do banco e retorna um slice de struct Album com todos os álbuns da tabela

  * AlbumPorID(db *sql.DB, id int64) (Album, bool, error)
    * Recebe o manipulador do banco e um Id de álbum.  Retorna a struct do álbum encontrado, um indicador com True se o álbum foi encontrado, e um erro (se houver)

  * AdicionaAlbum(db *sql.DB, alb Album) (int64, error)
    * Recebe o manipulador do banco e uma struct do Album a ser incluído no banco.  Retorna o ID do álbum inserido e um erro (se houver)


* modelos.go
    * type Album struct
        * Define a struct Album com os nomes dos campos da tabela, bem como o nome das campos para geração do JSON
                
