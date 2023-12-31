import React from 'react';
import Logo from "../assets/logo.png";
import {Link} from "react-router-dom"; //exporting a variable tāpēc {}
import '../styles/Navbar.css';
// import ReorderIcon from "@mui/icons-material/Reorder";
import SearchBar from './SearchBar';
import { useUser } from './UserContext';


function Navbar() {

    // const [openLinks, setOpenLinks] = useState(false)

    // const toggleNavbar = () => {
    //     setOpenLinks(!openLinks);
    // };

    const {user} = useUser();

    console.log('auth',user.isAuthenticated)
    console.log('token',user.token)
  return (
    <div className="navbar">
        <div className="leftSide" >
            <Link to="/"><img src={Logo} alt='Vakariņu iedvesma logo' /></Link> 
            {/* <div className='hiddenLinks'>
                <Link to="/recipes"> RECEPTES </Link>
                <Link to="/about"> PAR </Link>
                {user.isAuthenticated ? (
                <Link to="/profile"> PROFILS </Link>
            ) : (
                <Link to="/login"> PIERAKSTĪTIES </Link>
            )}
            </div> */}
        </div>
        <div className="middleSide" >
            <SearchBar />
        </div>
        <div className="rightSide">
            <Link to="/recipes"> RECEPTES </Link>
            <Link to="/about"> PAR </Link>
            {user.isAuthenticated ? (
                <Link to="/profile"> PROFILS </Link>
            ) : (
                <Link to="/login"> PIERAKSTĪTIES </Link>
            )}
            {/* <button onClick={toggleNavbar}>
                <ReorderIcon />
            </button> */}
        </div>
    </div>
  );
}

export default Navbar


// id={openLinks ? "open" : "close"}