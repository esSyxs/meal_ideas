import './App.css';
import Navbar from "./components/Navbar";
import Footer from "./components/Footer";
import Home from "./pages/Home";
import Recipes from "./pages/Recipes";
import About from "./pages/About"
import Contact from "./pages/Contact"
import SingleRecipe from './pages/SingleRecipe';
import SearchResults from './pages/SearchResults';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom'; //Switch ir noņemts
import Login from './pages/Login';
import Register from './pages/Register';

function App() {
  return (
    <div className="App">
      <Router>
        <Navbar />
        <Routes>
          <Route path='/' element={<Home />} />
          <Route path='/recipes' element={<Recipes />} />
          <Route path='/about' element={<About />} />
          <Route path='/login' element={<Login />}/>
          <Route path='/signup' element={<Register />}/>
          <Route path='/contact' element={<Contact />} />
          <Route path='/recipes/:id' element={<SingleRecipe />} />
          <Route path='/results' element={<SearchResults />} />
        </Routes>
        <Footer />
      </Router>
      
    </div>
  );
}

export default App;
