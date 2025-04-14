import Contratos from "../../components/Contratos";
import "./Home.css";

export default function Home() {
    
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
            <Contratos />
        </div>
        </main>
    </section>
    );
}
