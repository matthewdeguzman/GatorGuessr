describe("Visits Homepage", () => {
  it("Vists homepage of GatorGuessr", () => {
    cy.visit("http://localhost:4200/home");
    cy.url().should("eq", "http://localhost:4200/home");
  });
});

describe("Visits Login Page", () => {
  it("Vists login page of GatorGuessr", () => {
    cy.visit("http://localhost:4200/login");
    cy.url().should("eq", "http://localhost:4200/login");
  });
});

describe("Visits Registration Page", () => {
  it("Vists Registration page of GatorGuessr", () => {
    cy.visit("http://localhost:4200/register");
    cy.url().should("eq", "http://localhost:4200/register");
  });
});
describe("Buttons on Homepage", () => {
  it("Vists homepage of GatorGuessr", () => {
    cy.visit("http://localhost:4200/home");
    cy.get("a").contains("Login").click();
    cy.url().should("eq", "http://localhost:4200/login");
    cy.get("a").contains("Register").click();
    cy.url().should("eq", "http://localhost:4200/register");
    cy.get("[id^=home-button]").click();
    cy.url().should("eq", "http://localhost:4200/home");
  });
});
describe("Page not found", () => {
  it("Visits a page that does not exist", () => {
    cy.visit("http://localhost:4200/doesnotexist");
    cy.url().should("eq", "http://localhost:4200/page-not-found");
  });
});
describe("Page not found and goes back", () => {
  it("Visits a page that does not exist", () => {
    cy.visit("http://localhost:4200/home");
    cy.visit("http://localhost:4200/**");
    cy.url().should("eq", "http://localhost:4200/page-not-found");
    cy.get("button").contains("Go back to safety").click();
    cy.url().should("eq", "http://localhost:4200/home");
  });
});
