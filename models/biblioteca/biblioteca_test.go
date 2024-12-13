package biblioteca

import (
	liv "leet/models/livro"
	"testing"
)

func TestPatrimonioHistorico(t *testing.T) {
	testCases := []struct {
		nome        string
		anoFundacao int
		esperado    bool
	}{
		{
			nome:        "Biblioteca 1",
			anoFundacao: 1950,
			esperado:    true,
		},
		{
			nome:        "Biblioteca 2",
			anoFundacao: 1979,
			esperado:    true,
		},
		{
			nome:        "Biblioteca 3",
			anoFundacao: 1980,
			esperado:    false,
		},
		{
			nome:        "Biblioteca 4",
			anoFundacao: 2000,
			esperado:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.nome, func(t *testing.T) {
			biblioteca := NovaBiblioteca(tc.nome, "12345678901234", tc.anoFundacao)

			resultado := biblioteca.PatrimonioHistorico()

			if resultado != tc.esperado {
				t.Errorf("Falha no teste para %s. Esperado %v, mas obteve %v",
					tc.nome, tc.esperado, resultado)
			}
		})
	}
}

func TestBibliotecaGetterSetters(t *testing.T) {
	// nova biblioteca para testes
	biblioteca := NovaBiblioteca("Biblioteca Original", "98765432109876", 1960)

	// test GetNome()
	if biblioteca.GetNome() != "Biblioteca Original" {
		t.Errorf("Falha no GetNome. Esperado 'Biblioteca Original', mas obteve '%s'", biblioteca.GetNome())
	}

	// test SetNome()
	biblioteca.SetNome("Biblioteca Atualizada")
	if biblioteca.GetNome() != "Biblioteca Atualizada" {
		t.Errorf("Falha no SetNome. Esperado 'Biblioteca Atualizada', mas obteve '%s'", biblioteca.GetNome())
	}

	// test GetCNPJ()
	if biblioteca.GetCNPJ() != "98765432109876" {
		t.Errorf("Falha no GetCNPJ. Esperado '98765432109876', mas obteve '%s'", biblioteca.GetCNPJ())
	}

	// test SetCNPJ()
	biblioteca.SetCNPJ("12345678901234")
	if biblioteca.GetCNPJ() != "12345678901234" {
		t.Errorf("Falha no SetCNPJ. Esperado '12345678901234', mas obteve '%s'", biblioteca.GetCNPJ())
	}

	// test GetAnoFundacao()
	if biblioteca.GetAnoFundacao() != 1960 {
		t.Errorf("Falha no GetAnoFundacao. Esperado 1960, mas obteve %d", biblioteca.GetAnoFundacao())
	}

	// test SetAnoFundacao()
	biblioteca.SetAnoFundacao(1970)
	if biblioteca.GetAnoFundacao() != 1970 {
		t.Errorf("Falha no SetAnoFundacao. Esperado 1970, mas obteve %d", biblioteca.GetAnoFundacao())
	}
}

func TestIncluirLivro(t *testing.T) {
	biblioteca := NovaBiblioteca("Biblioteca Teste", "12345678901234", 1990)

	livro1 := liv.NovoLivro("Livro 1", 2023, "Autor 1", "ISBN-1")
	livro2 := liv.NovoLivro("Livro 2", 2022, "Autor 2", "ISBN-2")

	biblioteca.IncluirLivro(livro1)
	biblioteca.IncluirLivro(livro2)

	if biblioteca.QuantidadeLivros() != 2 {
		t.Errorf("Falha na inclusão de livros. Esperado 2, mas obteve %d", biblioteca.QuantidadeLivros())
	}

	livros := biblioteca.ConsultarLivros()
	if livros[0] != livro1 || livros[1] != livro2 {
		t.Errorf("Livros adicionados não correspondem aos esperados")
	}
}

func TestRemoverLivro(t *testing.T) {
	biblioteca := NovaBiblioteca("Biblioteca Teste", "12345678901234", 1990)

	livro1 := liv.NovoLivro("Livro 1", 2023, "Autor 1", "ISBN-1")
	livro2 := liv.NovoLivro("Livro 2", 2022, "Autor 2", "ISBN-2")
	livro3 := liv.NovoLivro("Livro 3", 2023, "Autor 3", "ISBN-3")

	biblioteca.IncluirLivro(livro1)
	biblioteca.IncluirLivro(livro2)

	removido := biblioteca.RemoverLivro(livro1)
	if !removido || biblioteca.QuantidadeLivros() != 1 {
		t.Errorf("Falha na remoção de livro existente")
	}

	removido = biblioteca.RemoverLivro(livro3)
	if removido {
		t.Errorf("Remoção de livro inexistente deveria retornar false")
	}
}

func TestConsultarLivros(t *testing.T) {
	biblioteca := NovaBiblioteca("Biblioteca Teste", "12345678901234", 1990)

	livro1 := liv.NovoLivro("Livro 1", 2023, "Autor 1", "ISBN-1")
	livro2 := liv.NovoLivro("Livro 2", 2022, "Autor 2", "ISBN-2")

	biblioteca.IncluirLivro(livro1)
	biblioteca.IncluirLivro(livro2)

	livros := biblioteca.ConsultarLivros()

	if len(livros) != 2 {
		t.Errorf("Consulta de livros retornou quantidade incorreta. Esperado 2, obteve %d", len(livros))
	}

	if livros[0] != livro1 || livros[1] != livro2 {
		t.Errorf("Livros consultados não correspondem aos livros adicionados")
	}
}

func TestQuantidadeLivros(t *testing.T) {
	biblioteca := NovaBiblioteca("Biblioteca Teste", "12345678901234", 1990)

	if biblioteca.QuantidadeLivros() != 0 {
		t.Errorf("Biblioteca nova deveria ter 0 livros, mas tem %d", biblioteca.QuantidadeLivros())
	}

	livro1 := liv.NovoLivro("Livro 1", 2023, "Autor 1", "ISBN-1")
	livro2 := liv.NovoLivro("Livro 2", 2022, "Autor 2", "ISBN-2")

	biblioteca.IncluirLivro(livro1)
	biblioteca.IncluirLivro(livro2)

	if biblioteca.QuantidadeLivros() != 2 {
		t.Errorf("Quantidade de livros incorreta após adição. Esperado 2, obteve %d", biblioteca.QuantidadeLivros())
	}

	biblioteca.RemoverLivro(livro1)
	if biblioteca.QuantidadeLivros() != 1 {
		t.Errorf("Quantidade de livros incorreta após remoção. Esperado 1, obteve %d", biblioteca.QuantidadeLivros())
	}
}

func TestAcervoPremium(t *testing.T) {
	biblioteca := NovaBiblioteca("Biblioteca Teste", "12345678901234", 1990)

	livro1 := liv.NovoLivro("Livro 1", 2023, "Autor 1", "ISBN-1")
	livro2 := liv.NovoLivro("Livro 2", 2023, "Autor 2", "ISBN-2")
	livro3 := liv.NovoLivro("Livro 3", 2023, "Autor 3", "ISBN-3")
	livro4 := liv.NovoLivro("Livro 4", 2023, "Autor 4", "ISBN-4")
	livro5 := liv.NovoLivro("Livro 5", 2023, "Autor 5", "ISBN-5")
	livro6 := liv.NovoLivro("Livro 6", 2022, "Autor 6", "ISBN-6")

	biblioteca.IncluirLivro(livro1)
	biblioteca.IncluirLivro(livro2)
	biblioteca.IncluirLivro(livro3)
	biblioteca.IncluirLivro(livro4)
	biblioteca.IncluirLivro(livro6)

	if biblioteca.AcervoPremium() {
		t.Errorf("Falha no AcervoPremium. Esperado false, mas obteve true")
	}

	biblioteca.IncluirLivro(livro5)

	if !biblioteca.AcervoPremium() {
		t.Errorf("Falha no AcervoPremium. Esperado true, mas obteve false")
	}
}
