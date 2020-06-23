import React, { useState, useContext } from 'react';
import { useHistory } from 'react-router-dom';
import CachingContext from '../context/CachingContext';
import { SearchItem } from '../model/Search';
import { PokemonDescription } from '../model/Pokemon';

export class SearchProps {
  onSuccess?: Function;
  onError?: Function;
}

export default function Search(props: SearchProps) {
  const cache = useContext(CachingContext);
  const [search, setSearch] = useState<string>('');
  const history = useHistory();

  const onSubmit = (evt: any) => {
    evt.preventDefault();
    let cached = cache.get(search);
    if (cached) {
      // if (props.onSuccess) {
      //   props.onSuccess(cached);
      // }
      console.log(cached);
      return;
    }
    fetch(`http://localhost:3001/pokemon/${search}`)
      // .then((response) => {
      //   if (response.status !== 200) {
      //     // response.json().then((err) => {
      //     //   if (props.onError) {
      //     //     props.onError(err);
      //     //   }
      //     // });
      //     throw new Error('asd');
      //   }
      //   return response;
      // })
      .then((response) => {
        if (response.status !== 200) {
          return null;
        }
        // console.log(response.status);
        return response.json();
      })
      .then((data: PokemonDescription | null) => {
        cache.addSearch(search, new SearchItem(data));
        let path = `/pokemon/${search}`;
        history.push(path);
        // if (props.onSuccess) {
        //   props.onSuccess(data);
        // }
      })
      .catch((err) => {
        console.log(err);
      });
  };

  return (
    <div className="SearchBox">
      <div className="SearchBoxLabel">Shakespoke</div>
      <div>
        <form name="search" onSubmit={onSubmit}>
          <input
            className="SearchBoxInput"
            type="text"
            id="s"
            name="s"
            placeholder="Pokemon Name"
            onChange={(e) => {
              setSearch(e.target.value);
            }}
          />
        </form>
      </div>
    </div>
  );
}
