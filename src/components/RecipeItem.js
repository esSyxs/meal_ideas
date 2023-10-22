import React from 'react'

function RecipeItem({name, produce, appliances}) {
  return (
    <div className='recipeItem'>
        <div></div>
        <h1> {name} </h1>
        <p>Produkti: {produce.map(item => item.Name).join(', ')}</p>
        <p>Kulinārijas iekārtas: {appliances.map(item => item.Name).join(', ')}</p>
        
    </div>
  )
}

export default RecipeItem

//<p>{recipe}</p>