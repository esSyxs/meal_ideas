import React from 'react'
import {Link} from "react-router-dom";
import BannerImage from '../assets/kitchen.jpg';
import '../styles/Home.css';

function Home() {
  return (
    <div className="home" style={{ backgroundImage: `url(${BannerImage})` }}>
        <div className="headerContainer" >
            <h1>Vienmēr kaut ko var pagatavot</h1>
            <p>Ievadi sev pieejamos produktus un atrodi sev patīkamas receptes.</p>
            <Link to="/recipes"><p>Meklēt</p></Link>
        </div>
    </div>
  )
}

export default Home