import React from 'react'
import { useNavigate } from 'react-router-dom';
import '../styles/Footer.css'
import {Link} from "react-router-dom";

function Footer() {

  const navigate = useNavigate();

  return (
    <div className="footer">
        <div className='contactUs'>
            Neatradi piemērotu recepti? 
            {/* <Link to='/contact'> Sazināties </Link> */}
            <button type='button' className='button-contact' onClick={() => navigate('/contact')}>
              Sazināties
          </button>
        </div>
        <p>&copy; 2023 Vakariņu Iedvesma</p>
    </div>
  )
}

export default Footer