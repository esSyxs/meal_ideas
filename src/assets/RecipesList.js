// import axios from 'axios'
// import React from 'react'

// const baseURL = "http://localhost:80/recipes/2";

// function Recipes(){
//     const [post, setPost] = React.useState(null);

//     React.useEffect(() => {
//       axios.get(baseURL).then((response) => {
//         setPost(response.data);
//       });
//     }, []);
  
//     if (!post) return null;

//     return(
//         (post.name), (post.id), (post.produce), (post.appliances)

//     )

// }
  
// export const RecipesList = [
//     {
//         id: Recipes().post.id,
//         name: Recipes().post.name,
//         produce: Recipes().post.produce,
//         appliances: Recipes().post.appliances
//     }
// ]


export const RecipesList = [
    {
        name: "Recepte1",
        produce: ["produkts1", "produkts2"],
        //produce: "produkti",
        culinary: ["iekārta1", "iekārta2"],
        //culinary: "iekārtas",
        recipe: "receptes pagatavošanas teksts"
    },
    {
        name: "Recepte2",
        produce: ["produkts1", "produkts2"],
        //produce: "produkti",
        culinary: ["iekārta1", "iekārta2"],
        //culinary: "iekārtas",
        recipe: "receptes pagatavošanas teksts"
    },
    {
        name: "Recepte3",
        produce: ["produkts1", "produkts2"],
        //produce: "produkti",
        culinary: ["iekārta1", "iekārta2"],
        //culinary: "iekārtas",
        recipe: "receptes pagatavošanas teksts"
    }
]

//pačekot, te var mēģināt slēgties pie datu bāzes un dabūt
// tās receptes