package placa

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"
	"zapys/src/db"
	"zapys/src/entities"
	"zapys/src/respostas"
)


func PLACAV1(w http.ResponseWriter, r *http.Request)  {
	placa := r.URL.Query().Get("placa")
	placa = strings.ToUpper(placa)
	if placa == ""{
		respostas.ERROR(w, http.StatusBadRequest, errors.New("Parametro placa vazio"))
		return
	}
	db, err := db.ConnectV2()
	if err != nil{
		respostas.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	var veiculo entities.Veiculo
	if err := db.QueryRow("SELECT COALESCE(id, 'não encontrado'), COALESCE(data_atualiacao, 'não encontrado'), COALESCE(chassi, 'não encontrado'), COALESCE(placa, 'não encontrado'), COALESCE(faturado, 'não encontrado'), COALESCE(ano_fabricacao, 'não encontrado'), COALESCE(municipio, 'não encontrado'), COALESCE(uf_placa, 'não encontrado'), COALESCE(marca_modelo, 'não encontrado'), COALESCE(combustivel, 'não encontrado'), COALESCE(potencia, 'não encontrado'), COALESCE(capacidade_carga, 'não encontrado'), COALESCE(nacionalidade, 'não encontrado'), COALESCE(linha, 'não encontrado'), COALESCE(carroceria, 'não encontrado'), COALESCE(caixa_cambio, 'não encontrado'), COALESCE(eixo_traseiro_dif, 'não encontrado'), COALESCE(terceiro_eixo, 'não encontrado'), COALESCE(motor, 'não encontrado'), COALESCE(tipo_doc_faturado, 'não encontrado'), COALESCE(uf_faturado, 'não encontrado'), COALESCE(tipo_doc_prop, 'não encontrado'), COALESCE(ano_modelo, 'não encontrado'), COALESCE(tipo_veiculo, 'não encontrado'), COALESCE(especie_veiculo, 'não encontrado'), COALESCE(tipo_carroceria, 'não encontrado'), COALESCE(cor_veiculo, 'não encontrado'), COALESCE(quantidade_passageiro, 'não encontrado'), COALESCE(situacao_chassi, 'não encontrado'), COALESCE(eixos, 'não encontrado'), COALESCE(tipo_montagem, 'não encontrado'), COALESCE(tipo_doc_importadora, 'não encontrado'), COALESCE(ident_importadora, 'não encontrado'), COALESCE(di, 'não encontrado'), COALESCE(registro_di, 'não encontrado'), COALESCE(unidade_local_srf, 'não encontrado'), COALESCE(ultima_atualizacao, 'não encontrado'), COALESCE(restricao_1, 'não encontrado'), COALESCE(restricao_2, 'não encontrado'), COALESCE(restricao_3, 'não encontrado'), COALESCE(restricao_4, 'não encontrado'), COALESCE(limite_restricao_trib, 'não encontrado'), COALESCE(cilindradas, 'não encontrado'), COALESCE(cap_maxima_tracao, 'não encontrado'), COALESCE(peso_bruto_total, 'não encontrado'), COALESCE(situacao_veiculo, 'não encontrado'), COALESCE(placa_modelo_antigo, 'não encontrado'), COALESCE(placa_modelo_novo, 'não encontrado'), COALESCE(placa_nova, 'não encontrado') FROM vehicles WHERE placa = ?", placa).Scan(
			&veiculo.ID, &veiculo.DataAtualizacao, &veiculo.Chassi, &veiculo.Placa,
			&veiculo.Faturado, &veiculo.AnoFabricacao, &veiculo.Municipio, &veiculo.UFPlaca,
			&veiculo.MarcaModelo, &veiculo.Combustivel, &veiculo.Potencia, &veiculo.CapacidadeCarga,
			&veiculo.Nacionalidade, &veiculo.Linha, &veiculo.Carroceria, &veiculo.CaixaCambio,
			&veiculo.EixoTraseiroDif, &veiculo.TerceiroEixo, &veiculo.Motor, &veiculo.TipoDocFaturado,
			&veiculo.UFFaturado, &veiculo.TipoDocProp, &veiculo.AnoModelo, &veiculo.TipoVeiculo,
			&veiculo.EspecieVeiculo, &veiculo.TipoCarroceria, &veiculo.CorVeiculo, &veiculo.QuantidadePassageiro,
			&veiculo.SituacaoChassi, &veiculo.Eixos, &veiculo.TipoMontagem, &veiculo.TipoDocImportadora,
			&veiculo.IdentImportadora, &veiculo.DI, &veiculo.RegistroDI, &veiculo.UnidadeLocalSRF,
			&veiculo.UltimaAtualizacao, &veiculo.Restricao1, &veiculo.Restricao2, &veiculo.Restricao3,
			&veiculo.Restricao4, &veiculo.LimiteRestricaoTrib, &veiculo.Cilindradas, &veiculo.CapMaximaTracao,
			&veiculo.PesoBrutoTotal, &veiculo.SituacaoVeiculo, &veiculo.PlacaModeloAntigo,
			&veiculo.PlacaModeloNovo, &veiculo.PlacaNova,
		); err != nil{
			if err == sql.ErrNoRows{
				respostas.ERROR(w, http.StatusNotFound, errors.New("Placa não encontrada"))
				return
			}
			respostas.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		respostas.JSON(w, http.StatusFound, veiculo)
}