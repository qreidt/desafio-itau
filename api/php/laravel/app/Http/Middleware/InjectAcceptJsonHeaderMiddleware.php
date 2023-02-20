<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;

class InjectAcceptJsonHeaderMiddleware
{
    public function handle(Request $request, Closure $next)
    {
        if (! $request->hasHeader('accept')) {
            $accept = $request->header('accept');
            $request->headers->set('accept', "$accept;application/json");
        } else {
            $request->headers->set('accept', 'application/json');
        }

        return $next($request);
    }
}
