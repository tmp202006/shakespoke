import React from 'react';
import { PokemonDescription } from '../model/Pokemon';

export class ResultProps {
  result?: PokemonDescription;
}

export default function Result(props: ResultProps) {
  if (!props.result) {
    return <div></div>;
  }

  return (
    <div>
      <div>Result</div>
      <div>Name: {props.result.name}</div>
      <div>Description: {props.result.description}</div>
    </div>
  );
}
