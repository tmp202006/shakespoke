import React from 'react';
import { SearchItem } from '../model/Search';

interface SearchResults {
  [key: string]: SearchItem;
}

export class CachingContextData {
  searchResults!: SearchResults;
  favourites!: SearchResults;
  constructor(favourites?: SearchItem[] | null) {
    this.searchResults = {};
    this.favourites = {};
    if (!favourites) {
      return;
    }

    let favMap: SearchResults = {};
    favourites?.forEach((f: SearchItem) => {
      if (f && f.pokemon && f.pokemon.name) {
        favMap[f.pokemon?.name] = f;
      }
    });
    this.favourites = favMap;
  }

  addSearch(key: string, result: SearchItem) {
    this.searchResults[key] = result;
  }

  addFavourite(key: string, result: SearchItem) {
    this.favourites[key] = result;
  }

  get(key: string): SearchItem | null {
    if (this.favourites[key]) {
      return this.favourites[key];
    }
    if (this.searchResults[key]) {
      return this.searchResults[key];
    }
    return null;
  }
}

const CachingContext = React.createContext<CachingContextData>(
  new CachingContextData()
);

export const CachingProvider = CachingContext.Provider;
export const CachingConsumer = CachingContext.Consumer;

export default CachingContext;
