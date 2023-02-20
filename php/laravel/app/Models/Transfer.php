<?php

namespace App\Models;

use App\Enums\TransferType;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class Transfer extends Model
{

    public $timestamps = ['created_at'];
    const UPDATED_AT = null;

    protected $fillable = [
        'sender_user_id',
        'receiver_user_id',
        'sender_account_id',
        'receiver_account_id',
        'type', 'value',
    ];

    protected $casts = [
        'type' => TransferType::class
    ];

    public function sender_user(): BelongsTo
    {
        return $this->belongsTo(User::class, 'sender_user_id');
    }

    public function receiver_user(): BelongsTo
    {
        return $this->belongsTo(User::class, 'receiver_user_id');
    }

    public function sender_account(): BelongsTo
    {
        return $this->belongsTo(BankAccount::class, 'sender_account_id');
    }

    public function receiver_account(): BelongsTo
    {
        return $this->belongsTo(BankAccount::class, 'receiver_account_id');
    }
}
