describe('User that already exists', () => {
    it('Tries to login with user that already exists', () => {
        cy.visit('http://localhost:4200/register')
        cy.get('#mat-input-0').type('ramajanco')
        cy.get('#mat-input-1').type('Password123')
        cy.get('button').click()
        cy.get('div').contains('Error: User already exists')
    })
})