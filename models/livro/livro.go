package livro

type Livro struct {
	nome   string
	edicao int
	autor  string
	isbn   string
}

func NovoLivro(nome string, edicao int, autor string, isbn string) *Livro {
	return &Livro{
		nome:   nome,
		edicao: edicao,
		autor:  autor,
		isbn:   isbn,
	}
}

func (l *Livro) GetNome() string {
	return l.nome
}

func (l *Livro) SetNome(nome string) {
	l.nome = nome
}

func (l *Livro) GetAutor() string {
	return l.autor
}

func (l *Livro) SetAutor(autor string) {
	l.autor = autor
}

func (l *Livro) GetEdicao() int {
	return l.edicao
}

func (l *Livro) SetEdicao(edicao int) {
	l.edicao = edicao
}

func (l *Livro) GetISBN() string {
	return l.isbn
}

func (l *Livro) SetISBN(isbn string) {
	l.isbn = isbn
}

func (l *Livro) VerificaLancamento() bool {
	return l.edicao > 2022
}
