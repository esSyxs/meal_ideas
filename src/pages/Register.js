// Register.js

import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import '../styles/RegLog.css'

const Register = () => {
  const [email, setEmail] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [passwordConfirm, setPasswordConfirm] = useState('');
  const navigate = useNavigate();

  const handleRegister = async (e) => {
    e.preventDefault();

    if (password !== passwordConfirm) {
      alert("Paroles nesakrīt, lūdzu ievadiet tās vēlreiz.");
      return;
    }

    try {
      const data = {
        username: `${username}`,
        password: `${password}`,
        email: `${email}`,
      }
      const response = await axios.post('http://localhost/api/public/signup', data);

      // Assuming your backend returns a success response upon successful registration
      if (response.status === 200) {
        alert('Reģistrācija veiksmīga! Ielogojieties, lai sāktu izmantot savu kontu!');
        navigate('/login');
      }
      else(
        alert('Reģitrācija nav izdevusies')
      )
    } catch (error) {
      // Handle registration errors (e.g., show an error message)
      console.error('Reģistrācija nav izdevusies:', error);
    }
  };

  return (

    <div className='RegLog'>
        <div className='leftSide1'>
          <h2>Reģistrēties</h2>
          <form onSubmit={handleRegister}>
          <div className="form-group">
            <input
              type="email"
              placeholder=" "
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              id='email'
            />
            <label>E-pasts</label>
          </div>

          <div className="form-group">
            <input
              type="text"
              placeholder=" "
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
              id='username'
            />
            <label>Lietotājvārds</label>
          </div>

          <div className="form-group">
            <input
              type="password"
              placeholder=" "
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              id='password'
            />
            <label>Parole</label>
          </div>

          <div className="form-group">
            <input
              type="password"
              placeholder=" "
              value={passwordConfirm}
              onChange={(e) => setPasswordConfirm(e.target.value)}
              required
              id='password-again'
            />
            <label>Parole atkārtoti</label>
          </div>

          <button type="submit">Reģistrēties</button>
          </form>
        </div>

        <div className='rightSide1'>
        <p>Jau ir konts? <br />
          <button type='button' onClick={() => navigate('/login')}>
              Pieslēgties
          </button>
        </p>
        </div>
      </div>

  );
};

export default Register;
