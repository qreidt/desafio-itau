<?php

namespace App\Enums;

enum Banks: string
{
    case BcoBrasil = 'BCO_BRASIL';
    case Itau = 'ITAU';
    case Caixa = 'CAIXA';
    case Nubank = 'NUBANK';
    case Bradesco = 'BRADESCO';
    case Citibank = 'CITIBANK';
}
