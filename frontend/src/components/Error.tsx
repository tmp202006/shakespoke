import React from 'react';
import { ErrorResponse } from '../model/Response';

export class ErrorProps {
  error?: ErrorResponse;
}

export default function Error(props: ErrorProps) {
  if (!props.error || !props.error.error) {
    return <div></div>;
  }
  return (
    <div>
      <div>Error</div>
      <div>Err: {props.error.error}</div>
      <div>Code: {props.error.code}</div>
      <div>Msg: {props.error.message}</div>
    </div>
  );
}
