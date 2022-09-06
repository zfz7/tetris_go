import '@testing-library/cypress/add-commands'

Cypress.Commands.add('resetDatabase', () => cy.exec(`psql postgres://tetris_test:tetris_test@localhost:5433/tetris_test -c "TRUNCATE TABLE meta_data"`))