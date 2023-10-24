import React from 'react'
import BannerImage from '../assets/light_blue.avif'
import '../styles/About.css'

function About() {
  return (
    <div className='about'>
        <div className='aboutLeft' style={{ backgroundImage: `url(${BannerImage})` }}></div>
        <div className='aboutRight'>
            <h1>PAR MUMS</h1>
            <p>Esam maza RTU 3.kursa studentu grupa,
                kas izveidojusi šo projektu, ar mērķi 
                uzlabot recepšu meklēšanu latviešu valodā.
            </p>
            <p>
              Grupas dalībnieki:<br></br>
              Patrīcija Jukša<br></br>
              Anna Nikola Pavasare<br></br>
              Ēriks Šneiders<br></br>
              Eduards Otomers<br></br>
            </p>
        </div>
    </div>
  )
}

export default About