import React from 'react'
import Bilde from '../assets/logo.png'
import '../styles/Contact.css'

function Contact() {
  return (
    <div className='contact'>
        <div className='leftSide1'
        style={{ backgroundImage: `url(${Bilde})` }}> </div>
        <div className='rightSide2'> 
            <h1>Sazinies:</h1>
            <form id='contactForm' className='contactForm' method="POST">
              <div className='contactForm'>
                <label htmlFor='email'>E-pasts</label>
                <input name='email' placeholder=' ' type='email'/>
              </div>
              <div className='contactForm'>
                <label htmlFor='message'>Ziņa</label>
                <textarea name='message' rows='6' placeholder=' ' required></textarea>
              </div>
                
                <button type='Submit'>Sūtīt</button>
            </form>
        </div>

    </div>
  )
}

export default Contact