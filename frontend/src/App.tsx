import React from 'react';
import './App.css';
import { BrowserRouter as Switch, Route } from 'react-router-dom';
// import { getFavourites } from './storage/Storage';
import Search from './components/Search';
import Pokemon from './components/Pokemon';

function App() {
  // const favs = getFavourites();

  return (
    <div className="App">
      <Switch>
        <Route exact path="/pokemon/:pokemonName">
          <Pokemon />
        </Route>
        <Route exact path="/favourites">
          Favourites
        </Route>
        <Route exact path="/">
          <div className="SearchWrapper">
            <Search hasFavs={true} />
          </div>
        </Route>
      </Switch>
    </div>
  );
}

export default App;
