import './App.css';
import Navbar from "./components/Navbar";
import Footer from "./components/Footer";
import Home from "./pages/Home";
import Recipes from "./pages/Recipes";
import About from "./pages/About"
import Contact from "./pages/Contact"
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom'; //Switch ir noņemts

function App() {
  return (
    <div className="App">
      <Router>
        <Navbar />
        <Routes>
          <Route path='/' element={<Home />} />
          <Route path='/recipes' element={<Recipes />} />
          <Route path='/about' element={<About />} />
          <Route path='/contact' element={<Contact />} />
        </Routes>
        <Footer />
      </Router>
      
    </div>
  );
}

export default App;



//<Route path="/" exact component={Home} />