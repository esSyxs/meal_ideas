import React, { useState } from 'react';
import axios from 'axios';
import RecipeItem from '../components/RecipeItem';
import {Link} from "react-router-dom";
import '../styles/SearchFilter.css'

function SearchFilter() {
  const [produce, setProduce] = useState('');
  const [appliance, setAppliance] = useState('');
  const [produceOnly, setProduceOnly] = useState(false);
  const [applianceOnly, setApplianceOnly] = useState(false);
  const [produceResults, setProduceResults] = useState([]);
  const [applianceResults, setApplianceResults] = useState([]);
  // const [noResultsMessage, setNoResultsMessage] = useState('');
  const [produceIDs, setProduceIDs] = useState([]);
  const [appliancesIDs, setAppliancesIDs] = useState([]);
  const [recipes, setRecipes] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const [produceQuery, setProduceQuery] = useState('');
  const [applianceQuery, setApplianceQuery] = useState('');

  

  const baseURL = 'http://localhost:80/api/public/recipes'

  const handleSearch = async () => {
    // setNoResultsMessage('');
// Split user input into an array of terms (produces and appliances)
    const produceTerms = produce.split(',').map(term => term.trim());
    const applianceTerms = appliance.split(',').map(term => term.trim());

    setProduceQuery(produce)
    setApplianceQuery(appliance)


    //get prouduce id
    const fetchDataProduce = async () => {
      
      // Fetch data based on the searchQuery
      try {
      const response = await axios.get('http://localhost/api/public/recipes');
      // Now you have the data in response.data, which is an object.
      // You need to convert it to an array of recipes.
      const recipesArray = Object.values(response.data);
      setRecipes(recipesArray);

      //
      const matchingProduceIDs = [];
      const lowercaseProduceTerms = produceTerms.map(term => term.toLowerCase());

      recipes.forEach((recipe, key) =>{
        const matchingProdIds = recipe.Produces
          .filter((prod, key) => lowercaseProduceTerms.includes(prod.Name.toLowerCase()))
          .map((prod, key) => prod.ID)

          matchingProduceIDs.push(...matchingProdIds)
      })

      // Remove duplicate IDs (if any)
    const uniqueMatchingProduceIDs = [...new Set(matchingProduceIDs)];
    produceIDs.push(...uniqueMatchingProduceIDs)

    } catch (error) {
      console.error('Produce ID Error:', error);
      setError(error);
    } finally {
      setIsLoading(false);
    } 
    };//fetchDataProduce
//=======================================================================

    // Build your request based on the user's input and checkboxes

    // Build and send a GET request for produce
  if (produce.length > 0) {
    //get produce ids
    await fetchDataProduce()

    
//katram id dabū query
    let produceURL = `${baseURL}?`;
    produceIDs.forEach(name => {
      produceURL += `produce_id=${name}&`;
    });
    if (produceOnly) {
      produceURL += 'produce_match_strict=true&';
    }
    // Remove the trailing '&' if present
    if (produceURL.endsWith('&')) {
      produceURL = produceURL.slice(0, -1);
    }

    try {
      const produceResponse = await axios.get(produceURL);
      // Handle the response for produce here
      const produceData = Object.values(produceResponse.data);
      setProduceResults(produceData); // Save the produce results in state
        

    } catch (error) {
      console.error('Produce Error:', error);
    }
  } else {
    // Clear produce results if there's no input in the appliance field
    setProduceResults([]);
  }

    //================================================
    const fetchDataAppliances = async () => {
      
      // Fetch data based on the searchQuery
      try {
      const response = await axios.get('http://localhost/api/public/recipes');
      // Now you have the data in response.data, which is an object.
      // You need to convert it to an array of recipes.
      const recipesArray = Object.values(response.data);
      setRecipes(recipesArray);

      //
      const matchingApplianceIDs = [];
      const lowercaseApplianceTerms = applianceTerms.map(term => term.toLowerCase());

      recipes.forEach((recipe, key) =>{
        const matchingApplIds = recipe.Appliances
          .filter((appl, key) => lowercaseApplianceTerms.includes(appl.Name.toLowerCase()))
          .map((appl, key) => appl.ID)

          matchingApplianceIDs.push(...matchingApplIds)
      })

      // Remove duplicate IDs (if any)
    const uniqueMatchingApplianceIDs = [...new Set(matchingApplianceIDs)];
    appliancesIDs.push(...uniqueMatchingApplianceIDs)

    } catch (error) {
      console.error('Appliance ID Error:', error);
      setError(error);
    } finally {
      setIsLoading(false);
    } 
    };
      // Build and send a GET request for appliances
  if (appliance.length > 0) {

    //get appliance ids
    await fetchDataAppliances()


    let applianceURL = `${baseURL}?`;
    appliancesIDs.forEach(name => {
      applianceURL += `appliance_id=${name}&`;
    });
    if (applianceOnly) {
      applianceURL += 'appliance_match_strict=true&';
    }
    // Remove the trailing '&' if present
    if (applianceURL.endsWith('&')) {
      applianceURL = applianceURL.slice(0, -1);
    }

    try {
      const applianceResponse = await axios.get(applianceURL);
      // Handle the response for appliances here
      const applianceData = Object.values(applianceResponse.data);
      setApplianceResults(applianceData); // Save the appliance results in state

    } catch (error) {
      console.error('Appliance Error:', error);
    }
  } else {
    // Clear appliance results if there's no input in the appliance field
    setApplianceResults([]);
  }

  // Check for both "no results" and error messages
  // if (produceResults.length === 0 && applianceResults.length === 0) {
  //   setNoResultsMessage('Pēc meklēšanas parametriem neviena recepte netika atrasta');
  // } else {
  //   setNoResultsMessage('');
  // }


    setProduce('');
    setAppliance('');
    setProduceOnly(false);
    setApplianceOnly(false);
    setProduceIDs([])
    setAppliancesIDs([])
  };//handleSearch




  return (
    <div className='filterTop'>
      <h1 className='filterTitle'>Padziļinātā meklēšana</h1>
      <div>
        <input
          type="text"
          placeholder="Ievadiet produktus, atdalot tos ar komatiem"
          value={produce}
          onChange={(e) => setProduce(e.target.value)}
        />
        <input
          type="checkbox"
          checked={produceOnly}
          onChange={() => setProduceOnly(!produceOnly)}
        />
        <label>Tikai šie produkti</label>
      </div>
      
      <div>
        <input
          type="text"
          placeholder="Ievadiet kulinārijas iekārtas, atdalot tās ar komatiem"
          value={appliance}
          onChange={(e) => setAppliance(e.target.value)}
        />
        <input
          type="checkbox"
          checked={applianceOnly}
          onChange={() => setApplianceOnly(!applianceOnly)}
        />
        <label>Tikai šīs iekārtas</label>
      </div>

      <div>
        <button onClick={handleSearch}>Meklēt</button>
      </div>

      

      {/* {noResultsMessage && <p>{noResultsMessage}</p>} */}

      {produceResults.length > 0 && (
        <div className='filters'>
          
          <h2 className='filterTitle2'>Meklēšanas rezultāti pēc produktiem "{produceQuery}":</h2>

          {isLoading && <p className='not-success-search'>Ielādē datus...</p>}
          {error && <p className='not-success-search'>Kļūda: {error.message}</p>}

          {!isLoading && !error && (
            <div className='filterList'>
              {produceResults.map((recipe, key) => (
              <Link to={`/recipes/${recipe.ID}`} className='filterItem'>
              <RecipeItem
                key={recipe.ID}
                name={recipe.Name}
                produce={recipe.Produces}
                appliances={recipe.Appliances}
              />
              </Link>
            ))}
            </div>
            
          )}
        </div>
      )}
      {produceResults.length === 0 && (
        <div>
          <h2 className='filterTitle2'>Pēc produktiem "{produceQuery}" nekas netika atrasts</h2>

        </div>
      )}


      {applianceResults.length > 0 && (
        <div className='filters'>
          <h2 className='filterTitle2'>Meklēšanas rezultāti pēc kulinārijas iekārtām "{applianceQuery}"</h2>

          {isLoading && <p className='not-success-search'>Ielādē datus...</p>}
          {error && <p className='not-success-search'>Kļūda: {error.message}</p>}

          {!isLoading && !error && (
            <div className='filterList'>
              {applianceResults.map((recipe, key) => (
                <Link to={`/recipes/${recipe.ID}`} className='filterItem'>
                <RecipeItem
                  key={recipe.ID}
                  name={recipe.Name}
                  produce={recipe.Produces}
                  appliances={recipe.Appliances}
                />
                </Link>
              ))}
            </div>
          )}
          
        </div>
      )}
      {applianceResults.length === 0 && (
        <div>
          <h2 className='filterTitle2'>Pēc kulinārijas iekārtām "{applianceQuery}" nekas netika atrasts</h2>

        </div>
      )}
    </div>
  );
}

export default SearchFilter;
