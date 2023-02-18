<?php

namespace App\Models;

use App\Enums\TransferType;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

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
}
