import "./menu.css"
import { useState } from "react"

export default function Menu() {

    const [isOpen, setIsopen] = useState(false)

    const menuOpen = () => {setIsopen(!isOpen)}

    return(
        <div>
            <div className={`menu ${isOpen ? "open" : ""}`}>
                <button className="close-btn" onClick={menuOpen}>×</button>
                <ul>
                    <li><a href="#">Perfil</a></li>
                    <li><a href="#">Sobre</a></li>
                    <li><a href="#">Serviços</a></li>
                    <li><a href="#">Contato</a></li>
                </ul>
            </div>

            {!isOpen && (
            <button className="menu-btn" onClick={menuOpen}>
                ☰
            </button>
            )}  
        </div>
    )
}