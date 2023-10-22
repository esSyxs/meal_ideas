import React from 'react'
import BannerImage from '../assets/light_blue.avif'
import '../styles/About.css'

function About() {
  return (
    <div className='about'>
        <div className='aboutTop' style={{ backgroundImage: `url(${BannerImage})` }}></div>
        <div className='aboutBottom'>
            <h1>PAR MUMS</h1>
            <p>Esam maza RTU 3.kursa studentu grupa,
                kas izveidojusi šo projektu, ar mēŗki 
                uzlabot recepšu meklēšanu latviešu valodā.
            </p>
        </div>
    </div>
  )
}

export default About