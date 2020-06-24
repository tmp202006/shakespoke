import React, { useState, useEffect } from 'react';
import { getFavourites } from '../storage/Storage';
import { useHistory } from 'react-router-dom';

export default function FavList() {
  const [favs, setFavs] = useState<JSX.Element[]>([]);
  const history = useHistory();

  useEffect(() => {
    let favItems = getFavourites().map((f: string) => {
      return (
        <a
          className="Fav"
          key={f}
          href="# "
          onClick={() => {
            let path = `/pokemon/${f}`;
            history.push(path);
          }}
        >
          <span className="FavName">{f}</span>
        </a>
      );
    });
    setFavs(favItems);
  }, [history]);

  return (
    <React.Fragment>
      {favs && favs.length > 0 && <h3>Favourites</h3>}
      <div className="Favs">{favs}</div>
    </React.Fragment>
  );
}
