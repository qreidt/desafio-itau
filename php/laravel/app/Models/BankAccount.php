<?php

namespace App\Models;

use App\Enums\Banks;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class BankAccount extends Model
{
    use HasFactory;

    protected $fillable = [
        'user_id',
        'number',
        'agency',
        'bank',
        'state',
        'city',
        'balance',
    ];

    protected $casts = [
        'bank' => Banks::class
    ];
}
