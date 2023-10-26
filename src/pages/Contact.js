import React from 'react'
import Bilde from '../assets/sazinaties.jpg'
import '../styles/Contact.css'

function Contact() {
  return (
    <div className='contact'>
        <div className='leftSide2'
        style={{ backgroundImage: `url(${Bilde})` }}>
       </div>
        <div className='rightSide2'> 
            <h1>Sazinies:</h1>
            <form id='contactForm' method="POST">
              <div className='contactForm'>
                <label>E-pasts</label>
                <input name='email' placeholder=' ' type='email'/>
              </div>
              <div className='contactForm'>
                <label >Ziņa</label>
                <textarea name='message' rows='6' placeholder=' ' required></textarea>
              </div>
              <div className='contactForm'>
                <button type='Submit'>Sūtīt</button>
              </div>
              
            </form>
        </div>

    </div>
  )
}

export default Contact