import React, { useEffect, useState } from 'react'
import { useUser } from '../components/UserContext';
// import { useNavigate } from 'react-router-dom';
import axios from 'axios';


function Profile() {
    const {user, logout} = useUser();
    // const navigate = useNavigate();
    const [userData, setUserData] = useState(null);

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
    }, [user.isAuthenticated, user.token]);





    if (user.isAuthenticated) {
        if(userData){
            return (
                <div>
                  <h1>Sveicināti, {userData.username}</h1>
                  <button onClick={logout}
                className='logout-button'>Izrakstīties</button>
                </div>
              );
        }
        else {
            // Show loading indicator while fetching user data
            return <div>Loading...</div>;
          }
    }

    else {
        // Render content for non-authenticated users
        return (
          <div>
            <h1>Please log in to view your profile.</h1>
          </div>
        );
      }
}

export default Profile