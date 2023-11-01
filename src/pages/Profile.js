import React, { useEffect, useState } from 'react'
import { useUser } from '../components/UserContext';
import axios from 'axios';
import RecipeItem from '../components/RecipeItem';
import {Link} from "react-router-dom";
import '../styles/Recipes.css'
import '../styles/Profile.css'
import { useNavigate } from 'react-router-dom';



function Profile() {
    const {user, logout} = useUser();
    const [userData, setUserData] = useState(null);
    const [reload, setReaload] = useState(true)

    const navigate = useNavigate();
    
    useEffect(() => {
      if(!user.isAuthenticated && reload){
        window.location.reload();
        setReaload(false)
      }
      
    }, [user.isAuthenticated, reload]);


    useEffect(() => {
        // Fetch user data only when the component mounts
    if (user.isAuthenticated) {
        const axiosInstance = axios.create({
          baseURL: 'http://localhost/api/protected',
          timeout: 5000,
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${user.token}`, // Include authorization header
          },
        });
        axiosInstance.get('/profile')
        .then((response) => {
          const userInfo = response.data;
          setUserData(userInfo);
        })
        .catch((error) => {
          // Handle authentication errors
          console.error("Couldn't get user data:", error);
        });
    }
    }, [user.isAuthenticated, user.token]);//user.token


              
      if(user.isAuthenticated && !userData) {
        return <div>Loading...</div>
      }
      else if(user.isAuthenticated){
        return (
          <div className='user-data'>
            <h1 className='user-greet'>Sveicināti, {userData.username}</h1>
            <h2 className='user-email'>E-pasts: {userData.email}</h2>
            <h2 className='user-fave-title'>Iecienītās receptes:</h2>
            {userData.foods ? (
            <div className='recipesList'>
              {userData.foods.map((recipeItem, key) => (
                  <Link to={`/recipes/${recipeItem.ID}`} className='recipeItem'>
                  <RecipeItem
                    key = {key}
                    name={recipeItem.Name}
                    produce={recipeItem.Produces}
                    appliances={recipeItem.Appliances}
                  />
                  </Link>
              ))}
            </div>):(
              <div>
                <h3 className='user-no-fave'>Neviena recepte nav atzīmēta kā iecienīta</h3>
              </div>
            )}
            
            
            <button onClick={logout}
          className='logout-button'>Izrakstīties</button>
          </div>
        );
      }


    else {
        // Render content for non-authenticated users
        return (
          <div className='not-logged-in'>
            <h1>Ieejiet savā kontā, lai skatītu šo lapu.</h1>
            {navigate('/')}
          </div>
        );
      }
}

export default Profile