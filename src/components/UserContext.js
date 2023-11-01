// UserContext.js

import { createContext, useContext, useState } from 'react';
import { useEffect } from 'react';
import {jwtDecode} from 'jwt-decode'

const UserContext = createContext();

export function useUser() {
  return useContext(UserContext);
}

export function UserProvider({ children }) {
  const [user, setUser] = useState({
    isAuthenticated: !!localStorage.getItem('token'), // Check if token exists in localStorage
    token: localStorage.getItem('token') || '',
  });

  const login = (token) => {
    localStorage.setItem('token', token);
    setUser({ isAuthenticated: true, token });
  };

  const logout = () => {
    localStorage.removeItem('token');
    setUser({ isAuthenticated: false, token: '' });
  };

  useEffect(() => {
    const checkTokenExpiration = () => {
      // Check if the token is expired
      const token = localStorage.getItem('token');
      if (token) {
        const decodedToken = jwtDecode(token); // You may need to install a JWT library
        const currentTime = Date.now() / 1000;

        if (decodedToken.exp < currentTime) {
          // Token is expired, log the user out
          logout();
        }
      }
    };

    //Check token expiration on component mount
    checkTokenExpiration();

    // Set up a timer to periodically check token expiration
    const tokenCheckInterval = setInterval(() => {
      checkTokenExpiration();
    }, 60000); // Check every minute (adjust as needed)

    return () => {
      // Clear the interval when the component unmounts
      clearInterval(tokenCheckInterval);
    };
  }, []);

  return (
    <UserContext.Provider value={{ user, login, logout }}>
      {children}
    </UserContext.Provider>
  );
}
