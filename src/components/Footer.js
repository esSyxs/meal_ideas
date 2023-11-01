import React from 'react'
import { useNavigate } from 'react-router-dom';
import '../styles/Footer.css'
import { useUser } from './UserContext';


function Footer() {

  const navigate = useNavigate();
  const {user} = useUser();

  return (
    <div className="footer">
        <div className='contactUs'>
            Neatradi piemērotu recepti? 

            {user.isAuthenticated ? (
              <button type='button' className='button-contact' onClick={() => navigate('/contact')}>
                Sazināties
              </button>
            ):(
              <p className='logout-mes'>Izveido kontu un sazinies ar mums!</p>
            )}

            
        </div>
        <p>&copy; 2023 Vakariņu Iedvesma</p>
    </div>
  )
}

export default Footer