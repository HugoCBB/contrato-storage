
import { useState } from "react";
import Card from "../../components/Card";
import "./Home.css";
import { useEffect } from "react";
import api from "../../Api";


export default function Home() {
    
    const [contratos, setContratos] = useState([])

    const getContratos = async () =>{
        try {
            const response = await api.get("/contrato", {
                withCredentials: true,
            })
            setContratos(response.data)
        
        } catch (error) {
            console.error("Error ao buscar contratos")
        }

    }
    useEffect(() => {
        getContratos();
    }, [])

    return (


    <section className="home">
        <main className="content">
        <div className="top-bar">
            <button className="add-btn">Adicionar</button>
            
            <input
            type="text"
            placeholder="Pesquisar com CPF ou nome"
            className="search-input"
            />
            
        </div>

        <div className="tables">
            {contratos.map(contrato => (
                <Card
                key={contrato.cpf}
                nome = {contrato.nome}
                cpf = {contrato.cpf}
                data = {contrato.data}

                />
                
            ))}
        </div>
        </main>
    </section>
    );
}
