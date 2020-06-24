const FAVS_KEY = 'favs';

export function getFavourites(): string[] {
  let data = localStorage.getItem(FAVS_KEY);
  if (!data) {
    return [];
  }

  let parsed = JSON.parse(data);
  return parsed;
}

export function addFavourites(item: string) {
  let favs = getFavourites();
  if (favs.includes(item)) {
    return;
  }
  return localStorage.setItem(FAVS_KEY, JSON.stringify(favs.concat(item)));
}

export function removeFavourites(item: string) {
  let favs = getFavourites();
  if (!favs.includes(item)) {
    return;
  }
  return localStorage.setItem(
    FAVS_KEY,
    JSON.stringify(favs.filter((f: string) => f !== item))
  );
}
