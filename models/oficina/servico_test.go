package oficina

import "testing"

func TestServico_OrdemServico(t *testing.T) {
	mecanico := NovoMecanico(12345, "João Almeida", []string{"Motor", "Suspensão"})

	veiculo := NovoVeiculo("ABC3D567", "SUV", 0, 2018)

	tests := []struct {
		nome     string
		servico  Servico
		motivo   string
		esperado int
	}{
		{
			nome: "Primeira ocorrência",
			servico: Servico{
				id:       1,
				motivo:   "Problema no óleo",
				mecanico: mecanico,
				veiculo:  veiculo,
			},
			motivo:   "Troca de óleo",
			esperado: 1,
		},
		{
			nome: "Segunda ocorrência",
			servico: Servico{
				id:       2,
				motivo:   "Problema de alinhamento",
				mecanico: mecanico,
				veiculo:  veiculo,
			},
			motivo:   "Alinhamento",
			esperado: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.nome, func(t *testing.T) {
			test.servico.OrdemServico(test.motivo)
			if test.servico.veiculo.numOcorrencias != test.esperado {
				t.Errorf("Esperado %d ocorrências, mas obteve %d", test.esperado, test.servico.veiculo.numOcorrencias)
			}
		})
	}
}
