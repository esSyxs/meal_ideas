import React, { useState, useEffect } from 'react';
import axios from 'axios'
//import {RecipesList} from '../assets/RecipesList' //exporting a variable, tāpēc {}
import RecipeItem from '../components/RecipeItem';
import '../styles/Recipes.css'


function Recipes() {

  // Define a state variable for recipe data
  const [recipeData, setRecipeData] = useState({
    ID: 2, //default value
    Name: '',
    Produces: [],
    Appliances: [],
  });

//Function to fetch recipe data from the server
const fetchRecipesData = async () => {
  try {
    const response = await axios.get('http://localhost/recipes/2'
    ); // Modify the URL as needed
    if(response.status === 200){
      const recipeData = response.data;
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
      // You can update your state with the fetched data here
      // For example, you can use the "useState" hook to manage your state.
      setRecipeData(recipeData);
      console.log(recipeData);
    });
  }, []);


  return (
    <div className='recipes'>
        <h1 className='recipesTitle'>Receptes</h1>
        <p></p>
        <div className='recipesList'>
          {
            recipeData  && (
              <RecipeItem
                name={recipeData.Name}
                produce={recipeData.Produces.map(item => item.Name)}
                appliances={recipeData.Appliances.map(item => item.Name)}
                />
            )
          }
        </div>
    </div>
  )
}

export default Recipes


//{RecipesList.map((recipeItem, key) => {
  //return <div> {recipeItem.name} {recipeItem.produce} </div>


  /* {recipeData && recipeData.map((recipeItem, key) => {
                return (
                    <RecipeItem 
                    key={key}
                    name={recipeItem.Name}
                    produce={recipeItem.Produces}
                    appliances={recipeItem.Appliances}
                    //recipe={recipeItem.recipe}
                    />
                )
            })} */