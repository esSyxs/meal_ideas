import React, {useState} from 'react';
import Logo from "../assets/logo.png";
import {Link} from "react-router-dom";
import '../styles/Navbar.css';
import ReorderIcon from "@mui/icons-material/Reorder";

function NavbarNoLog() {

    const [openLinks, setOpenLinks] = useState(false)

    const toggleNavbar = () => {
        setOpenLinks(!openLinks);
    };

  return (
    <div className="navbar">
        <div className="leftSide" id={openLinks ? "open" : "close"}>
            <Link to="/"><img src={Logo} alt='Vakariņu iedvesma logo' /></Link> 
            <div className='hiddenLinks'>
                <Link to="/recipes"> RECEPTES </Link>
                <Link to="/about"> PAR </Link>
                <Link to="/login"> PIERAKSTĪTIES </Link>
            </div>
        </div>
        <div className="rightSide">
            <Link to="/recipes"> RECEPTES </Link>
            <Link to="/about"> PAR </Link>
            <Link to="/login"> PIERAKSTĪTIES </Link>
            <button onClick={toggleNavbar}>
                <ReorderIcon />
            </button>
        </div>
    </div>
  );
}

export default NavbarNoLog