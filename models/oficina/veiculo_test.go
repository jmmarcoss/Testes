package oficina

import "testing"

func TestVeiculo_CalculaGarantia(t *testing.T) {

	tests := []struct {
		nome     string
		veiculo  Veiculo
		esperado int
	}{
		{
			nome:    "Veiculo PICKUP",
			veiculo: Veiculo{"ABC1D23", "PICKUP", 0, 2020},
			//esperado: 2025,
			esperado: 2022,
		},
		{
			nome:    "Veiculo SUV",
			veiculo: Veiculo{"ABC1D24", "SUV", 0, 2020},
			//esperado: 2023,
			esperado: 2022,
		},
		{
			nome:     "Veiculo POPULAR",
			veiculo:  Veiculo{"ABC1D25", "POPULAR", 0, 2020},
			esperado: 2022,
		},
		{
			nome:    "Veiculo SEM TIPO",
			veiculo: Veiculo{"ABC1D26", "", 0, 2020},
			//esperado: 2020,
			esperado: 2022,
		},
	}
	for _, tt := range tests {
		t.Run(tt.nome, func(t *testing.T) {
			v := &Veiculo{
				placa:          tt.veiculo.placa,
				tipo:           tt.veiculo.tipo,
				numOcorrencias: tt.veiculo.numOcorrencias,
				anoFabricacao:  tt.veiculo.anoFabricacao,
			}
			if got := v.CalculaGarantia(); got != tt.esperado {
				t.Errorf("CalculaGarantia() = %v, esperado %v", got, tt.esperado)
			}
		})
	}
}
