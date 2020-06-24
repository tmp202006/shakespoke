import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { PokemonDescription } from '../model/Pokemon';
import {
  getFavourites,
  addFavourites,
  removeFavourites,
} from '../storage/Storage';
import Search from './Search';

export default function Pokemon() {
  const [loading, setLoading] = useState(true);
  const { pokemonName } = useParams();
  const [isFav, setIsFav] = useState<boolean>(false);
  const [pokemon, setPokemon] = useState<PokemonDescription | undefined>(
    undefined
  );

  useEffect(() => {
    fetch(`http://localhost:3001/pokemon/${pokemonName}`)
      .then((response) => {
        if (response.status !== 200) {
          return;
        }
        return response.json();
      })
      .then((data: PokemonDescription) => {
        setPokemon(data);
      })
      .finally(() => {
        setLoading(false);
      })
      .catch((err) => {
        console.log(err);
        setLoading(false);
      });
    if (getFavourites().includes(pokemonName)) {
      setIsFav(true);
    }
  }, [pokemonName]);

  const toggleFavs = () => {
    if (!pokemon || !pokemon.name) {
      return;
    }
    if (isFav) {
      removeFavourites(pokemon?.name);
      setIsFav(false);
    } else {
      addFavourites(pokemon?.name);
      setIsFav(true);
    }
  };

  return (
    <React.Fragment>
      <div className="PokemonSearch">
        <Search hasFavs={false} />
      </div>
      {loading && <div className="Loading">Loading</div>}
      {!loading && (
        <div className="Pokemon">
          <h1 className="PokemonName">
            {pokemon ? pokemon.name : 'NOT FOUND'}
            {pokemon && (
              <span onClick={toggleFavs} className="PokemonFav">
                {isFav && <span>&#9733;</span>}
                {!isFav && <span>&#9734;</span>}
              </span>
            )}
          </h1>
          {pokemon && pokemon.description && (
            <div className="PokemonDescription">
              <span className="Quote">&quot;</span>
              <span>{pokemon.description}</span>
            </div>
          )}
        </div>
      )}
    </React.Fragment>
  );
}
