import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';
import FavList from './FavList';

export class SearchProps {
  hasFavs?: boolean;
}

export default function Search(props: SearchProps) {
  const [search, setSearch] = useState<string>('');
  const history = useHistory();

  const onSubmit = (evt: any) => {
    evt.preventDefault();
    let path = `/pokemon/${search}`;
    history.push(path);
  };

  return (
    <div>
      <div className="SearchBox">
        <div className="SearchBoxLabel">
          <a className="HomeLink" href="/">
            Shakespoke
          </a>
        </div>
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
      {props.hasFavs && <FavList />}
    </div>
  );
}
