<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers;
use App\Http\Controllers\Auth;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::post('/register', [Auth\RegisteredUserController::class, 'store'])
    ->name('register');

Route::post('/login', [Auth\AuthenticatedSessionController::class, 'store'])
    ->name('login');

Route::middleware('auth:sanctum')->group(function () {
    Route::get('/auth', function () {
        return auth()->user()->load('bank_accounts');
    });

    Route::delete('/logout', [Auth\AuthenticatedSessionController::class, 'destroy'])
        ->name('logout');

    Route::apiResource('/transfers', Controllers\TransferController::class);
});
