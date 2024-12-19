package oficina

type Mecanico struct {
	matricula      int64
	nome           string
	especialidades []string
}

func NovoMecanico(matricula int64, nome string, especialidades []string) *Mecanico {
	return &Mecanico{
		nome:           nome,
		matricula:      matricula,
		especialidades: especialidades,
	}
}

func (m *Mecanico) GetMatricula() int64 {
	return m.matricula
}

func (m *Mecanico) SetMatricula(matricula int64) {
	m.matricula = matricula
}

func (m *Mecanico) GetNome() string {
	return m.nome
}

func (m *Mecanico) SetNome(nome string) {
	m.nome = nome
}

func (m *Mecanico) GetEspecialidades() []string {
	return m.especialidades
}

func (m *Mecanico) SetEspecialidades(especialidades []string) {
	m.especialidades = especialidades
}

func (m *Mecanico) FuncionarioPremium() bool {
	return len(m.especialidades) > 1
}
