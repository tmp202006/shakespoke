import React, { useState } from 'react';
import './App.css';
import { BrowserRouter as Switch, Route } from 'react-router-dom';
import { CachingProvider, CachingContextData } from './context/CachingContext';
import { getFavourites } from './storage/Storage';
import Search from './components/Search';
import Result from './components/Result';
import Error from './components/Error';

function App() {
  const [result, setResult] = useState<any>(undefined);
  const [error, setError] = useState<any>(undefined);
  const favs = getFavourites();
  const cache = new CachingContextData(favs);

  return (
    <div className="App">
      <CachingProvider value={cache}>
        <Switch>
          <Route exact path="/pokemon/:pokemonName">
            Pokemon
          </Route>
          <Route exact path="/favourites">
            Favourites
          </Route>
          <Route exact path="/">
            <Search
              onSuccess={(r: any) => {
                setResult(r);
                setError(undefined);
              }}
              onError={(e: any) => {
                setResult(undefined);
                setError(e);
              }}
            />
            <Result result={result} />
            <Error error={error} />
          </Route>
          <Route>Not Found</Route>
        </Switch>
      </CachingProvider>
    </div>
  );
}

export default App;
