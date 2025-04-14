
import "./Card.css"

export default function Card({nome, cpf, data}) {

    return (
        <div className="Contrato_container">
            <h2>
                <strong>
                    Contratos
                </strong>
            </h2>

            <div className="Contrato_card">
                <div className="Contrato_info">
                    
                    <span>
                        <strong>
                            {nome}
                        </strong>
                    </span>
                    <span>{cpf}</span>
                    <span>{data}</span>
                    <a href="/">Visualizar contrato</a>
                </div>
            </div>

        
        </div>
    )
}
