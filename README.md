# go-lang-full-cycle-cobra-cli
Repositório para armazenar o código do módulo sobre cobra CLI do curso Go-Expert do Full Cycle


Para adicionar novas aplicações, use:
```bash
cobra-cli add <nome da aplicação>
```

Para criar comandos numa aplicação, use:
```bash
cobra-cli add <sub-comand> -p <'command principal, ex.: categoryCmd'>
```

## Rodando o projeto
- Dependências
    - [Cobra CLI](https://github.com/spf13/cobra)
- Criar as tabelas no banco de dados:
  ```sh
  sqlite3 database.db;
  CREATE TABLE categories (id string, name string, description string);
  CREATE TABLE courses (id string, name string, description string, category_id string);
  ```