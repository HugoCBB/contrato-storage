
import "./Card.css"

export default function Card({nome, cpf, data}) {

    return (
        <section className="contrato">
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
        </section>
    )
}
