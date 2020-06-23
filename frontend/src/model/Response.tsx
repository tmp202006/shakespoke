export class ErrorResponse {
  error?: string;
  code?: string;
  message?: string;
}

export class SuccessResponse {
  name!: string;
  description!: string;
}
