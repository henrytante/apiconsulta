package entities

type CPFV2 struct {
	CPF                 string `json:"cpf"`
	PAI                 string `json:"pai"`
	MAE                 string `json:"mae"`
	MunicipioNascimento string `json:"municipioNascimento"`
	Municipio           string `json:"municipio"`
	Logradouro          string `json:"logradouro"`
	Numero              string `json:"numero"`
	Bairro              string `json:"bairro"`
	CEP                 string `json:"cep"`
	RGNumero            string `json:"rgNumero"`
	RGOrgaoEmissor      string `json:"rgOrgaoEmissor"`
	RGUf                string `json:"rgUf"`
	RGDataEmissao       string `json:"rgDataEmissao"`
	CNS                 string `json:"cns"`
	Telefone            string `json:"telefone"`
	TelefoneSecundario  string `json:"telefoneSecundario"`
}

type CPFV1 struct{
	CPF string `json:"cpf"`
	NOME string `json:"nome"`
	SEXO string `json:"sexo"`
	NASCIMENTO string `json:"nascimento"`
}

type NOMEV1 struct{
	CPF string `json:"cpf"`
	NOME string `json:"nome"`
	SEXO string `json:"sexo"`
	NASCIMENTO string `json:"nascimento"`
}

type Veiculo struct {
	ID                    string `json:"id"`
	DataAtualizacao       string `json:"data_atualização"`
	Chassi                string `json:"chassi"`
	Placa                 string `json:"placa"`
	Faturado              string   `json:"faturado"`
	AnoFabricacao         string    `json:"ano_fabricacao"`
	Municipio             string `json:"municipio"`
	UFPlaca               string `json:"uf_placa"`
	MarcaModelo           string `json:"marca_modelo"`
	Combustivel           string `json:"combustivel"`
	Potencia              string `json:"potencia"`
	CapacidadeCarga       string `json:"capacidade_carga"`
	Nacionalidade         string `json:"nacionalidade"`
	Linha                 string `json:"linha"`
	Carroceria            string `json:"carroceria"`
	CaixaCambio           string `json:"caixa_cambio"`
	EixoTraseiroDif       string `json:"eixo_traseiro_dif"`
	TerceiroEixo          string `json:"terceiro_eixo"`
	Motor                 string `json:"motor"`
	TipoDocFaturado       string `json:"tipo_doc_faturado"`
	UFFaturado            string `json:"uf_faturado"`
	TipoDocProp           string `json:"tipo_doc_prop"`
	AnoModelo             string    `json:"ano_modelo"`
	TipoVeiculo           string `json:"tipo_veiculo"`
	EspecieVeiculo        string `json:"especie_veiculo"`
	TipoCarroceria        string `json:"tipo_carroceria"`
	CorVeiculo            string `json:"cor_veiculo"`
	QuantidadePassageiro  string    `json:"quantidade_passageiro"`
	SituacaoChassi        string `json:"situacao_chassi"`
	Eixos                 string    `json:"eixos"`
	TipoMontagem          string `json:"tipo_montagem"`
	TipoDocImportadora    string `json:"tipo_doc_importadora"`
	IdentImportadora      string `json:"ident_importadora"`
	DI                    string `json:"di"`
	RegistroDI            string `json:"registro_di"`
	UnidadeLocalSRF       string `json:"unidade_local_srf"`
	UltimaAtualizacao     string `json:"ultima_atualizacao"`
	Restricao1            string `json:"restricao_1"`
	Restricao2            string `json:"restricao_2"`
	Restricao3            string `json:"restricao_3"`
	Restricao4            string `json:"restricao_4"`
	LimiteRestricaoTrib   string `json:"limite_restricao_trib"`
	Cilindradas           string `json:"cilindradas"`
	CapMaximaTracao       string `json:"cap_maxima_tracao"`
	PesoBrutoTotal        string `json:"peso_bruto_total"`
	SituacaoVeiculo       string `json:"situacao_veiculo"`
	PlacaModeloAntigo     string `json:"placa_modelo_antigo"`
	PlacaModeloNovo       string `json:"placa_modelo_novo"`
	PlacaNova             string `json:"placa_nova"`
}
