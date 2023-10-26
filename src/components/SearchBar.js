// SearchBar.js
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { IconButton } from '@mui/material';
import SearchIcon from '@mui/icons-material/Search';
import TuneIcon from '@mui/icons-material/Tune';
import '../styles/SearchBar.css'

function SearchBar() {
  const [searchTerm, setSearchTerm] = useState('');
  const navigate = useNavigate();

  const handleSearch = () => {
    const searchTermsArray = searchTerm.split(',').map((term) => term.trim());
    // Redirect to the results page with the search term as a query parameter
    navigate(`/results?search=${searchTermsArray.join(',')}`);
    setSearchTerm('');
  };

  const clearSearch = () => {
    setSearchTerm(''); // Clear the search query when clicking outside the input field
  };

  const handleKeyUp = (e) => {
    if (e.key === 'Enter') {
      handleSearch();
    }
  };

  const handleFilter = () => {
    navigate('/results/filter');
  }

  return (
    <div className='search-container' onClick={clearSearch}>
      <IconButton className='filter-button' onClick={handleFilter}><TuneIcon /></IconButton>
      <input
        type="text"
        placeholder="Ievadi parametrus, atdalot tos ar komatiem"
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
        onClick={(e) => e.stopPropagation()}
        onKeyUp={handleKeyUp}
        className='search-input'
      />
      <IconButton className='search-button' onClick={handleSearch}><SearchIcon /> </IconButton>
    </div>
  );
}

export default SearchBar;
