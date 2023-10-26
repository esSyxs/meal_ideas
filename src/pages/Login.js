import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import '../styles/RegLog.css'

const Login = () => {
  const [identifier, setIdentifier] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();

    try {
      const response = await axios.post('http://localhost/api/public/login', {
        identifier,
        password,
      });

      // Assuming your backend returns a token upon successful login
      const token = response.data.token;

      // Store the token in local storage or in a secure cookie for future requests
      localStorage.setItem('token', token);

      // Redirect to the profile page or some other protected route
      navigate('/profile');
    } catch (error) {
      // Handle authentication errors (e.g., show an error message)
      console.error('Authentication failed:', error);
    }
  };

  return (
    <div className='formAll'>
      <h2 className='topSide1'>Pierakstīties</h2>
      <div className='formReg'>
      <form className='leftSide1' onSubmit={handleLogin}>
        <div className='form-group'>
          <input
            type="text"
            placeholder=' '
            value={identifier}
            onChange={(e) => setIdentifier(e.target.value)}
            required
            //className={identifier ? 'active' : ''}
          />
          <label>E-pasts vai lietotājvārds</label>
          {/*<label className={identifier ? 'active' : ''}>E-pasts vai lietotājvārds</label>
        */}
        </div>
        <div className='form-group'>
          <input
            type="password"
            placeholder=' '
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            //className={password ? 'active' : ''}
          />
          <label>Parole</label>
          {/*<label className={password ? 'active' : ''}>Parole</label>
        */}</div>
        <button type="submit">Pieslēgties</button>
      </form>
      </div>
      <div className='rightSide1'>
      <p>
        Neesi lietotājs? 
        <button type='button' onClick={() => navigate('/signup')}>
            Reģistrēties
        </button>
      </p>
      </div>
      
    </div>
  );
};

export default Login;
