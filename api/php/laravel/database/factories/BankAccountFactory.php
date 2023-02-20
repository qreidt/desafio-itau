<?php

namespace Database\Factories;

use App\Enums\Banks;
use App\Models\BankAccount;
use App\Models\User;
use Illuminate\Database\Eloquent\Factories\Factory;

class BankAccountFactory extends Factory
{
    protected $model = BankAccount::class;
    public function definition(): array
    {
        $number = $this->getRandomNumber(length: $this->faker->numberBetween(6, 10));

        return [
            'user_id' => User::factory(),
            'number' => $number, 'balance' => 0,
        ];
    }

    private function getRandomNumber(int $length): string
    {
        $number = '';
        for ($i = 0; $i < $length; $i++) {
            $number .= $this->faker->randomDigit();
        }

        return "$number-". $this->faker->randomDigit();
    }
}
