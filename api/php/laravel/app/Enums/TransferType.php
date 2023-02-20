<?php

namespace App\Enums;

enum TransferType: string
{
    case PIX = 'PIX';
    case TED = 'TED';
    case DOC = 'DOC';
}
