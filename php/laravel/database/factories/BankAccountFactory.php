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
        [$state, $city] = $this->getRandomStateAndCity();

        return [
            'user_id' => User::factory(),
            'number' => $this->randomNumber(),
            'agency' => str_pad($this->faker->numberBetween(1, 9999), 4, '0', STR_PAD_LEFT),
            'bank' => $this->faker->randomElement(Banks::cases()),
            'state' => $state,
            'city' => $city,
            'balance' => 0,
        ];
    }

    private function randomNumber(): string
    {
        return $this->faker->randomElement([
            '79892-0',
            '1994019-5',
            '45554322-5',
            '44181-9',
            '1253738-1',
            '234971-X',
            '1719459-3',
            '0516723-P',
            '0837119-9',
            '8409134402-0',
            '1082589-4',
            '1030889-X',
            '1299038-8',
            '47369-4',
            '69843-6',
            '0441755-0',
        ]);
    }

    private function getRandomStateAndCity(): array
    {
        $state = $this->faker->randomElement([
            'AC', 'AL', 'AP', 'AM', 'BA', 'CE',
            'DF', 'ES', 'GO', 'MA', 'MT', 'MS',
            'MG', 'PA', 'PB', 'PR', 'PE', 'PI',
            'RJ', 'RN', 'RS', 'RO', 'RR', 'SC',
            'SP', 'SE', 'TO',
        ]);

        $city = $this->faker->city();

        return [$state, $city];
    }
}
