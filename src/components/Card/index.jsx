
import "./Card.css"

export default function Card({nome, cpf, data}) {

    return (
        <section className="contrato">

            <div className="Contrato_container">
                {/* 
                
                Card dos contratos salvos
                
                */}
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
        
        </section>
    )
}
