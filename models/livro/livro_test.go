package livro

import "testing"

func TestVerificaLancamento(t *testing.T) {
	testCases := []struct {
		nome     string
		edicao   int
		esperado bool
	}{
		{
			nome:     "Livro 1",
			edicao:   2023,
			esperado: true,
		},
		{
			nome:     "Livro 2",
			edicao:   2024,
			esperado: true,
		},
		{
			nome:     "Livro 3",
			edicao:   2022,
			esperado: false,
		},
		{
			nome:     "Livro 4",
			edicao:   2020,
			esperado: false,
		},
		{
			nome:     "Livro 5",
			edicao:   2021,
			esperado: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.nome, func(t *testing.T) {
			livro := NovoLivro(tc.nome, tc.edicao, "Autor Teste", "ISBN-TESTE")
			resultado := livro.VerificaLancamento()

			if resultado != tc.esperado {
				t.Errorf("Falha no teste para %s. Esperado %v, mas obteve %v",
					tc.nome, tc.esperado, resultado)
			}
		})
	}
}

func TestLivroGetterSetters(t *testing.T) {
	// novo livro para testes
	livro := NovoLivro("Livro Original", 2023, "Autor Original", "ISBN-ORIGINAL")

	// test GetNome()
	if livro.GetNome() != "Livro Original" {
		t.Errorf("Falha no GetNome. Esperado 'Livro Original', mas obteve '%s'", livro.GetNome())
	}

	// test SetNome()
	livro.SetNome("Livro Atualizado")
	if livro.GetNome() != "Livro Atualizado" {
		t.Errorf("Falha no SetNome. Esperado 'Livro Atualizado', mas obteve '%s'", livro.GetNome())
	}

	// test GetEdicao()
	if livro.GetEdicao() != 2023 {
		t.Errorf("Falha no GetEdicao. Esperado 2023, mas obteve %d", livro.GetEdicao())
	}

	// test SetEdicao()
	livro.SetEdicao(2024)
	if livro.GetEdicao() != 2024 {
		t.Errorf("Falha no SetEdicao. Esperado 2024, mas obteve %d", livro.GetEdicao())
	}

	// test GetAutor()
	if livro.GetAutor() != "Autor Original" {
		t.Errorf("Falha no GetAutor. Esperado 'Autor Original', mas obteve '%s'", livro.GetAutor())
	}

	// test SetAutor()
	livro.SetAutor("Novo Autor")
	if livro.GetAutor() != "Novo Autor" {
		t.Errorf("Falha no SetAutor. Esperado 'Novo Autor', mas obteve '%s'", livro.GetAutor())
	}

	// test GetISBN()
	if livro.GetISBN() != "ISBN-ORIGINAL" {
		t.Errorf("Falha no GetISBN. Esperado 'ISBN-ORIGINAL', mas obteve '%s'", livro.GetISBN())
	}

	// test SetISBN()
	livro.SetISBN("NOVO-ISBN")
	if livro.GetISBN() != "NOVO-ISBN" {
		t.Errorf("Falha no SetISBN. Esperado 'NOVO-ISBN', mas obteve '%s'", livro.GetISBN())
	}

}
