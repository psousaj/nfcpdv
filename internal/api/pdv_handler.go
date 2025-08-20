package api

/*
Criar rotas (/pdvs, /pdvs/:id/check, etc)

Endpoints
POST /pdvs / GET /pdvs / PATCH /pdvs/:id
POST /jobs/scan (body: filtros) → enfileira scans
POST /jobs/migrate (body: pdvIds|filtros, opções de backup) → rename
POST /jobs/sql (body: pdvIds|filtros, templateId|sqlRaw, params)
GET /jobs / GET /jobs/:id
GET /pdvs?implanted=true|false&online=true|false (usa último scan)
GET /health (API) e GET /health/worker


Handlers só chamam app.PdvService
*/
