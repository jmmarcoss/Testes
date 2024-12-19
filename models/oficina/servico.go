package oficina

type Servico struct {
	id       int
	motivo   string
	mecanico *Mecanico
	veiculo  *Veiculo
}

func NovoServico(id int, matricula int64, nome string, especialidades []string, placa, tipo string, numOcorrencias, anoFabricacao int) *Servico {
	mecanico := NovoMecanico(matricula, nome, especialidades)
	veiculo := NovoVeiculo(placa, tipo, numOcorrencias, anoFabricacao)
	return &Servico{
		id:       id,
		mecanico: mecanico,
		veiculo:  veiculo,
	}
}

func (s *Servico) GetID() int {
	return s.id
}

func (s *Servico) SetID(id int) {
	s.id = id
}

func (s *Servico) GetMotivo() string {
	return s.motivo
}

func (s *Servico) SetMotivo(motivo string) {
	s.motivo = motivo
}

func (s *Servico) GetMecanico() *Mecanico {
	return s.mecanico
}

func (s *Servico) GetVeiculo() *Veiculo {
	return s.veiculo
}

func (s *Servico) OrdemServico(motivo string) {
	s.motivo = motivo
	s.veiculo.numOcorrencias++
}
