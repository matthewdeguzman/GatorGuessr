describe("Nonexistent User", () => {
  it("Tries to login with user that doesnt exist", () => {
    cy.visit("http://localhost:4200/login");
    cy.get("[id^=username]").type("NotAUser");
    cy.get("[id^=password]").type("password");
    cy.get("[id^=login-button]").click();
    cy.get("div").contains("Error: No User Found");
  });
});
describe("Incorrect Password", () => {
  it("Tries to login with incorrect password", () => {
    cy.visit("http://localhost:4200/login");
    cy.get("[id^=username]").type("Frontend");
    cy.get("[id^=password]").type("password");
    cy.get("[id^=login-button]").click();
    cy.get("div").contains("Error: Incorrect Password");
  });
});
describe("Correct Login", () => {
  it("Logs in with correct username and password", () => {
    cy.visit("http://localhost:4200/login");
    cy.get("[id^=username]").type("Frontend");
    cy.get("[id^=password]").type("FrontendIsNumber1");
    cy.get("[id^=login-button]").click();
    cy.url().should("eq", "http://localhost:4200/landing-page");
  });
});
