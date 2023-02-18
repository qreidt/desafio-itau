<?php

namespace App\Http\Controllers;

use App\Enums\TransferType;
use App\Models\BankAccount;
use App\Models\Transfer;
use App\Models\User;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Http\Response;
use Illuminate\Support\Facades\DB;
use Illuminate\Validation\Rules\Enum;

class TransferController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index(): JsonResponse
    {
        $user = auth()->user();

        $transfers = Transfer::query()
            ->where('sender_user_id', '=', $user->id)
            ->orWhere('receiver_user_id', '=', $user->id)
            ->get();

        return response()->json($transfers);
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(): JsonResponse
    {
        $data = $this->validateSenderData();
        $data = [...$data, ...$this->validateReceiverData()];
        $data = [...$data, ...$this->validateTransferType()];
        $data = [...$data, ...$this->validateTransferValue($data['type'])];

        DB::beginTransaction();

        $transfer = Transfer::create($data);
        BankAccount::query()->find($data['sender_account_id'])
            ?->increment('balance', $data['value']);
        BankAccount::query()->find($data['receiver_account_id'])
            ?->increment('balance', $data['value']);

        DB::commit();

        return response()->json($transfer);
    }

    public function validateSenderData(): array
    {
        $user_id = auth()->id();
        $data = request()->validate([
            'sender_account_id' => [
                'required', 'integer',
                fn($id) => BankAccount::query()
                    ->where('user_id', '=', $user_id)
                    ->where('id', '=', $id)
                    ->exists(),
            ],
        ]);

        return [ 'sender_user_id' => $user_id, ...$data ];
    }

    public function validateReceiverData(): array
    {
        ['receiver_user_id' => $receiver_user_id] = request()->validate([
            'receiver_user_id' => ['required', 'integer', 'exists:users,id' ],
        ]);

        ['receiver_account_id' => $receiver_account_id] = request()->validate([
            'receiver_account_id' => [
                'required', 'integer',
                fn($id) => BankAccount::query()
                    ->where('user_id', '=', auth()->id())
                    ->where('id', '=', $id)
                    ->exists(),
            ]
        ]);

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
     * Display the specified resource.
     */
    public function show(Transfer $transfer): JsonResponse
    {
        if (in_array(auth()->id(), [$transfer->sender_user_id, $transfer->receiver_user_id])) {
            abort(403);
        }

        return response()->json($transfer);
    }
}
