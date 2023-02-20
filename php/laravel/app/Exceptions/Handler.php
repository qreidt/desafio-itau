<?php

namespace App\Exceptions;

use Illuminate\Foundation\Exceptions\Handler as ExceptionHandler;
use Illuminate\Http\JsonResponse;
use Illuminate\Validation\ValidationException;
use Throwable;

class Handler extends ExceptionHandler
{
    /**
     * A list of exception types with their corresponding custom log levels.
     *
     * @var array<class-string<\Throwable>, \Psr\Log\LogLevel::*>
     */
    protected $levels = [
        //
    ];

    /**
     * A list of the exception types that are not reported.
     *
     * @var array<int, class-string<\Throwable>>
     */
    protected $dontReport = [
        //
    ];

    /**
     * A list of the inputs that are never flashed to the session on validation exceptions.
     *
     * @var array<int, string>
     */
    protected $dontFlash = [
        'current_password',
        'password', 'password_confirmation',
        'transaction_password', 'transaction_password_confirmation'
    ];

    /**
     * Register the exception handling callbacks for the application.
     */
    public function register(): void
    {
        $this->reportable(function (Throwable $e) {
            //
        });
    }

    /**
     * Convert a validation exception into a JSON response.
     *
     * @param \Illuminate\Http\Client\Request $request
     * @param ValidationException $exception
     * @return JsonResponse
     */
    protected function invalidJson($request, ValidationException $exception): JsonResponse
    {
        return response()->json([
            'message' => $exception->getMessage(),
            'errors'  => $this->transformErrors($exception),
        ], $exception->status);
    }

    /** transform the error messages */
    private function transformErrors(ValidationException $exception): array
    {
        $messages = $exception->validator->errors()->getMessages();
        $failed_rules = $exception->validator->failed();

        $result = [];

        // add custom thrown messages
        foreach ($messages as $key => $item) {
            if (preg_match('/^(?:[-A-Za-z0-9]+\.)+[A-Za-z]{2,20}$/', $key)) {
                $result[$key] = $item[0];
            }
        }

        // add failed validation rules
        foreach ($failed_rules as $input => $rules) {

            foreach (array_keys($rules) as $i => $rule) {
                $rule = strtolower($rule);

                $key = "$input.$rule";

                $result[$key] = $messages[$input][$i];
            }
        }

        return $result;
    }
}
