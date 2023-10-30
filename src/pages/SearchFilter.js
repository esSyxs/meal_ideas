import React, { useState } from 'react';
import axios from 'axios';
import RecipeItem from '../components/RecipeItem';
import {Link} from "react-router-dom";
import '../styles/SearchFilter.css'
import { useEffect } from 'react';

function SearchFilter() {
  const [produce, setProduce] = useState('');
  const [appliance, setAppliance] = useState('');
  const [produceOnly, setProduceOnly] = useState(false);
  const [applianceOnly, setApplianceOnly] = useState(false);
  const [resultsMessage, setResultsMessage] = useState('');
  const [recipesResults, setRecipesResults] = useState([]);
  const [produceIDs, setProduceIDs] = useState([]);
  const [appliancesIDs, setAppliancesIDs] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const [produceQuery, setProduceQuery] = useState('');
  const [applianceQuery, setApplianceQuery] = useState('')
  const [noResultsMessage, setNoResultsMessage] = useState('')


  const baseURL = 'http://localhost:80/api/public/recipes'

//update state immediately after setting it
  useEffect(() =>{
    setProduceQuery(produce)
    setApplianceQuery(appliance)
  }, [produce, appliance])



  const handleSearch = async () => {

    setRecipesResults([]);
    setNoResultsMessage('');
    setResultsMessage('');
// Split user input into an array of terms (produces and appliances)
    const produceTerms = produce.split(',').map(term => term.trim());
    const applianceTerms = appliance.split(',').map(term => term.trim());

    setProduceQuery(produce)
    setApplianceQuery(appliance)

    
    //get prouduce id
    const fetchData = async () => {

      
      // Fetch data based on the searchQuery
      if(produce.length > 0 || appliance.length > 0){
        try {
            
            const response = await axios.get('http://localhost/api/public/recipes');
            // Now you have the data in response.data, which is an object.
            // You need to convert it to an array of recipes.
            const recipesArray = Object.values(response.data);


            const matchingProduceIDs = [];
            const lowercaseProduceTerms = produceTerms.map(term => term.toLowerCase());
            const matchingApplianceIDs = [];
            const lowercaseApplianceTerms = applianceTerms.map(term => term.toLowerCase());
      
            recipesArray.forEach((recipe, key) =>{
              if(produce.length > 0){
                  const matchingProdIds = recipe.Produces
                  .filter((prod, key) => lowercaseProduceTerms.includes(prod.Name.toLowerCase()))
                  .map((prod, key) => prod.ID)
        
                  matchingProduceIDs.push(...matchingProdIds)
              }

              if (appliance.length > 0) {
                  const matchingApplIds = recipe.Appliances
                  .filter((appl, key) => lowercaseApplianceTerms.includes(appl.Name.toLowerCase()))
                  .map((appl, key) => appl.ID)
        
                  matchingApplianceIDs.push(...matchingApplIds)
              }
      
              
            })
      
            // Remove duplicate IDs (if any)
          const uniqueMatchingProduceIDs = [...new Set(matchingProduceIDs)];
          produceIDs.push(...uniqueMatchingProduceIDs)
          const uniqueMatchingApplianceIDs = [...new Set(matchingApplianceIDs)];
          appliancesIDs.push(...uniqueMatchingApplianceIDs)

          setResultsMessage(`Meklēšanas rezultāti "${produceQuery} ${applianceQuery}":`)

      
          } catch (error) {
            console.error('Produce / Appliance ID Error:', error);
            setError(error);
          } finally {
            setIsLoading(false);
          } 
      } else{
        setResultsMessage('')
      }
      
    };//fetchData

    
//=======================================================================

    // Build your request based on the user's input and checkboxes

    // Build and send a GET request for produce

    //get produce/appliance ids
    await fetchData()

    
//katram id dabū query


    let sendurl = `${baseURL}?`;
    console.log('prodlen', produce.length)
    if(produce.length > 0) {
        produceIDs.forEach(name => {
            sendurl += `produce_id=${name}&`;
          });
          if (produceOnly) {
            sendurl += 'produce_match_strict=true&';
          }
    }
    console.log('applen', appliance.length)
    if(appliance.length > 0) {
        appliancesIDs.forEach(name => {
            sendurl += `appliance_id=${name}&`;
          });
          if (applianceOnly) {
            sendurl += 'appliance_match_strict=true&';
          }
    }
    
    // Remove the trailing '&' if present
    if (sendurl.endsWith('&')) {
      sendurl = sendurl.slice(0, -1);
    }


        try {
            if(produceIDs.length > 0 || appliancesIDs.length > 0){
                const recipesResponse = await axios.get(sendurl);
                // Handle the response for produce here
                const recipesData = Object.values(recipesResponse.data);
                setRecipesResults(recipesData)

                if(recipesData.length === 0){
                  setNoResultsMessage(`Pēc meklēšanas parametriem "${produceQuery} ${applianceQuery}" nekas netika atrasts`)
                }
                

                setResultsMessage(`Meklēšanas rezultāti "${produceQuery} ${applianceQuery}":`)
                //setNoResultsMessage('')
            }
            else{
                setNoResultsMessage(`Pēc meklēšanas parametriem "${produceQuery} ${applianceQuery}" nekas netika atrasts`)
                setResultsMessage('')
            }         
            
    
        } catch (error) {
          console.error('Produce /Appliance Error:', error);
        }

        if(recipesResults.length === 0){
            setNoResultsMessage(`Pēc meklēšanas parametriem "${produceQuery} ${applianceQuery}" nekas netika atrasts`)
        }
        
    setProduce('');
    setAppliance('');
    setProduceOnly(false);
    setApplianceOnly(false);
    setProduceIDs([])
    setAppliancesIDs([])
  };//handleSearch

  const handleKeyUp = (e) => {
    if (e.key === 'Enter') {
      handleSearch();
    }
  };

  


  return (
    <div className='filterTop'>
      <h1 className='filterTitle'>Padziļinātā meklēšana</h1>
      <div className='filter-container'>
        <input
          type="text"
          placeholder="Ievadiet produktus, atdalot tos ar komatiem"
          value={produce}
          onChange={(e) => setProduce(e.target.value)}
          onKeyUp={handleKeyUp}
          className='filter-input'
        />
        <input
          type="checkbox"
          checked={produceOnly}
          onChange={() => setProduceOnly(!produceOnly)}
        />
        <label>Tikai šie produkti</label>
      </div>
      
      <div className='filter-container'>
        <input
          type="text"
          placeholder="Ievadiet kulinārijas iekārtas, atdalot tās ar komatiem"
          value={appliance}
          onChange={(e) => setAppliance(e.target.value)}
          onKeyUp={handleKeyUp}
          className='filter-input'
        />
        <input
          type="checkbox"
          checked={applianceOnly}
          onChange={() => setApplianceOnly(!applianceOnly)}
        />
        <label>Tikai šīs iekārtas</label>
      </div>

      <div>
        <button onClick={handleSearch} className='filter-search'>Meklēt</button>
      </div>

    


      {recipesResults.length > 0 ? (
        <div className='filters'>
          
          {/* <h2 className='filterTitle2'>Meklēšanas rezultāti "{produceQuery} {applianceQuery}":</h2> */}
          {resultsMessage && <h2 className='filterTitle2'>{resultsMessage}</h2>}

          {isLoading && <p className='not-success-search'>Ielādē datus...</p>}
          {error && <p className='not-success-search'>Kļūda: {error.message}</p>}

          {!isLoading && !error && (
            <div className='filterList'>
              {recipesResults.map((recipe, key) => (
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
      ):(
        <div>
          <h2 className='filterTitle2'>{noResultsMessage}</h2>

        </div>
      )}      

    </div>
  );
}

export default SearchFilter;
