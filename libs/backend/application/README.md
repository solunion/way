# @way/backend-application

Libreria per la gestione delle applicazioni nel sistema multi-tenant.

## Funzionalità

- Creazione di nuove applicazioni
- Lettura e ricerca di applicazioni
- Aggiornamento di applicazioni esistenti
- Eliminazione di applicazioni
- Supporto multi-tenant

## API

- REST Endpoint: `/applications`
- GraphQL Query/Mutation: `Application`

## Configurazione

Importare il modulo `ApplicationModule` nel modulo principale dell'applicazione.

## Test

Eseguire i test con `nx test backend-application` 