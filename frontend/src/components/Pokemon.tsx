import React, { useState, useContext, useEffect } from 'react';
import CachingContext from '../context/CachingContext';
import { useParams } from 'react-router-dom';
import { PokemonDescription } from '../model/Pokemon';
import { SearchItem } from '../model/Search';

export default function Pokemon() {
  const { pokemonName } = useParams();
  const cache = useContext(CachingContext);
  const [pokemon, setPokemon] = useState<PokemonDescription | undefined>(
    undefined
  );

  useEffect(() => {
    let cached = cache.get(pokemonName);
    if (cached) {
      console.log(pokemonName);
      setPokemon(cached.hasResult ? cached.pokemon : undefined);
      return;
    }

    fetch(`http://localhost:3001/pokemon/${pokemonName}`)
      .then((response) => {
        if (response.status !== 200) {
          return;
        }
        return response.json();
      })
      .then((data: PokemonDescription | null) => {
        cache.addSearch(pokemonName, new SearchItem(data));
        setPokemon(data);
      })
      .catch((err) => {
        console.log(err);
      });
  });

  return (
    <React.Fragment>
      <h1 className="PokemonName">{pokemon ? pokemon.name : 'NOT FOUND'}</h1>
      {pokemon && pokemon.description && (
        <div className="PokemonDescription">{pokemon}</div>
      )}
    </React.Fragment>
  );
}
