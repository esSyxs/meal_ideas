import React, {useState, useEffect} from 'react'
import axios from 'axios';
import '../styles/SingleRecipe.css'
import { useParams } from 'react-router-dom'; // Import useParams

// function SingleRecipe({parRecId}) {
  function SingleRecipe() {

    const { id } = useParams();
    console.log(id);
    // Define a state variable for recipe data
  const [recipeData, setRecipeData] = useState({
    ID: 2, //default value
    Name: '',
    Produces: [],
    Appliances: [],
  });

//Function to fetch recipe data from the server

  useEffect(() => {

    const fetchRecipesData = async () => {
      try {
        const response = await axios.get(`http://localhost/api/public/recipes/${id}`
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

    // Fetch recipe data and update state
    fetchRecipesData().then((recipeData) => {
      // You can update your state with the fetched data here
      // For example, you can use the "useState" hook to manage your state.
      setRecipeData(recipeData);
      console.log(recipeData);
    });
  }, [id]);

  return (
    <div className='recipes1'>
        <h1 className='recipeTitle1'>{recipeData.Name}</h1>
        <p className='recProduce1'>Produkti: {recipeData.Produces.map(item => item.Name).join(', ')}</p>
        <p className='recAppliances1'>Kulinārijas iekārtas: {recipeData.Appliances.map(item => item.Name).join(', ')}</p>
        <div className='recipeText1'>
            <p>Pagatavošana:</p>
            <p>{recipeData.Desciption}</p>
        </div>
        
    </div>
  )
}

export default SingleRecipe