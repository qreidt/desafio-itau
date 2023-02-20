<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     */
    public function up(): void
    {
        Schema::create('transfers', function (Blueprint $table) {
            $table->id();
            $table->foreignId('sender_user_id')->constrained('users')->cascadeOnUpdate()->cascadeOnUpdate();
            $table->foreignId('receiver_user_id')->constrained('users')->cascadeOnUpdate()->cascadeOnUpdate();
            $table->foreignId('sender_account_id')->constrained('bank_accounts')->cascadeOnUpdate()->cascadeOnDelete();
            $table->foreignId('receiver_account_id')->constrained('bank_accounts')->cascadeOnUpdate()->cascadeOnDelete();
            $table->enum('type', array_map(fn($enum) => $enum->value, \App\Enums\TransferType::cases()));
            $table->unsignedInteger('value');
            $table->timestamp('created_at');
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('transfers');
    }
};
