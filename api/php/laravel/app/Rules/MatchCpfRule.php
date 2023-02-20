<?php

namespace App\Rules;

use Closure;
use Illuminate\Contracts\Validation\DataAwareRule;
use Illuminate\Contracts\Validation\ValidationRule;

class MatchCpfRule implements ValidationRule
{

    /** Run the validation rule. */
    public function validate(string $attribute, mixed $value, Closure $fail): void
    {
        if (! $value) {
            return;
        }

        if (! preg_match('/^\d{3}\.\d{3}\.\d{3}-\d{2}$/', $value)) {
            $fail();
        }
    }
}
