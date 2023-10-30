// UserContext.js

import { createContext, useContext, useState } from 'react';

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

  return (
    <UserContext.Provider value={{ user, login, logout }}>
      {children}
    </UserContext.Provider>
  );
}
