package biblioteca

import liv "testes/models/livro"

type Biblioteca struct {
	nome        string
	cnpj        string
	anoFundacao int
	livros      []*liv.Livro
}

func NovaBiblioteca(nome string, cnpj string, anoFundacao int) *Biblioteca {
	return &Biblioteca{
		nome:        nome,
		cnpj:        cnpj,
		anoFundacao: anoFundacao,
		livros:      []*liv.Livro{},
	}
}

func (b *Biblioteca) GetNome() string {
	return b.nome
}

func (b *Biblioteca) SetNome(nome string) {
	b.nome = nome
}

func (b *Biblioteca) GetCNPJ() string {
	return b.cnpj
}

func (b *Biblioteca) SetCNPJ(cnpj string) {
	b.cnpj = cnpj
}

func (b *Biblioteca) GetAnoFundacao() int {
	return b.anoFundacao
}

func (b *Biblioteca) SetAnoFundacao(anoFundacao int) {
	b.anoFundacao = anoFundacao
}

func (b *Biblioteca) PatrimonioHistorico() bool {
	return b.anoFundacao < 1950
}

func (b *Biblioteca) ConsultarLivros() []*liv.Livro {
	return b.livros
}

func (b *Biblioteca) IncluirLivro(livro *liv.Livro) {
	b.livros = append(b.livros, livro)
}

func (b *Biblioteca) RemoverLivro(livro *liv.Livro) bool {
	for i, l := range b.livros {
		if l == livro {
			b.livros = append(b.livros[:i], b.livros[i+1:]...)
			return true
		}
	}
	return false
}

func (b *Biblioteca) AcervoPremium() bool {
	contadorLancamentos := 0
	for _, livro := range b.livros {
		if livro.VerificaLancamento() {
			contadorLancamentos++
		}
	}
	return contadorLancamentos >= 5
}

func (b *Biblioteca) QuantidadeLivros() int {
	return len(b.livros)
}
