import React from 'react'
import Bilde from '../assets/logo.png'
import '../styles/Recipes.css'

function RecipeItem({name, produce, appliances}) {
  return (
    <div className='recipeItem1'>
      <div className='recipeBack' style={{ backgroundImage: `url(${Bilde})` }}></div>
        <div>
          <h1 className='recTitle'> {name} </h1>
          <p><span className='rectext'>Produkti:</span> {produce?.map(item => item?.Name).join(', ')}
            <br />
          <span className='rectext'>Kulinārijas iekārtas:</span> {appliances?.map(item => item?.Name).join(', ')}
          </p>
        </div>
        
        
    </div>
  )
}

export default RecipeItem
