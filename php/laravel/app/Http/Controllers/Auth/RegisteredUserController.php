<?php

namespace App\Http\Controllers\Auth;

use App\Enums\UserType;
use App\Http\Controllers\Controller;
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
            'transaction_password' => ['required', 'confirmed', 'string', 'min:4', 'max:10' ],
            'type' => ['required', new Enum(UserType::class)]
        ]);

        ['document' => $document] = request()->validate([
            'document' => [
                'required', 'string',
                match ($data['type']) {
                    UserType::Fisical->value => new MatchCpfRule(),
                    UserType::Legal->value => new MatchCnpjRule(),
                }
            ]
        ]);

        $data['document'] = $document;
        $data['password'] = bcrypt($data['password']);
        $data['transaction_password'] = bcrypt($data['transaction_password']);

        $user = User::create($data);

        event(new Registered($user));

        $token = $user->createToken(
            request()->header('user-agent', 'no-device')
        );

        return response()->json([
            'user_id' => $user->id,
            'token' => $token->plainTextToken
        ]);
    }
}
