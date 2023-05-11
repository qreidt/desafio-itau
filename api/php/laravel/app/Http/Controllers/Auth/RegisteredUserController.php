<?php

namespace App\Http\Controllers\Auth;

use App\Enums\UserType;
use App\Http\Controllers\Controller;
use App\Models\BankAccount;
use App\Models\User;
use App\Rules\MatchCnpjRule;
use App\Rules\MatchCpfRule;
use Illuminate\Auth\Events\Registered;
use Illuminate\Http\JsonResponse;
use Illuminate\Validation\Rules\Enum;
use Illuminate\Validation\Rules\Password;

class RegisteredUserController extends Controller
{
    /**
     * Handle an incoming registration request.
     *
     * @throws \Illuminate\Validation\ValidationException
     */
    public function store(): JsonResponse
    {

        $data = request()->validate([
            'name' => ['required', 'string', 'max:255'],
            'password' => ['required', 'confirmed', Password::min(8)->symbols()->mixedCase() ],
            'transaction_password' => ['required', 'confirmed', 'string', 'digits_between:4,10' ],
        ]);

        ['document' => $document] = request()->validate([
            'document' => [ 'required', 'regex:/(^\d{3}\.\d{3}\.\d{3}-\d{2}$)|(^\d{2}\.\d{3}\.\d{3}\/\d{4}-\d{2}$)/' ]
        ]);

        $data['document'] = $document;
        $data['password'] = bcrypt($data['password']);
        $data['transaction_password'] = bcrypt($data['transaction_password']);
        $data['type'] = UserType::tryFromLength($document);

        $user = User::create($data);

        $bank_account = BankAccount::factory()->create(['user_id' => $user->id]);

        event(new Registered($user));

        $token = $user->createToken(
            request()->header('user-agent', 'no-device')
        );

        $user->bank_accounts = collect([$bank_account]);

        return response()->json([
            'user' => $user,
            'token' => $token->plainTextToken
        ]);
    }
}
