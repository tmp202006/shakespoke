import { SearchItem } from '../model/Search';

const FAVS_KEY = 'favs';

export function getFavourites(): SearchItem[] | null {
  let data = localStorage.getItem(FAVS_KEY);
  if (!data) {
    return null;
  }

  let parsed = JSON.parse(data);
  return parsed.map((item: SearchItem) => {
    return new SearchItem(item.pokemon);
  });
}

export function setFavourites(items: SearchItem[]) {
  return localStorage.setItem(FAVS_KEY, JSON.stringify(items));
}
