describe('Visits Homepage', () => {
  it('Vists homepage of GatorGuessr', () => {
    cy.visit('http://localhost:4200/home')
  })
})

describe('Visits Login Page', () => {
  it('Vists login page of GatorGuessr', () => {
    cy.visit('http://localhost:4200/login')
  })
})

describe('Visits Registration Page', () => {
  it('Vists Registration page of GatorGuessr', () => {
    cy.visit('http://localhost:4200/register')
  })
})
describe('Buttons on Homepage', () => {
  it('Vists homepage of GatorGuessr', () => {
    cy.visit('http://localhost:4200/home')
    cy.get('a').contains('Login').click();
    cy.get('a').contains('Register').click();
  })
})    