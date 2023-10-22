import React from 'react'
import Bilde from '../assets/logo.png'
import '../styles/Contact.css'

function Contact() {
  return (
    <div className='contact'>
        <div className='leftSide'
        style={{ backgroundImage: `url(${Bilde})` }}> </div>
        <div className='rightSide'> 
            <h1>Sazinies:</h1>
            <form id='contactForm' method="POST">
                <label htmlFor='name'>Vārds</label>
                <input name='name' placeholder='vārds' type='text'/>
                <label htmlFor='email'>E-pasts</label>
                <input name='email' placeholder='e-pasts' type='email'/>
                <label htmlFor='message'>Ziņa</label>
                <textarea name='message' rows='6' placeholder='ziņa' required></textarea>
                <button type='Submit'>Sūtīt</button>
            </form>
        </div>

    </div>
  )
}

export default Contact