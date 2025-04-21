
import { useState } from "react";
import Card from "../../components/Card";
import "./Home.css";
import { useEffect } from "react";
import api from "../../Api";
import Menu from "../../components/Menu";


export default function Home() {
    
    const [contratos, setContratos] = useState([])

    const getContratos = async () =>{
        try {
            const response = await api.get("/contrato/", {
                withCredentials: true,
            })
            setContratos(response.data)
            console.log(response.data.data_criacao)
        
        } catch (error) {
            console.error("Error ao buscar contratos")
        }

    }
    useEffect(() => {
        getContratos();
    }, [])

    return (
    
    <section  >

        <section className="home">
        <Menu/>
            <main className="content">
            <div className="top-bar">
                <button className="add-btn">Adicionar</button>
                
                <input
                type="text"
                placeholder="Digite o nome ou cpf"
                className="search-input"
                />
                
            </div>

            <div className="tables">
            <h2 style={{justifyContent:"center", textAlign:"center"}}>Contratos</h2>
                
                <div className="contrato_container">
                    <div className="Contrato_card">
                    
                    {contratos && contratos.length > 0 ? (contratos.map(contrato => (
                            <Card
                            key={contrato.cpf}
                            nome = {contrato.nome}
                            cpf = {contrato.cpf}
                            data = {contrato.data_criacao}
                            />
                    ))):(
                    <p1 style={{textAlign:"center",  }} ><strong>Nenhum contrato adicionado</strong></p1>)}
                    <div/>
                </div>
            
                </div>
            </div>
            </main>
        </section>
    </section>
    );
}
