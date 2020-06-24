import { PokemonDescription } from './Pokemon';

export class SearchItem {
  hasResult!: boolean;
  pokemon?: PokemonDescription;
  constructor(pokemon?: PokemonDescription | null | undefined) {
    this.hasResult = Boolean(pokemon);
    this.pokemon = pokemon || undefined;
  }
}
