import React, { useState, useEffect } from 'react';
import axios from 'axios'
import RecipeItem from '../components/RecipeItem';
import '../styles/Recipes.css'
import {Link} from "react-router-dom";


function Recipes() {

  const [recipeData, setRecipeData] = useState([]);

//for pagination
  const [currentPage, setCurrentPage] = useState(1);
  const recipesPerPage = 6;

//Function to fetch recipe data from the server
const fetchRecipesData = async () => {
  try {
    const response = await axios.get('http://localhost/api/public/recipes'
    ); // Modify the URL as needed
    if(response.status === 200){
      const recipeData = Object.values(response.data);
      setRecipeData(recipeData);//
      return recipeData;
    } else {
      throw new Error('Failed to fetch recipe data');
    }
  } catch (error) {
    console.error('Error fetching recipe data:', error);
  }
};
  useEffect(() => {
    // Fetch recipe data and update state
    fetchRecipesData().then((recipeData) => {
      setRecipeData(recipeData);
      
      console.log('Fetched recipe data:', recipeData);
    });
  }, []);

  console.log('Current state of recipeData:', recipeData);

   // Handle page change
   const handlePageChange = (pageNumber) => {
    setCurrentPage(pageNumber);
  };

  
  const getCurrentPageRecipes = () => {
    const indexOfLastRecipe = currentPage * recipesPerPage;
    const indexOfFirstRecipe = indexOfLastRecipe - recipesPerPage;
    return recipeData.slice(indexOfFirstRecipe, indexOfLastRecipe);
  };
 



  return (
    <div className='recipes'>
        <h1 className='recipesTitle'>Receptes</h1>
        <p></p>
        
          <div className='recipesList'>
        
            {getCurrentPageRecipes().map((recipeItem, key) => (
              <Link to={`/recipes/${recipeItem.ID}`} className='recipeItem'>
              <RecipeItem
                key = {key}
                name={recipeItem.Name}
                produce={recipeItem.Produces}
                appliances={recipeItem.Appliances}
              />
              </Link>
            )
            )}
        </div>

        {/* Pagination */}
        <div className='pagination'>
          {Array.from({length: Math.ceil(recipeData.length / recipesPerPage)}).map((_, index) => (
            <button className={`pagination-button ${currentPage === index + 1 ? 'active' : ''}`} key={index} onClick={() => handlePageChange(index + 1)}>
              {index + 1}
            </button>
          ))}
        </div>
    </div>
  )
}

export default Recipes