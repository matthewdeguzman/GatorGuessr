describe('Nonexistent User', () => {
    it('Tries to login with user that doesnt exist', () => {
        cy.visit('http://localhost:4200/login')
        cy.get('#mat-input-0').type('NotAUser')
        cy.get('#mat-input-1').type('password')
        cy.get('button').click()
        cy.get('div').contains('Incorrect username or password')
    })
})
describe('Incorrect Password', () => {
    it('Tries to login with incorrect password', () => {
        cy.visit('http://localhost:4200/login')
        cy.get('#mat-input-0').type('Frontend')
        cy.get('#mat-input-1').type('password')
        cy.get('button').click()
        cy.get('div').contains('Incorrect username or password')
    })
})
describe('Correct Login', () => {
    it('Logs in with correct username and password', () => {
        cy.visit('http://localhost:4200/login')
        cy.get('#mat-input-0').type('Frontend')
        cy.get('#mat-input-1').type('FrontendIsNumber1')
        cy.get('button').click()
        cy.get('div').contains('Welcome back')
    })
})