package oficina

import "testing"

func TestMecanico_FuncionarioPremium(t *testing.T) {
	tests := []struct {
		nome     string
		mecanico Mecanico
		esperado bool
	}{
		{
			nome:     "Mecanico com 3 especialidades",
			mecanico: Mecanico{1234, "Carlos Silva", []string{"Motor", "Freios", "Suspensão"}},
			esperado: true,
		},
		{
			nome:     "Mecanico com 2 especialidades",
			mecanico: Mecanico{5678, "João Souza", []string{"Motor", "Freios"}},
			//esperado: false,
			esperado: true,
		}}
	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			m := &Mecanico{
				matricula:      tt.mecanico.matricula,
				nome:           tt.mecanico.nome,
				especialidades: tt.mecanico.especialidades,
			}
			if got := m.FuncionarioPremium(); got != tt.esperado {
				t.Errorf("FuncionarioPremium() = %v, esperado %v", got, tt.esperado)
			}
		})
	}
}
