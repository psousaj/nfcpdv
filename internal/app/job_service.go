package app

/*
RegisterPDV, UpdatePDV, ScanPDVs, MigrateFolder, ExecuteSQL.

Aqui vai a regra de negócio:
Implantado = !sefazPresent
Idempotência

Preocupação: esta camada não fala direto com Fiber nem SSH, só com interfaces (ports).
*/
