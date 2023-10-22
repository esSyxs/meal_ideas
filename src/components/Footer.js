import React from 'react'
import '../styles/Footer.css'
import {Link} from "react-router-dom";

function Footer() {
  return (
    <div className="footer">
        <div className='contactUs'>
            Neatradi piemērotu recepti? 
            <Link to='/contact'> Sazināties </Link>
        </div>
        <p>&copy; 2023 Vakariņu Iedvesma</p>
    </div>
  )
}

export default Footer