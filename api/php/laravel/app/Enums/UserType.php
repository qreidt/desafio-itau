<?php

namespace App\Enums;

enum UserType: string
{
    case Fisical = 'FISICAL';
    case Legal = 'LEGAL';

    public static function tryFromLength(string $document): self
    {
        return match (strlen($document)) {
            18 => UserType::Legal,
            14 => UserType::Fisical,
            default => throw new \Exception('Unmatched document type'),
        };
    }
}
