import React, {useState, useEffect} from 'react'
import axios from 'axios';
import '../styles/SingleRecipe.css'
import { useParams } from 'react-router-dom'; // Import useParams
import { useUser } from '../components/UserContext';
import { IconButton } from '@mui/material';
import StarIcon from '@mui/icons-material/Star';
import StarOutlineIcon from '@mui/icons-material/StarOutline';

// function SingleRecipe({parRecId}) {
  function SingleRecipe() {

    const {user} = useUser();
    const [isFavorited, setIsFavorited] = useState(false); // Track if the recipe is favorited


    const { id } = useParams();

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
      // console.log('recipedata',recipeData);
    });

    if (user.isAuthenticated) {
      const axiosInstance = axios.create({
        baseURL: 'http://localhost/api/protected',
        timeout: 5000,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${user.token}`, // Include authorization header
        },
      });
      const intId = parseInt(id, 10);
      axiosInstance.get('/profile')
      .then((response) => {
        const userInfo = response.data;
        if(userInfo.foods) {
          const favoriteRecipeIds = userInfo.foods.map(item => item.ID);

          setIsFavorited(favoriteRecipeIds.includes(intId))
        }
        else{
          setIsFavorited(false)
        }
      })
      .catch((error) => {
        // Handle authentication errors
        console.error("Couldn't get user data:", error);
      });
    }
  }, [id, user.isAuthenticated, user.token]);


  const addFave = async () => {
    if(isFavorited){
      try {
        const intId = parseInt(id, 10); // Convert the id to an integer
        if (isNaN(intId)) {
          throw new Error('Invalid ID');
        }
        const data = {
          id: intId,
        }
        const axiosInstance = axios.create({
          baseURL: 'http://localhost/api/protected',
          timeout: 5000,
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${user.token}`, // Include authorization header
          },
        });
  
        const response = await axiosInstance.put(`/favourite`, data);
        if(response.status === 200){
          const faveData = response.data;
          setIsFavorited(!isFavorited);
          return faveData;
        } else {
          throw new Error('Failed remove fave recipe data');
        }
      } catch (error) {
        console.error('Error removing fave recipe data:', error);
      }
    }
    else{
      try {
        const intId = parseInt(id, 10); // Convert the id to an integer
        if (isNaN(intId)) {
          throw new Error('Invalid ID');
        }
        const data = {
          id: intId,
        }
        const axiosInstance = axios.create({
          baseURL: 'http://localhost/api/protected',
          timeout: 5000,
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${user.token}`, // Include authorization header
          },
        });
  
        const response = await axiosInstance.post(`/favourite`, data);
        if(response.status === 200){
          const faveData = response.data;
          setIsFavorited(!isFavorited);
          return faveData;
        } else {
          throw new Error('Failed add fave recipe data');
        }
      } catch (error) {
        console.error('Error adding fave recipe data:', error);
      }
    }
    
  }

  return (
    <div className='recipes1'>
      {user.isAuthenticated ? (
            <div className='title-button'>
              <h1 className='recipeTitle1'>{recipeData.Name}</h1>
              <IconButton className='fave-button' onClick={addFave}>
                {isFavorited ? <StarIcon /> : <StarOutlineIcon />}
              </IconButton>
            </div>
            ) : (
              <h1 className='recipeTitle1'>{recipeData.Name}</h1>
            )}
        <p className='recProduce1'>Produkti: {recipeData.Produces.map(item => item?.Name).join(', ')}</p>
        <p className='recAppliances1'>Kulinārijas iekārtas: {recipeData.Appliances.map(item => item?.Name).join(', ')}</p>
        <div className='recipeText1'>
            <p>Pagatavošana:</p>
            <p>{recipeData.Desciption}</p>
        </div>
        
    </div>
  )
}

export default SingleRecipe