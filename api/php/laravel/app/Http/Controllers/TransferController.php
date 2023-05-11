<?php

namespace App\Http\Controllers;

use App\Enums\TransferType;
use App\Models\BankAccount;
use App\Models\Transfer;
use App\Models\User;
use App\Traits\RateLimitsRequests;
use Illuminate\Http\JsonResponse;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Facades\RateLimiter;
use Illuminate\Support\Str;
use Illuminate\Validation\Rules\Enum;
use Illuminate\Validation\ValidationException;
use Illuminate\Http\Response;

class TransferController extends Controller
{
    use RateLimitsRequests;

    /**
     * Display a listing of the resource.
     */
    public function index(): JsonResponse
    {
        $user = auth()->user();

        $transfers = Transfer::query()
            ->where('sender_user_id', '=', $user->id)
            ->orWhere('receiver_user_id', '=', $user->id)
            ->with([
                'sender_user' => fn($q) => $q->select(['id', 'document', 'name']),
                'receiver_user' => fn($q) => $q->select(['id', 'document', 'name']),
                'sender_account' => fn($q) => $q->select(['id', 'number']),
                'receiver_account' => fn($q) => $q->select(['id', 'number']),
            ])
            ->when(request('order-by'), fn($q, $value) => $q->orderBy($value))
            ->get();

        return response()->json($transfers);
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(): JsonResponse
    {
        $data = $this->validateStoreRequest();

        if (! $this->authorizeStoreRequest()) {
            abort(Response::HTTP_UNAUTHORIZED, 'Transação não autorizada');
        }

        DB::beginTransaction();

        $transfer = Transfer::create($data);
        BankAccount::query()->where('id', '=', $data['sender_account_id'])->decrement('balance', $data['value']);
        BankAccount::query()->where('id', '=', $data['receiver_account_id'])->increment('balance', $data['value']);

        DB::commit();

        $transfer->load([
            'sender_user' => fn($q) => $q->select(['id', 'document', 'name']),
            'receiver_user' => fn($q) => $q->select(['id', 'document', 'name']),
            'sender_account' => fn($q) => $q->select(['id', 'number', 'balance']),
            'receiver_account' => fn($q) => $q->select(['id', 'number', 'balance']),
        ]);

        return response()->json($transfer);
    }

    /**
     * Display the specified resource.
     */
    public function show(Transfer $transfer): JsonResponse
    {
        if (! in_array(auth()->id(), [$transfer->sender_user_id, $transfer->receiver_user_id])) {
            abort(403);
        }

        $transfer->load([
            'sender_user' => fn($q) => $q->select(['id', 'document', 'name']),
            'receiver_user' => fn($q) => $q->select(['id', 'document', 'name']),
            'sender_account' => fn($q) => $q->select(['id', 'number']),
            'receiver_account' => fn($q) => $q->select(['id', 'number']),
        ]);

        return response()->json($transfer);
    }

    public function authorizeStoreRequest(): bool
    {
        $this->ensureIsNotRateLimited();

        ['password_combinations' => $password_combinations] = request()->validate([
            'password_combinations' => ['required', 'array'],
        ]);

        $account_password = auth()->user()->transaction_password;

        foreach ($password_combinations as $cadidate_password) {
            if (Hash::check($cadidate_password, $account_password)) {
                RateLimiter::clear($this->throttleKey());
                return true;
            }
        }

        RateLimiter::hit($this->throttleKey());
        return false;
    }

    public function validateStoreRequest(): array
    {
        $data = $this->validateSenderData();
        $data = [...$data, ...$this->validateReceiverData()];
        $data = [...$data, ...$this->validateTransferType()];
        $data = [...$data, ...$this->validateTransferValue($data['type'])];

        if ($data['sender_account_id'] == $data['receiver_account_id']) {
            throw ValidationException::withMessages([
                'receiver_account_number' => 'Não é permitido realizar uma transferência para a mesma conta bancária.'
            ]);
        }

        return $data;
    }

    public function validateSenderData(): array
    {
        $user_id = auth()->id();
        $data = request()->validate([
            'sender_account_id' => [
                'required', 'integer',
                fn($id) => BankAccount::query()
                    ->where('user_id', '=', $user_id)
                    ->where('id', '=', $id)->exists(),
            ],
        ]);

        return [ 'sender_user_id' => $user_id, ...$data ];
    }

    public function validateReceiverData(): array
    {
        ['receiver_user_document' => $receiver_user_document] = request()->validate([
            'receiver_user_document' => ['required', 'string', 'exists:users,document' ],
        ]);

        ['receiver_account_number' => $receiver_account_number] = request()->validate([
            'receiver_account_number' => [
                'required', 'string',
                fn($v) => BankAccount::query()
                    ->where('user_id', '=', auth()->id())
                    ->where('number', '=', $v)->exists(),
            ]
        ]);

        $receiver_user_id = User::query()->select('id')
            ->firstWhere('document', '=', $receiver_user_document)?->id;

        $receiver_account_id = BankAccount::query()->select('id')
            ->firstWhere('number', '=', $receiver_account_number)?->id;

        return compact('receiver_user_id', 'receiver_account_id');
    }

    public function validateTransferType(): array
    {
        ['type' => $type] = request()->validate([
            'type' => ['required', new Enum(TransferType::class)]
        ]);

        return ['type' => TransferType::from($type)];
    }

    public function validateTransferValue(TransferType $transfer_type): array
    {
        return request()->validate([
            'value' => [
                'required', 'integer',
                ...(match ($transfer_type) {
                    TransferType::PIX => ['min:1', 'max:500000'],       # min 0.01     max 5000.00
                    TransferType::TED => ['min:500001', 'max:1000000'], # min 5000.01  max 10000.00
                    TransferType::DOC => ['min:1000001'],               # min 10000.01
                })
            ]
        ]);
    }

    /**
     * Get the rate limiting throttle key for the request.
     */
    public function throttleKey(): string
    {
        return Str::transliterate(Str::lower(request('document')).'|'.request()->ip());
    }
}
