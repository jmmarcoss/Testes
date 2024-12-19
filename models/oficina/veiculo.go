package oficina

type Veiculo struct {
	placa          string
	tipo           string
	numOcorrencias int
	anoFabricacao  int
}

func NovoVeiculo(placa, tipo string, numOcorrencias, anoFabricacao int) *Veiculo {
	return &Veiculo{
		placa:          placa,
		tipo:           tipo,
		numOcorrencias: numOcorrencias,
		anoFabricacao:  anoFabricacao,
	}
}

func (v *Veiculo) GetPlaca() string {
	return v.placa
}

func (v *Veiculo) SetPlaca(placa string) {
	v.placa = placa
}

func (v *Veiculo) GetTipo() string {
	return v.tipo
}

func (v *Veiculo) SetTipo(tipo string) {
	v.tipo = tipo
}

func (v *Veiculo) GetNumOcorrencias() int {
	return v.numOcorrencias
}

func (v *Veiculo) SetNumOcorrencias(numOcorrencias int) {
	v.numOcorrencias = numOcorrencias
}

func (v *Veiculo) GetAnoFabricacao() int {
	return v.anoFabricacao
}

func (v *Veiculo) SetAnoFabricacao(anoFabricacao int) {
	v.anoFabricacao = anoFabricacao
}

func (v *Veiculo) CalculaGarantia() int {
	switch v.tipo {
	case "PICKUP":
		return v.anoFabricacao + 2
	case "SUV":
		return v.anoFabricacao + 2
	case "POPULAR":
		return v.anoFabricacao + 2
	default:
		return v.anoFabricacao + 2
	}
}
