// SearchResults.js
import React, { useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom';
import axios from 'axios';
import RecipeItem from '../components/RecipeItem';
import '../styles/SearchResults.css'
import {Link} from "react-router-dom";

function SearchResults() {
  const location = useLocation();
  const searchQuery = new URLSearchParams(location.search).get('search');
  const searchTerms = searchQuery ? searchQuery.split(',') : [];
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const [recipes, setRecipes] = useState([]);

  //for pagination
  const [currentPage, setCurrentPage] = useState(1);
  const recipesPerPage = 3;


  const fetchData = async () => {
    setIsLoading(true); // Set loading state to true
    // Fetch data based on the searchQuery
    try {
    const response = await axios.get('http://localhost/api/public/recipes');
    // Now you have the data in response.data, which is an object.
    // You need to convert it to an array of recipes.
    const recipesArray = Object.values(response.data);
    setRecipes(recipesArray);
    console.log('recipesArray: ', recipesArray)
  } catch (error) {
    setError(error);
  } finally {
    setIsLoading(false);
  }
  };

  console.log('searchQuery', searchTerms)

  const searchRecipes = (searchTerms) => {
    return recipes.filter((recipe) => {
      return searchTerms.every((term) =>
        recipe.Name.toLowerCase().includes(term.toLowerCase()) ||
        recipe.Produces.some((produce) =>
          produce.Name.toLowerCase().includes(term.toLowerCase())
        ) ||
        recipe.Appliances.some((appliances) =>
        appliances.Name.toLowerCase().includes(term.toLowerCase())
        )
      );
    });
  };

  const filteredRecipes = searchRecipes(searchTerms);



  useEffect(() => {
    if (searchQuery) {
      fetchData().then(() => {
        console.log('Recipes data:', recipes);
      });
    }
  }, [searchQuery]);

// Handle page change
const handlePageChange = (pageNumber) => {
  setCurrentPage(pageNumber);
};


const getCurrentPageRecipes = () => {
  const indexOfLastRecipe = currentPage * recipesPerPage;
  const indexOfFirstRecipe = indexOfLastRecipe - recipesPerPage;
  return filteredRecipes.slice(indexOfFirstRecipe, indexOfLastRecipe);
};



  return (
    <div className='searches'>
      <h1 className='searchTitle'>Meklēšanas rezultāti "{searchQuery}"</h1>

      {isLoading && <p className='not-success-search'>Ielādē datus...</p>}

      {error && <p className='not-success-search'>Kļūda: {error.message}</p>}

      {!isLoading && !error && (
        <div className='searchList'>
          {filteredRecipes.length === 0 ? (
            <p className='not-success-search'>Neviena recepte netika atrasta.</p>
          ) : (
              getCurrentPageRecipes().map((recipe, key) => (
                <Link to={`/recipes/${recipe.ID}`} className='searchItem'>
                <RecipeItem
                  key={recipe.ID}
                  name={recipe.Name}
                  produce={recipe.Produces}
                  appliances={recipe.Appliances}
                />
                </Link>
            ))
          )}
        </div>
      )}
      {/* Pagination */}
      <div className='pagination'>
          {Array.from({length: Math.ceil(filteredRecipes.length / recipesPerPage)}).map((_, index) => (
            <button className={`pagination-button ${currentPage === index + 1 ? 'active' : ''}`} key={index} onClick={() => handlePageChange(index + 1)}>
              {index + 1}
            </button>
          ))}
      </div>
    </div>
  );
}

export default SearchResults;
